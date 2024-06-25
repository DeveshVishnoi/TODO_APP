package models

type User struct {
	UUID     string `json:"uuid" bson:"uuid"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password"`
	EmailId  string `json:"email_id" bson:"email_id"`
}

type Task struct {
	UUID     string `json:"uuid" bson:"uuid"`
	TaskName string `json:"task_name" bson:"task_name"`
	TaskDate int64  `json:"task_date" bson:"task_date"`
	UserId   string `json:"user_id" bson:"user_id"`
	Desc     string `json:"desc" bson:"desc"`
}
