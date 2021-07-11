package main

import (
	"fmt"
	"log"

	"github.com/osdifera/gRPC/pbmodels"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main(){
	ms:=doPerson()
	msAsString:=toJson(ms)
	fmt.Printf("%s \n",msAsString)

	pb:=&pbmodels.Person{}
	fromJson(msAsString,pb)
	fmt.Println(pb)
}

func toJson(pb proto.Message) []byte {

	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatal("The message could not being marshaled", err)
		return nil
	}
	return out
}

func fromJson(str []byte, pb proto.Message) {
	err := protojson.Unmarshal(str, pb)
	if (err != nil){
		log.Fatal("Could not unmarshal ", err)
	}
}

func doPerson() *pbmodels.Person{
	person := &pbmodels.Person{
		Name: "Hiro",
		Lastname: "Watanabe",
		Age: 15,
	}

	fmt.Println(person)
	person.Age = 25
	fmt.Println(person)

	return person
}