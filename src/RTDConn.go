package src

import (
	"net"
	"syscall"
)

type RTDConn struct {
	conn net.Conn
	fdSet RtdFdSet
}

type RtdFdSet struct {
	read  syscall.FdSet
	write syscall.FdSet
	exec  syscall.FdSet
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

// 无等待时间的select
func (r *RTDConn) Select() {

}

// 指定等待时间
//func (r *RTDConn) Select(timeout_sec int64) error {
//	timeVal := syscall.Timeval{Sec: timeout_sec, Usec: 0, Pad_cgo_0: [4]byte{}}
	// unix 使用整数数组存储fd描述符，
	// 假如使用int32数组，第一个描述符在1～31，第二个在32～63，第三个在64～95
	// 而poll函数没有这个数组长度的限制，这也是poll函数与select的主要区别
	//err := syscall.Select(1, &r.fdSet.read, &r.fdSet.write, &r.fdSet.exec, &timeVal)
	//if err != nil {
	//	return RTDError{msg: "select error: " + err.Error()}
	//}
	//return nil
//}
// write msg to the other party
func (r *RTDConn) WriteStr(msg string) (n int, err error) {
	return r.conn.Write([]byte(msg))
}

func (r *RTDConn) LoopRead() {
	data := make([]byte, 1024)
	for {
		_, err := r.conn.Read(data)
		if err != nil {
			panic("rtd loop read error " + err.Error())
		}
	}
}

// close the rtd conn
func (r *RTDConn) Close() error {
	return r.conn.Close()
}
