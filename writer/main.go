package main

import (
	"github.com/golang/protobuf/proto"
	userpb "github.com/menyasosali/grpc-golang/gen/go/user/v1"
	"log"
	"os"
)

func main() {
	u := userpb.User{
		Uuid:      "1-2-3-4",
		FullName:  "Mario",
		BirthYear: 1900,
	}

	b, err := proto.Marshal(&u)
	if err != nil {
		log.Fatalln("Marshaling error", err)
	}

	if err = os.WriteFile("user.bin", b, 0644); err != nil {
		log.Fatalln("Writing error", err)
	}
}
