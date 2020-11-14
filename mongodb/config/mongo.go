package config

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {
	hostname := getEnv("MONGO_DB_HOST", "localhost:27017")
	dbname := getEnv("MONGO_DB_NAME", "belajar_golang")
	// username := getEnv("MONGO_DB_USER", "reyhan_go")
	// password := getEnv("MONGO_DB_PASS", "reyhan_go")

	CONNECTION_URI := "mongodb://" + hostname + "/" + dbname
	fmt.Println(CONNECTION_URI)

	// info := &mgo.DialInfo{
	// 	Addrs:    []string{hostname},
	// 	Timeout:  10 * time.Second,
	// 	Database: dbname,
	// 	Username: username,
	// 	Password: password,
	// }

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

	return value
}
