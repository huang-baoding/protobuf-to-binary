package main

import (
	"fmt"
	"github.com/huang-baoding/protobuf-to-binary/sample"
	"github.com/huang-baoding/protobuf-to-binary/serializer"
)

func main() {

	binaryFile := "./tmp/laptop.bin"
	jasonFile := "./tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	fmt.Println(err)
	fmt.Printf("	laptop1:	%v\n\n", laptop1)

	laptop2 := sample.NewLaptop()
	fmt.Printf("	laptop2.1:	%v\n\n", laptop1)
	err = serializer.ReadProtobufFromBinaryFile(laptop2, binaryFile)
	fmt.Println(err)
	fmt.Printf("	laptop2.2:	%v\n\n", laptop1)

	err = serializer.WriteProtobufToJSONFile(laptop1, jasonFile)
	fmt.Println(err)
}
