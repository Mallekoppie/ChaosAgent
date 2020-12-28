package repositories

import (
	"github.com/Mallekoppie/goslow/platform"
	"go.uber.org/zap"
	"mallekoppie/ChaosGenerator/ChaosMaster/models"

	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoCollectionNameTestGroups      string = "testgroups"
	mongoCollectionNameTestCollections string = "testcollections"
	mongoIdFieldName                   string = "_id"
	mongoGroupIdFieldName              string = "groupid"
)

var (
	mongoClient   *mongo.Client
	mongoDBName   string
	serviceConfig models.ServiceConfig

	ErrUpdateCountWrong                 = errors.New("Incorrect update count")
	ErrNoGroupExistsForTestCollectionId = errors.New("No group exists for ID")
)

func init() {
	connectToMongo()

	serviceConfig, _ = GetConfig()
	mongoDBName = serviceConfig.MongoDBName
}

func connectToMongo() {
	config, err := GetConfig()
	if err != nil {
		log.Panicln("Unable to read service config to get mongoDB details: ", err.Error())
		os.Exit(1)
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", config.MongoDBHost, config.MongoDBPort))
	mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panicln("Unable to connect to mongo:", err.Error())
		os.Exit(1)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Panicln("Unable to ping mongodb: ", err.Error())
		os.Exit(1)
	}

	log.Println("Mongo connection successfull!")
}

func AddTestGroup(testGroup models.TestGroup) error {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestGroups)
	testCollections := testGroup.TestCollections
	testGroup.TestCollections = nil

	_, err := collection.InsertOne(context.TODO(), testGroup)
	if err != nil {
		log.Println("Unable to insert TestGroup into MongoDB Collection: ", err.Error())
		return err
	}

	for index := range testCollections {
		item := testCollections[index]
		item.GroupId = testGroup.ID

		err = AddTestCollection(item)
		if err != nil {
			logger.Error("Error saving test collection for test group: ", err.Error())
			return err
		}
	}

	return nil
}

func GetTestGroup(id string) (testGroup models.TestGroup, errr error) {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestGroups)

	findFilter := bson.D{{mongoIdFieldName, id}}

	testGroup = models.TestGroup{}

	result := collection.FindOne(context.TODO(), findFilter)
	if result.Err() != nil {
		log.Println("Error retrieving Test Group: ", result.Err().Error())
		return testGroup, result.Err()
	}

	err := result.Decode(&testGroup)
	if err != nil {
		log.Println("Unable to decode test group: ", err.Error())
		return testGroup, err
	}

	testCollections, err := GetTestCollectionsForGroup(testGroup.ID)
	if err != nil {
		logger.Error("unable to retrieve test collections for Test Group: ", err.Error())

	} else {
		testGroup.TestCollections = testCollections
	}

	return testGroup, nil
}

func DeleteAllTestGroups() error {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestGroups)

	filter := bson.D{}

	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println("Failed to delete all test groups: ", err.Error())
		return err
	}

	testCollection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestGroups)

	_, err = testCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println("Failed to delete all test collections: ", err.Error())
		return err
	}

	log.Println("Number of test groups deleted: ", result.DeletedCount)

	return nil
}

func DeleteTestGroup(id string) error {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestGroups)

	filter := bson.D{{mongoIdFieldName, id}}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		logger.Error("Unable to delete record:", err.Error())
		return err
	}

	if result.DeletedCount != 1 {
		logger.Error("Delete count incorrect: ", result.DeletedCount)
	}

	testCollection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestCollections)

	testCollectionFilter := bson.D{{mongoGroupIdFieldName, id}}

	collectionResult, err := testCollection.DeleteMany(context.TODO(), testCollectionFilter)
	if err != nil {
		logger.Error("Error when deleting Test Collections linked to Test Group: ", err.Error())
		return err
	}

	logger.Info("Test Collections deleted for test group: ", collectionResult.DeletedCount)

	return nil
}

func DeleteTestCollection(id string) error {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestCollections)

	filter := bson.D{{mongoIdFieldName, id}}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		logger.Error("Unable to delete record:", err.Error())
		return err
	}

	if result.DeletedCount != 1 {
		logger.Error("Delete count incorrect: ", result.DeletedCount)
	}

	return nil
}

func UpdateTestGroup(testGroup models.TestGroup) error {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestGroups)

	testCollections := testGroup.TestCollections
	testGroup.TestCollections = nil

	findFilter := bson.M{mongoIdFieldName: bson.M{"$eq": testGroup.ID}}

	result, err := collection.ReplaceOne(context.TODO(), findFilter, testGroup)
	if err != nil {
		log.Println("Error updating record: ", err.Error())
		return err
	}

	if result.ModifiedCount != 1 {
		log.Println("TestGroup update failed")
		return ErrUpdateCountWrong
	}

	for index := range testCollections {
		col := testCollections[index]

		err = UpdateTestCollection(col)
		if err != nil {
			logger.Error("Unable to update TestCollection: ", err.Error())
			return err
		}
	}

	return nil
}

func UpdateTestCollection(col models.TestCollection) error {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestCollections)

	findFilter := bson.M{mongoIdFieldName: bson.M{"$eq": col.ID}}

	result, err := collection.ReplaceOne(context.TODO(), findFilter, col)
	if err != nil {
		logger.Error("unable to replace Test Collection: ", err.Error())
		return err
	}

	if result.ModifiedCount != 1 {
		log.Println("Test Collection update failed for: ", col.ID)

		err = AddTestCollection(col)
		if err != nil {
			logger.Info("Could not update so we will add it as a new test collection")
			return err
		}
	}

	return nil
}

func GetAllTestGroups() (testGroups []models.TestGroup, err error) {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestGroups)

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error retrieving all test groups: ", err.Error())
		return nil, err
	}

	testGroups = make([]models.TestGroup, 0)

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		group := models.TestGroup{}

		err = cursor.Decode(&group)
		if err != nil {
			log.Println("Error decoding testgroup: ", err.Error())
			return nil, err
		}

		result, err := GetTestCollectionsForGroup(group.ID)
		if err != nil {
			logger.Error("Unable to retrieve TestCollections for Test Group: ", err.Error())
		} else {
			group.TestCollections = result
		}

		testGroups = append(testGroups, group)
	}

	return testGroups, nil
}

func AddTestCollection(tests models.TestCollection) error {
	groupCollection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestGroups)

	findGroupFilter := bson.D{{mongoIdFieldName, tests.GroupId}}

	groupResult := groupCollection.FindOne(context.TODO(), findGroupFilter)
	if groupResult.Err() != nil {
		if groupResult.Err() == mongo.ErrNoDocuments {
			logger.Error("The Test Group must be created before the test collection")
			return ErrNoGroupExistsForTestCollectionId
		} else {
			logger.Error("Error finding group for test collection: ", groupResult.Err().Error())
			return groupResult.Err()
		}
	}

	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestCollections)

	_, err := collection.InsertOne(context.TODO(), tests)
	if err != nil {
		logger.Error("Error inserting document: ", err.Error())
		return err
	}

	return nil
}

func GetTestCollectionsForGroup(id string) (tests []models.TestCollection, err error) {
	tests = make([]models.TestCollection, 0)
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionNameTestCollections)

	findFilter := bson.D{{mongoGroupIdFieldName, id}}

	cursor, err := collection.Find(context.TODO(), findFilter)
	if err != nil {
		logger.Error("Unable to retrieve test collections for Test Group: ", err.Error())
		return tests, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		test := models.TestCollection{}
		err = cursor.Decode(&test)
		if err != nil {
			logger.Error("Unable to decode test collection: ", err.Error())
			return tests, err
		}

		tests = append(tests, test)
	}

	return tests, nil
}
