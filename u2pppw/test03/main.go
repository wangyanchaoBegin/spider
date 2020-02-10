package main

import (
	"u2pppw/test/engine"
	"u2pppw/test/scheduler"
	"u2pppw/test/zhenai/parser"
)

func main() {
	//
	//engine.SimpleEngine.Run(engine.Request{
	//	Url : "http://www.zhenai.com/zhenghun",
	//	ParserFunc : parser.ParserCityList,
	//
	//})


	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleSheduler{},
		WorkCount: 10 ,
	}

	e.Run(engine.Request{
		Url : "http://www.zhenai.com/zhenghun",
		ParserFunc : parser.ParserCityList,

	})


}
