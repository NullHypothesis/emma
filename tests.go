package main

import (
	"fmt"
	"net"
	"time"
)

// timeout specifies the number of seconds we're willing to wait until we
// decide that the given destination is offline.
const timeout time.Duration = 3 * time.Second

// IsTCPPortReachable returns `true' if it can establish a TCP connection with
// the given IP address and port.  If not, it returns `false' and the
// respective error, as reported by `net.DialTimeout'.
func IsTCPPortReachable(addrTuple string) (bool, error) {

	conn, err := net.DialTimeout("tcp", addrTuple, timeout)
	if err != nil {
		return false, err
	}
	conn.Close()
	return true, nil
}

func DoesDomainResolve(domain string, expected map[string]bool) error {
	addrs, err := net.LookupHost(domain)
	if err != nil {
		return err
	}

	for _, addr := range addrs {
		if !expected[addr] {
			return fmt.Errorf("%s != %s", addrs, expected)
		}
	}
	return nil
}
