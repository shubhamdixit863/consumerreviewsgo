package server

import (
	"consumerreviewsgo/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
