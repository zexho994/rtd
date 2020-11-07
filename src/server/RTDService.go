package main

import (
	"flag"
	"fmt"
	"github.com/zouzhihao-994/rtd/src"
)

var host = flag.String("host", "127.0.0.1", "listening host , ip4 or ip6 or localhost")
var port = flag.Int("port", 9944, "listening port , 5000 ~ 65535")

func main() {
	// bind
	ra, err := src.RTDBind(*host, *port)
	if err != nil {
		panic("rtd bind error " + err.Error())
	}

	// listen
	conn, err := src.RTDListen(ra)
	if err != nil {
		panic("rtd listen error " + err.Error())
	}
	fmt.Println("start listening...")

	// read
	data := make([]byte, 1024)
	for {
		msg, err := conn.Read(data)
		if err != nil {
			fmt.Println("rtd conn read error ,", err.Error())
		} else {
			fmt.Println("rcv: ", msg)
		}
	}
}
