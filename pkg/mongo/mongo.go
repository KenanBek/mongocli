/*
Copyright © 2019 KANAN RAHIMOV <mail@kenanbek.me>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

func (mc *MongoClient) CountDocuments(db, coll string) (int64, error) {
	filter := bson.D{{}}
	collection := mc.client.Database(db).Collection(coll)

	return collection.CountDocuments(mc.context, filter)
}

// ListDocuments PRINTS documents in a given database and collection (TODO)
func (mc *MongoClient) ListDocuments(databaseName, collectionName string) {
	filter := bson.D{{}}
	collection := mc.client.Database(databaseName).Collection(collectionName)
	cur, err := collection.Find(mc.context, filter)
	if err != nil {
		log.Fatal("Error on finding all the documents", err)
	}
	for cur.Next(mc.context) {
		fmt.Println(cur)
	}
}
