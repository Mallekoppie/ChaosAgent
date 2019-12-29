package repositories

import (
	"mallekoppie/ChaosGenerator/ChaosMaster/models"

	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoCollectionNameTestGroups  string = "testgroups"
	MongoTestCollectionIdFieldName string = "testcollections.testCollectionId"
	MongoTestGroupIdFieldName      string = "testgroupid"
)

var (
	MongoClient   *mongo.Client
	MongoDBName   string
	ServiceConfig models.ServiceConfig
)

func init() {
	ConnectToMongo()

	ServiceConfig, _ = GetConfig()
	MongoDBName = ServiceConfig.MongoDBName
}

func ConnectToMongo() {
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

func DeleteTestCollection(testGroupId string, testCollectionId string) error {
	collection := MongoClient.Database(MongoDBName).Collection(MongoCollectionNameTestGroups)

	findFilter := bson.D{{MongoTestGroupIdFieldName, testGroupId}}

	findResult := collection.FindOneAndDelete(context.TODO(), findFilter)
	if findResult.Err() != nil {
		log.Println("Unable to find testgroup to delete: ", findResult.Err().Error())
		return findResult.Err()
	}

	testGroup := models.TestGroup{}
	err := findResult.Decode(&testGroup)
	if err != nil {
		log.Println("Unable to decode test group for find: ", err.Error())
		return err
	}

	var removeId int
	for id := range testGroup.TestCollections {
		item := testGroup.TestCollections[id]
		if item.TestCollectionId == testCollectionId {
			removeId = id
			break
		}
	}

	testGroup.TestCollections = append(testGroup.TestCollections[:removeId], testGroup.TestCollections[removeId+1:]...)

	return AddTestGroup(testGroup)
}
