package parser

import (
	_ "fmt"
	"regexp"
	"spider/config"
	"spider/engine"
)

//解析信息
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(config.GetSetting("City::Template", 0).(string))
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, c := range all {
		result.Items = append(result.Items, "User:"+string(c[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c)
			},
		})
	}

	return result
}
