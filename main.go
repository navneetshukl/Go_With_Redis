package main

import (
	"Go_With_Redis/Redis"
	"encoding/json"
	"fmt"
	"log"
)

type Name struct {
	First string
	Last  string
	Age   int
}

func main() {

	fmt.Println("Go Redis Tutorial")

	Redis.InsertIntoRedis("name","Navneet Shukla")

	name:=Name{
		First: "Navneet",
		Last: "Shukla",
		Age :22,
	}
	json,err:=json.Marshal(name)
	if err!=nil{
		log.Fatal("Error in Marshaling Struct ",err)
	}
	Redis.InsertIntoRedis("json",json)

	Redis.GetFromRedis("name")
	Redis.GetFromRedis("json")


}
