package src

import (
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
	return string(b[:l]), nil
}

func (r *RTDConn) WriteBytes(b []byte) (int, error) {
	l, err := r.conn.Write(b)
	if err != nil {
		return -1, RTDError{msg: "rtd conn write bytes error"}
	}
	return l, nil
}

// write msg to the other party
func (r *RTDConn) WriteStr(msg string) (n int, err error) {
	return r.WriteBytes([]byte(msg))
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
