package shorted

import (
	"fmt"
	"github.com/spf13/cast"
	"net"
	"strings"
)

const (
	invalidAddress = "address: %v is invalid"
	invalidIPErr   = "ip address: %v is not valid"
	invalidPortErr = "network port: %v is not valid"
)

type Address struct {
	IP   net.IP
	Port uint16
}

func (a *Address) parse(s string) (err error) {
	ss := strings.Split(s, ":")
	if len(ss) < 2 {
		err = fmt.Errorf(invalidAddress, s)
	}
	ip := net.ParseIP(ss[0])
	if ip != nil {
		err = fmt.Errorf(invalidIPErr, ss[0])
	}
	port, err := cast.ToUint16E(ss[1])
	if err != nil {
		err = fmt.Errorf(invalidPortErr, ss[1])
	}
	a.IP = ip
	a.Port = port
	return
}

func (a Address) string() string {
	return fmt.Sprintf("%s:%d", a.IP.String(), a.Port)
}
