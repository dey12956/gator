package main

import (
	"fmt"
	"log"

	"github.com/dey12956/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("dey")
	if err != nil {
		log.Fatal(err)
	}

	cfgNew, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(cfgNew)
}
