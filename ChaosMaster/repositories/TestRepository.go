package repositories

import (
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
	MongoCollectionNameTestGroups string = "testgroups"
	MongoTestGroupIdFieldName     string = "_id"
)

var (
	MongoClient   *mongo.Client
	MongoDBName   string
	ServiceConfig models.ServiceConfig

	ErrorUpdateCountWrong = errors.New("Incorrect update count")
)

func init() {
	connectToMongo()

	ServiceConfig, _ = GetConfig()
	MongoDBName = ServiceConfig.MongoDBName
}

func connectToMongo() {
	config, err := GetConfig()
	if err != nil {
		log.Panicln("Unable to read service config to get mongoDB details: ", err.Error())
		os.Exit(1)
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", config.MongoDBHost, config.MongoDBPort))
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panicln("Unable to connect to mongo:", err.Error())
		os.Exit(1)
	}

	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Panicln("Unable to ping mongodb: ", err.Error())
		os.Exit(1)
	}

	log.Println("Mongo connection successfull!")
}

func AddTestGroup(testGroup models.TestGroup) error {
	collection := MongoClient.Database(MongoDBName).Collection(MongoCollectionNameTestGroups)

	_, err := collection.InsertOne(context.TODO(), testGroup)
	if err != nil {
		log.Println("Unable to insert TestGroup into MongoDB Collection: ", err.Error())
		return err
	}

	return nil
}

func GetTestGroup(id string) (testGroup models.TestGroup, errr error) {
	collection := MongoClient.Database(MongoDBName).Collection(MongoCollectionNameTestGroups)

	findFilter := bson.D{{MongoTestGroupIdFieldName, id}}

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

	return testGroup, nil
}

func DeleteAllTestGroups() error {
	collection := MongoClient.Database(MongoDBName).Collection(MongoCollectionNameTestGroups)

	filter := bson.D{}

	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println("Failed to delete all test groups: ", err.Error())
		return err
	}

	log.Println("Number of test groups deleted: ", result.DeletedCount)

	return nil
}

func UpdateTestGroup(testGroup models.TestGroup) error {
	collection := MongoClient.Database(MongoDBName).Collection(MongoCollectionNameTestGroups)

	findFilter := bson.M{MongoTestGroupIdFieldName: bson.M{"$eq": testGroup.ID}}

	result, err := collection.ReplaceOne(context.TODO(), findFilter, testGroup)
	if err != nil {
		log.Println("Error updating record: ", err.Error())
		return err
	}

	if result.ModifiedCount != 1 {
		log.Println("Update failed")
		return ErrorUpdateCountWrong
	}

	return nil
}

func GetAllTestGroups() (testGroups []models.TestGroup, err error) {
	collection := MongoClient.Database(MongoDBName).Collection(MongoCollectionNameTestGroups)

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

		testGroups = append(testGroups, group)
	}

	return testGroups, nil
}
