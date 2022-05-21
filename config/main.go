package config

import "github.com/andikanugraha11/social-monkey/config/schema"

type Schema struct {
	App schema.App
}

func Init(service string) (*Schema, error) {
	var conf Schema
	_ = readConfig(&conf, service)
	return &conf, nil
}
