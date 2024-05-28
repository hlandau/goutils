package net

import "net"

// Returns true if the given string is a hostname:port string, and false if it
// is just a hostname. Returns an error if the string cannot be parsed.
func HasPort(hostport string) (bool, error) {
	_, _, err := net.SplitHostPort(hostport)
	if err != nil {
		if e, ok := err.(*net.AddrError); ok && e.Err == "missing port in address" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// Like net.SplitHostPort, but supports absence of port. For example,
// "192.0.2.1" is accepted.
func TrySplitHostPort(hostport string) (host, port string, hasPort bool, err error) {
	host, port, err = net.SplitHostPort(hostport)
	if err != nil {
		if e, ok := err.(*net.AddrError); ok && e.Err == "missing port in address" {
			host, port, err = net.SplitHostPort(hostport + ":")
		}
	} else {
		hasPort = true
	}

	return
}
