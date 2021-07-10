package main

import (
	"fmt"

	"github.com/osdifera/gRPC/proto"
)

func main(){
	doMessage()
}

func doMessage(){
	person := &proto.Person{
		Name: "Hiro",
		Lastname: "Watanabe",
		Age: 15,
	}

	fmt.Println(person)

	person.Age = 25
	fmt.Println(person)
}