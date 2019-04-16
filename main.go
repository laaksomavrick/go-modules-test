package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello, world")
	uuid, _ := uuid.NewRandom()
	fmt.Println(uuid.String())
}
