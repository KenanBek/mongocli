package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoClient is exported.
type MongoClient struct {
	host     string
	port     int
	uri      string
	client   *mongo.Client
	context  context.Context
	database *mongo.Database

	Cancel func()
}

// New is exported.
func New(host string, port int) *MongoClient {
	mongoClient := MongoClient{}
	var err error
	// client
	mongoClient.uri = fmt.Sprintf("mongodb://%s:%d", host, port)
	mongoClient.client, err = mongo.NewClient(options.Client().ApplyURI(mongoClient.uri))
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database new client error"))
	}

	// context and cancel
	mongoClient.context, mongoClient.Cancel = context.WithTimeout(context.Background(), 10*time.Second)

	// connection
	err = mongoClient.client.Connect(mongoClient.context)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database connection error"))
	}

	return &mongoClient
}

// Ping connection.
func (mc *MongoClient) Ping() error {
	return mc.client.Ping(mc.context, readpref.Primary())
}

// ListDatabaseNames returns all database names.
func (mc *MongoClient) ListDatabaseNames() ([]string, error) {
	filter := bson.D{{}}
	return mc.client.ListDatabaseNames(mc.context, filter)
}

// ListCollectionNames returns collection names in a given database.
func (mc *MongoClient) ListCollectionNames(databaseName string) ([]string, error) {
	filter := bson.D{{}}
	return mc.client.Database(databaseName).ListCollectionNames(mc.context, filter)
}

// CountDocuments return number of documents in the collection.
func (mc *MongoClient) CountDocuments(db, coll string) (int64, error) {
	filter := bson.D{{}}
	collection := mc.client.Database(db).Collection(coll)

	return collection.CountDocuments(mc.context, filter)
}

// ListDocuments return list of decoded documents and errors during the process.
func (mc *MongoClient) ListDocuments(databaseName, collectionName string) ([]bson.M, []error) {
	var docs []bson.M
	var errs []error

	filter := bson.D{{}}
	collection := mc.client.Database(databaseName).Collection(collectionName)

	cur, err := collection.Find(mc.context, filter)
	if err != nil {

		log.Fatal("Error on finding all the documents", err)
	}

	for cur.Next(mc.context) {
		var doc bson.M
		err := cur.Decode(&doc)
		if err != nil {
			errs = append(errs, err)
		} else {
			docs = append(docs, doc)
		}
	}

	if err := cur.Err(); err != nil {
		errs = append(errs, err)
	}

	return docs, errs
}
