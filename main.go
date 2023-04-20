package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"phonenumber/gen"
)

func main() {
	phone := &gen.PhoneNumber{}
	bytes := []byte{10, 45, 10, 8, 74, 111, 104, 110, 32, 68, 111, 101, 16, 210, 9, 26, 16, 106, 100, 111, 101, 64, 101, 120, 97, 109, 112, 108, 101, 46, 99, 111, 109, 34, 12, 10, 8, 53, 53, 53, 45, 52, 51, 50, 49, 16, 3}
	err := proto.Unmarshal(bytes, phone)
	if err != nil {
		fmt.Errorf("error occured during unmarshalling %s", err)
	}
	fmt.Print(phone.Number)
}
