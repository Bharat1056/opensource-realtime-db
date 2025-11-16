package main

import (
	"fmt"
	"log"

	"github.com/Bharat1056/real_time_db/photon"
)

func main() {

	// user := map[string]string{
	// 	"Name": "Bharat Panigrahi",
	// 	"Age": "21",
	// }


	db, err := photon.New()

	if err != nil {
		log.Fatal(err)
	}

	coll, err := db.CreateCollection("users")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", coll)

}
