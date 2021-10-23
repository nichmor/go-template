package events

import "go-template/pkg/db"

func StartApp() {
	db.ConnectMongo()
	db.ConnectRedis()
}

func StopApp() {
	db.CloseMongo()
	db.ConnectRedis()
}
