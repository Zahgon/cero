package main

import (
	"regexp"
)

/*
	expands IP/IPv6 CIDR into atomic IPs

returns channel from which string IPs must be consumed
returns error if mask is too wide, or CIDR is not syntaxed properly
supported masks:
  - for IPv4: /[0-32] (whole IPv4 space)
  - for IPv6: /[64-128]: (up to 2^64 IPs)
*/
func expandCIDR(CIDR string) (chan string, error) {
	_ = "STUB: not implemented"
	// parse CIDR
	return nil, nil
}

// general check for unsupported cases

// create channel to deliver output

// switch branch to IPv4 / IPv6

// IPv4:

// convert to uint32, for convenient bitwise operation

// create buffer

// build IP as byte slice

// yield stringified IP

// IPv6

// convert lower halves to uint64, for convenient bitwise operation

// write portion of IP that will not change during expansion

// build IP as byte slice

// yield stringified IP

/*
	every value with slash is condiered as CIDR

if it's not a valid one, it will fail at later processing
*/
func isCIDR(value string) bool { _ = "STUB: not implemented"; return false }

var portRegexp, bracketRegexp *regexp.Regexp

func init() {
	portRegexp = regexp.MustCompile(`^(.*?)(:(\d+))?$`)
	bracketRegexp = regexp.MustCompile(`^\[.*\]$`)
}

/*
	parses input addr into -> host, port.

if port is not specified, returns ports as empty string.
tolerates IPv6 port specification without enclosing IP into square brackets.
in truly ambiguous cases for IPv6, treat as portless
Doesn't check for errors, just splits
*/
func splitHostPort(addr string) (host, port string) {
	_ = "STUB: not implemented"
	// split host and port
	return "", ""
}

// skip further checks for bracketed IPv6

// no port found, skip futher checks

// skip futher checks for CIDR

// check ambiguous cases for IPv6

// if port is longer than 4 digits -> it is truly a port

// cancel port if whole thing parses as valid IPv6

// isDomainName checks if a string is a presentation-format domain name
// (currently restricted to hostname-compatible "preferred name" LDH labels and
func isDomainName(s string) bool {
	_ = "STUB: not implemented"
	// See RFC 1035, RFC 3696.
	// Presentation format has dots before every label except the first, and the
	// terminal empty label is optional here because we assume fully-qualified
	// (absolute) input. We must therefore reserve space for the first and last
	// labels' length octets in wire format, where they are necessary and the
	// maximum total length is 255.
	// So our _effective_ maximum is 253, but 254 is not rejected if the last
	// character is a dot.
	return false
}

// true once we've seen a letter or hyphen

// fine

// Byte before dash cannot be dot.

// Byte before dot cannot be dot, dash.
