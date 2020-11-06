package src

type RTDAddr struct {
	host string
	port int
}

func (r RTDAddr) Host() string {
	return r.host
}

func (r RTDAddr) Port() int {
	return r.port
}
