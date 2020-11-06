package src

import (
	"net"
	"strconv"
)

func Dial(host string, port int) (*RTDConn, error) {
	if host == "" {
		return nil, RTDError{msg: "host is blank"}
	}
	if port < 0 || port >= 65535 {
		return nil, RTDError{msg: "port is invalid"}
	}

	// udp dial
	conn, err := net.Dial("udp", host+":"+strconv.Itoa(port))

	//todo syn ack

	if err != nil {
		return nil, err
	}
	rc := RTDConn{
		conn: conn,
	}
	return &rc, nil
}
