package parser

import (

	"fmt"
)

func PageParser(seed string,startPage int, endPage int)[]Request {

	requestList := []Request{}
	for startPage < endPage{
		startPage += 1
		url := fmt.Sprintf("%s%d",seed,startPage)
		request := Request{
			Url:    url,
			Parser: titleListParser,
		}

		requestList = append(requestList, request)
	}
	return requestList
}
