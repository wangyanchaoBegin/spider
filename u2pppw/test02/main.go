package main

import (
	"u2pppw/test/engine"
	"u2pppw/test/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url : "http://www.zhenai.com/zhenghun",
		ParserFunc : parser.ParserCityList,

	})
}
