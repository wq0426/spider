package parser

import (
	_ "fmt"
	"regexp"
	"spider/config"
	"spider/engine"
)

//解析城市信息
func ParseCityList(contents []byte) engine.ParseResult {
	tem := config.GetSetting("CityList::Template", 0)
	re := regexp.MustCompile(tem.(string))
	all := re.FindAllSubmatch(contents, -1)
	length := len(all)                                 //城市总数量（470个）
	need_len := config.GetSetting("Common::Length", 1) //需要采集的城市数量
	i := 0
	result := engine.ParseResult{}
	for _, c := range all {
		result.Items = append(result.Items, string(c[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseCity(c)
			},
		})
		i++
		if (i == need_len) || (i == length) {
			break
		}
	}

	return result
}
