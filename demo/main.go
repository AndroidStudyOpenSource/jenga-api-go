package main

import (
	"github.com/AndroidStudyOpenSource/jenga-api-go"
	"log"
)

const (
	username = "8485200649"                        // sandbox --> change to yours
	password = "Hb8jahNZDnPjCw0T9RXxWH8KGvwQJweZK" // sandbox --> change to yours
)

func main() {

	jeng, err := jenga.New(username, password, jenga.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := jeng.Auth()
	if err != nil {
		log.Println(err)
	}
	log.Println(res)

}
