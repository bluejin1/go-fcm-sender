package configs

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"google.golang.org/api/option"
)

var FcmClient *messaging.Client
var FcmCtx context.Context = context.Background()
var GoogleKeyFileData string = ``

// Firebase APP 초기화
func FcmInit() error {

	fmt.Println("[initApp] FcmInit start :")

	var err error

	type AccountKey struct {
		ProjectID string `json:"project_id"`
	}
	accountkey := AccountKey{}

	opt := option.WithCredentialsJSON([]byte(GoogleKeyFileData))

	if err := json.Unmarshal([]byte(GoogleKeyFileData), &accountkey); err != nil {
		fmt.Println("[initApp] getConfig error :", err)
		return err
	}
	config := &firebase.Config{ProjectID: accountkey.ProjectID}

	//app, err := firebase.NewApp(FcmCtx, config, opt)
	app, err := firebase.NewApp(FcmCtx, config, opt)
	if err != nil {
		fmt.Println("[initApp] initializing app error :", err)
		return err
	}

	// Obtain a messaging.Client from the App.
	FcmClient, err = app.Messaging(FcmCtx)
	if err != nil {
		fmt.Println("[initApp] getting Messaging client error :", err)
		return err
	}
	fmt.Println("[initApp] FcmClient :", FcmClient)
	return nil
}
