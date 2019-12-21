package main

import (
	"exsample/crawler/caoliu/parser"
	"exsample/crawler/engine"
)


func main() {
	seed := parser.Domain + "thread0806.php?fid=16&search=&page="
	engine.Run(seed)
}














