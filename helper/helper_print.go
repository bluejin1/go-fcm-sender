package helper

import (
	"encoding/json"
	"fcm-sender/helper/env"
	"fmt"
)

func PrettyPrint(data interface{}) {
	if env.GetEnvEnvironment() != "development" {
		return
	}

	var p []byte
	//    var err := perror
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
