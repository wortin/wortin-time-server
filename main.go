package main

import (
	"wortin-time-server/api"
	"wortin-time-server/conf"
)

func main() {
	port := conf.Vip.GetString("server.port")
	err := api.Server.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
