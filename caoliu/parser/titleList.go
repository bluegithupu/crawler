package parser

import (

	"log"
	"regexp"
)

const Domain  = "http://cc.jb6v.icu/"

func titleListParser(html string) []Request {

	re := regexp.MustCompile(`<a href="(.+.html)" target="_blank" id="">`)
	match := re.FindAllStringSubmatch(html,-1)
	list := []Request{}
	for _, m := range match{
		url := Domain + m[1]
		log.Println("url:", url)
		request := Request{
			Url:    url,
			Parser: imgeListParse,
		}
		list = append(list, request)
	}
	return list
}
