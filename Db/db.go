package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DatabaseName = "TODO_APP"
var UserCollectionName = "Users"
var TaskCollectionName = "Tasks"

type Manager struct {
	Connection *mongo.Client
	Ctx        context.Context
}

var MongoManager Manager

func ConnectDB(URL string) (*Manager, error) {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))
	if err != nil {
		fmt.Println("unable to make connection ", err)
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Unable to ping : ", err)
		return nil, err
	}
	MongoManager = Manager{Connection: client, Ctx: ctx}
	fmt.Println("Database Connected successfully.....!!!!!!!!!!!!")
	return &MongoManager, nil

}

func (mongoDBManager *Manager) DisconnectDb() {
	ctx := mongoDBManager.Ctx
	err := mongoDBManager.Connection.Disconnect(ctx)
	if err != nil {
		fmt.Println("unable to Disconnect the Db")
		return
	}
	fmt.Println("Successfully Disconnect the Db......!!!!!")

}
