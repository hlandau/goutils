package os

// Returns the machine hostname as a fully-qualified domain name.
func MachineFQDNUncached() (string, error) {
	return machineFQDN()
}

var cachedFQDN = ""

// Returns the machine hostname as a fully-qualified domain name. Caches the
// hostname for fast lookup. Since retrieval of the hostname can involve
// process execution on some systems, you should use this instead of
// MachineFQDNUncached unless you need to support hostname changes.
func MachineFQDN() (string, error) {
	if cachedFQDN != "" {
		return cachedFQDN, nil
	}

	fqdn, err := MachineFQDNUncached()
	if err != nil {
		return "", err
	}

	cachedFQDN = fqdn
	return fqdn, nil
}
