package src

import (
	"fmt"
	"net"
)

type RTDConn struct {
	conn net.Conn
}

// read msg from the other party
func (r *RTDConn) Read(b []byte) (msg string, err error) {
	l, err := r.conn.Read(b)
	if err != nil {
		return "", err
	}
	data := string(b[:l])
	fmt.Println("receive: " + data)
	return data, nil
}

// write msg to the other party
func (r *RTDConn) Write(msg string) (n int, err error) {
	l, err := r.conn.Write([]byte(msg))
	if err != nil {
		return -1, err
	}
	return l, nil
}

func (r *RTDConn) LoopRead() {
	data := make([]byte, 1024)
	for {
		_, err := r.Read(data)
		if err != nil {
			panic("rtd loop read error " + err.Error())
		}
	}
}

// close the rtd conn
func (r *RTDConn) Close() error {
	return r.conn.Close()
}
