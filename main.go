package main

import (
	"fcm-sender/cmd"
	"fcm-sender/configs"
	"fmt"
	"log"
	"time"
)

var (
	GitCommit string
	BuildTime string
)

func main() {

	// start time
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("start time : ", time.Now().UTC())
	log.Println(configs.SERVICE_NAME)
	now := time.Now().UnixNano() / 1000000
	fmt.Printf("%d", now)

	BuildTime = time.Now().String()
	//cmd.SgEventServiceServerStart(GitCommit, BuildTime)

	cmd.OrderServerStart(GitCommit, BuildTime)

}
