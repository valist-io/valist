package main

import (
	"fmt"
	"log"

	"github.com/valist-io/valist/internal/command"
)

func main() {
	md, err := command.NewApp().ToMarkdown()
	if err != nil {
		log.Fatalf("failed to gen docs %v", err)
	}

	fmt.Println(md)
}
