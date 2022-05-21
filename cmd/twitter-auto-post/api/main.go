package main

import "github.com/andikanugraha11/social-monkey/domain/twitter-auto-post/cmd/api"

const (
	service = "twitter-auto-post"
	mode    = "API"
)

func main() {
	// var conf *config.Schema
	// conf, err := config.Init(service)
	// if err != nil {
	// 	log.Fatal("cannot read config: ", err.Error())
	// }

	// log.Println(conf)

	api.Init(service)

}
