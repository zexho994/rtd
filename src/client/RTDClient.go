package main

import (
	"flag"
	"github.com/zouzhihao-994/rtd/src"
)

var host = flag.String("host", "127.0.0.1", "listening host , ip4 or ip6 or localhost")
var port = flag.Int("port", 9944, "listening port , 5000 ~ 65535")

func main() {
	// 连接rtd
	rc, err := src.Dial(*host, *port)
	if err != nil {
		panic("rtd dial error" + err.Error())
	}
	_, err = rc.Write("hello")
	if err != nil {
		panic("rtd conn write error" + err.Error())
	}
}
