package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// timeout specifies the number of seconds we're willing to wait until we
// decide that the given destination is offline.
const timeout time.Duration = 3 * time.Second

type Result struct {
	Target        string
	Error         error
	ExecutionTime time.Duration
}

func (r Result) String() string {
	if r.Error == nil {
		return fmt.Sprintf("  ✓ %-35s %6s\n", r.Target, r.ExecutionTime.Round(time.Millisecond))
	} else {
		return fmt.Sprintf("  ✗ %-35s %6s  error: %v\n", r.Target, r.ExecutionTime.Round(time.Millisecond), r.Error)
	}
}

func showExecutionTime(start time.Time, r *Result) {
	r.ExecutionTime = time.Since(start)
}

// IsTCPPortReachable returns `true' if it can establish a TCP connection with
// the given IP address and port.  If not, it returns `false' and the
// respective error, as reported by `net.DialTimeout'.
func IsTCPPortReachable(addrTuple string) (r Result) {

	defer func(s time.Time) {
		r.ExecutionTime = time.Since(s)
	}(time.Now())
	r.Target = addrTuple

	conn, err := net.DialTimeout("tcp", addrTuple, timeout)
	if err != nil {
		r.Error = err
		return
	}
	conn.Close()
	r.Error = nil
	return
}

func DoesDomainResolve(domain string, expected map[string]bool) error {
	addrs, err := net.LookupHost(domain)
	if err != nil {
		return err
	}

	for _, addr := range addrs {
		if !expected[addr] {
			return fmt.Errorf("%q != %q", addrs, expected)
		}
	}
	return nil
}

func DoesWebsiteContainStr(rawurl, substring string) (r Result) {

	defer func(s time.Time) {
		r.ExecutionTime = time.Since(s)
	}(time.Now())

	u, err := url.Parse(rawurl)
	if err == nil {
		r.Target = u.Host
	} else {
		r.Target = rawurl
	}

	resp, err := http.Get(rawurl)
	if err != nil {
		r.Error = err
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Error = err
		return
	}

	if strings.Contains(string(body), substring) {
		r.Error = nil
	} else {
		r.Error = errors.New("could not find expected string in HTML body")
	}
	return
}
