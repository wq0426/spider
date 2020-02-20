package main

import (
	_ "fmt"
	"spider/config"
	"spider/engine"
	"spider/parser"
)

func main() {
	url := config.GetSetting("Home::Url", 0)
	engine.Run(engine.Request{
		Url:        url.(string),
		ParserFunc: parser.ParseCityList,
	})
}
