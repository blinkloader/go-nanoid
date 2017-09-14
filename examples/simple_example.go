package main

import (
	"fmt"
	"github.com/matoous/go-nanoid"
)

func main(){
	id, err := gonanoid.Generate()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated default id: %s\n", id)
}