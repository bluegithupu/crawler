package engine

import (
	"exsample/crawler/caoliu/parser"
	"exsample/crawler/fetcher"
	"log"
)

func Run(seed string)  {
	startPage := 15
	endPage := 20
	requestList := []parser.Request{}
	requestList = append(requestList, parser.PageParser(seed,startPage,endPage)...)
	log.Println(requestList)

	for len(requestList) != 0{
		requst := requestList[0]
		requestList = requestList[1:]
		html, err := fetcher.Fetch(requst.Url)
		if err != nil{
			continue
		}
		newRequst := requst.Parser(html)
		if newRequst != nil{
			requestList = append(requestList, newRequst...)
		}
	}
}
