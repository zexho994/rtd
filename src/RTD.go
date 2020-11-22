package src

import "net"

type RTD struct {
}

func RTDBind(host string, port int) (*RTDAddr, error) {
	if host == "" {
		return nil, RTDError{msg: "host is blank"}
	}
	if port < 0 || port >= 65535 {
		return nil, RTDError{msg: "port is invalid"}
	}
	return &RTDAddr{host: host, port: port}, nil
}

func RTDListen(addr *RTDAddr) (*RTDConn, error) {
	ip := net.ParseIP(addr.host)
	ua := net.UDPAddr{IP: ip, Port: addr.port, Zone: ""}
	conn, err := net.ListenUDP("udp", &ua)
	if err != nil {
		return nil, RTDError{msg: "rtd listen error , " + err.Error()}
	}
	rc := RTDConn{conn: conn}
	return &rc, nil
}
