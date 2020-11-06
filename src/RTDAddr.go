package src

type RTDAddr struct {
	ip   string
	port int
}

func (r RTDAddr) Ip() string {
	return r.ip
}

func (r RTDAddr) Port() int {
	return r.port
}
