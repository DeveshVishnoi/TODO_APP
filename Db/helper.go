package db

import (
	"context"
	"fmt"
	"todo/models"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var DbFunction DbHelperProvider

type DbHelperProvider interface {
	GetAllUsers() ([]models.User, error)
	GetUser(email_id string) (models.User, error)
	InsertUser(models.User) (string, error)
	GetTask(task_id string) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	InsertTaskInUser(models.Task) (string, error)
	UpdateTaskInUser(models.Task, string) (models.Task, error)
	DeleteTaskInUser(string) (string, error)
	GetUserTasks(string) ([]models.Task, error)
}

func (mongoDBManager *Manager) GetAllUsers() ([]models.User, error) {
	var users []models.User

	orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(UserCollectionName)
	cursor, err := orgCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error getting the Find Query to get Users", err)
		return users, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &users); err != nil {
		fmt.Println("Error decoding users", err)
		return users, err
	}

	return users, nil
}

func (mongoDBManager *Manager) GetUser(email_id string) (models.User, error) {

	fmt.Println("email_id ", email_id)

	var user models.User

	filter := bson.M{"email_id": email_id}
	orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(UserCollectionName)
	result := orgCollection.FindOne(context.TODO(), filter)
	err := result.Decode(&user)
	if err != nil {
		return user, err
	}

	fmt.Println("User", user)
	return user, nil

}

func (mongoDBManager *Manager) InsertUser(data models.User) (string, error) {
	fmt.Println("Insert data ----", data)

	isExits := false

	_, err := mongoDBManager.GetUser(data.EmailId)
	if err == mongo.ErrNoDocuments {
		isExits = false
	} else {
		isExits = true
		fmt.Println("Error", err)
		return "Same EmailID Data is Presnt...!!!", err
	}
	fmt.Println("isExist", isExits)

	if !isExits {
		orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(UserCollectionName)
		result, err := orgCollection.InsertOne(context.TODO(), data)
		if err != nil {
			fmt.Println("error getting when inserting the data")
			return "", err
		}

		fmt.Println(result.InsertedID, "Successfully Inserted")
		return "Successfully Inserted", nil
	}

	return "Same EmailID Data is Presnt...!!!", nil

}

func (mongoDBManager *Manager) GetTask(task_id string) (models.Task, error) {
	fmt.Println("task_id ", task_id)

	var task models.Task

	filter := bson.M{"email_id": task_id}
	orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(UserCollectionName)
	result := orgCollection.FindOne(context.TODO(), filter)
	err := result.Decode(&task)
	if err != nil {
		return task, err
	}

	fmt.Println("task", task)
	return task, nil
}

func (mongoDBManager *Manager) GetAllTasks() ([]models.Task, error) {

	var tasks []models.Task

	orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(TaskCollectionName)
	cursor, err := orgCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error getting the Find Query to get all tasks", err)
		return tasks, err
	}

	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &tasks); err != nil {
		fmt.Println("Error decoding tasks", err)
		return tasks, err
	}

	return tasks, nil

}

func (mongoDBManager *Manager) InsertTaskInUser(task_data models.Task) (string, error) {

	fmt.Println("Task_Data", task_data)

	// Check user is exit or not.

	orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(TaskCollectionName)
	result, err := orgCollection.InsertOne(context.TODO(), task_data)
	if err != nil {
		fmt.Println("Unable to insert task data in the Database")
		return "", err
	}
	fmt.Println(result.InsertedID, "Successfully inserted task")
	return fmt.Sprintf("%v", result.InsertedID), nil

}

func (mongoDBManager *Manager) UpdateTaskInUser(task_data models.Task, task_id string) (models.Task, error) {

	filter := bson.M{"uuid": task_id}
	orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(TaskCollectionName)
	result, err := orgCollection.ReplaceOne(context.TODO(), filter, task_data)
	if err != nil {
		fmt.Println("unable to replace data", err)
		return models.Task{}, err
	}
	fmt.Println(result.MatchedCount, result.ModifiedCount, "Successfully updated")
	return task_data, nil
}

func (mongoDBManager *Manager) DeleteTaskInUser(task_id string) (string, error) {

	filter := bson.M{"uuid": task_id}
	orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(TaskCollectionName)
	result, err := orgCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("unmbal to delete the task", err)
		return "unable to delete", err
	}
	if result.DeletedCount == 0 {
		return "no documents in the db", nil
	}

	return "Successfully Deleted Task...!!!", nil
}

func (mongoDBManager *Manager) GetUserTasks(user_id string) ([]models.Task, error) {

	var tasks []models.Task

	filter := bson.M{"user_id": user_id}
	orgCollection := mongoDBManager.Connection.Database(DatabaseName).Collection(TaskCollectionName)
	cursor, err := orgCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error getting the Find Query to get all tasks", err)
		return tasks, err
	}

	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &tasks); err != nil {
		fmt.Println("Error decoding tasks", err)
		return tasks, err
	}

	return tasks, nil
}
