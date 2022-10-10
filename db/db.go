package db

import (
	"consumerreviewsgo/config"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var DB *elasticsearch.Client

func Init() {
	c := config.GetConfig()
	fmt.Println(c)
	cfg := elasticsearch.Config{
		Addresses: []string{
			c.GetString("elastic.server"),
		},
	}
	db, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}

func GetDb() *elasticsearch.Client {
	return DB

}
