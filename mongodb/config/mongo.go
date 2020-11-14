package config

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {
	hostname := getEnv("MONGO_DB_HOST", "belajarmongodb.4epud.mongodb.net")
	username := getEnv("MONGO_DB_USER", "reyhan_go")
	password := getEnv("MONGO_DB_PASS", "reyhan_go")
	dbname := getEnv("MONGO_DB_NAME", "belajar_go")

	// mongodb atlas
	CONNECTION_URI := "mongodb+srv://" + username + ":" + password + "@" + hostname + "/" + dbname + "?retryWrites=true&w=majority"
	fmt.Println(CONNECTION_URI)

	session, err := mgo.Dial(CONNECTION_URI)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connect to MongoDB")

	return session.DB(dbname), nil
}

func getEnv(key string, fallback string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		value = fallback
	}

	fmt.Println(value)

	return value
}
