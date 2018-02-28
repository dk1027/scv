package main

import (
	"flag"
	"fmt"
	"github.com/dk1027/scv/core"
	"io/ioutil"
	"os"
)

func main() {
	configFile := flag.String("conf", "", "Config file path (required)")
	flag.Parse()
	if *configFile == "" {
		fmt.Println("conf is required.")
		os.Exit(1)
	}
	fmt.Printf("config: %s\n", *configFile)

	raw, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Printf("Error opening config file: %+v\n", err.Error())
		os.Exit(1)
	}

	tc, _ := scv.ReadConfig(raw)

	scv.NewEngine(tc)
}
