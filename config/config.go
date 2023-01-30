package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token string

	config *configStruct
)

type configStruct struct {
	Token string `json:"token"`
}

func ReadConfig() error {
	fmt.Println("Считываю config.json...")

	file, err := os.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	return nil
}
