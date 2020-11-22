package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/zouzhihao-994/rtd/src"
	"os"
)

var host = flag.String("host", "127.0.0.1", "listening host , ip4 or ip6 or localhost")
var port = flag.Int("port", 9944, "listening port , 5000 ~ 65535")

func main() {
	// 连接rtd
	rc, err := src.Dial(*host, *port)
	if err != nil {
		panic("rtd dial error" + err.Error())
	}
	fmt.Println("connect success...")
	for {
		// type
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		b, _, err := in.ReadLine()
		// write to rtd
		_, err = rc.WriteBytes(b)
		if err != nil {
			panic("rtd conn write error" + err.Error())
		}
	}
}
