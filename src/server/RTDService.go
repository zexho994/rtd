package main

import (
	"flag"
	"github.com/zouzhihao-994/rtd/src"
)

var host = flag.String("host", "127.0.0.1", "listening host , ip4 or ip6 or localhost")
var port = flag.Int("port", 9944, "listening port , 5000 ~ 65535")

func main() {
	// bind
	ra, err := src.Bind(*host, *port)
	if err != nil {
		panic("rtd bind error " + err.Error())
	}
	// listen
	conn, err := src.Listen(ra)
	if err != nil {
		panic("rtd listen error " + err.Error())
	}
	conn.LoopRead()
}
