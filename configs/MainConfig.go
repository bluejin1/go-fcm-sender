package configs

import (
	"fcm-sender/helper/env"
	"sync"
)

var (
	mainServerConfig     MainServer
	mainServerConfigOnce sync.Once
)

type MainServer struct {
	Name    string
	Address AddressInfo
}

func DefaultMainServerConfigFromEnv() MainServer {
	mainServerConfigOnce.Do(func() {
		mainServerConfig = MainServer{
			Name: env.GetEnv("SERVICE_NAME", SERVICE_NAME),
			Address: AddressInfo{
				Host: env.GetEnv("HOST_NAME", SERVICE_HOST_NAME),
				Port: env.GetEnv("GRPC_PORT", SERVICE_PORT),
			},
		}
	})
	return mainServerConfig
}
