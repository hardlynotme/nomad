package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

func main() {

	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		panic(err)
	}

	err = session.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", os.Getenv("MONGO_URL"))
}
