package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"io/ioutil"
	"log"

	"./src/simple"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)

	smAsString := toJSON(sm)
	fmt.Println("smAsString:", smAsString)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}

	readFromFile("simple.bin", sm2)
	fmt.Println("Read Sm2 content:", sm2)
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read from file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)

	if err2 != nil {
		log.Fatalln("Couldn't put bytes into protobuf struct", err2)
		return err2
	}

	return nil
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SimpleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)
	fmt.Println("The Id is: ", sm.GetId())

	return &sm
}
