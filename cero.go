package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

/* result of processing a domain name */
type procResult struct {
	addr  string
	names []string
	err   error
}

// run parameters (filled from CLI arguments)
var (
	verbose              bool
	concurrency          int
	defaultPorts         []string
	timeout              int
	onlyValidDomainNames bool
)

var usage = "" +
	`usage: cero [options] [targets]
if [targets] not provided in commandline arguments, will read from stdin
`

func main() {
	// parse CLI arguments
	var ports string

	flag.BoolVar(&verbose, "v", false, `Be verbose: Output results as 'addr -- [result list]', output errors to stderr as 'addr -- error message'`)
	flag.IntVar(&concurrency, "c", 100, "Concurrency level")
	flag.StringVar(&ports, "p", "443", "TLS ports to use, if not specified explicitly in host address. Use comma-separated list")
	flag.IntVar(&timeout, "t", 4, "TLS Connection timeout in seconds")
	flag.BoolVar(&onlyValidDomainNames, "d", false, "Output only valid domain names (e.g. strip IPs, wildcard domains and gibberish)")

	// set custom usage text
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		fmt.Fprintln(os.Stderr, "options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	// parse default port list into string slice
	defaultPorts = strings.Split(ports, `,`)

	// channels
	chanInput := make(chan string)
	chanResult := make(chan *procResult)

	// a common dialer
	dialer := &net.Dialer{
		Timeout: time.Duration(timeout) * time.Second,
	}

	// create and start concurrent workers
	var workersWG sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		workersWG.Add(1)
		go func() {
			for addr := range chanInput {
				result := &procResult{addr: addr}
				result.names, result.err = grabCert(addr, dialer, onlyValidDomainNames)
				chanResult <- result
			}
			workersWG.Done()
		}()
	}

	// close result channel when workers are done
	go func() {
		workersWG.Wait()
		close(chanResult)
	}()

	// create and start result-processing worker
	var outputWG sync.WaitGroup
	outputWG.Add(1)
	go func() {
		for result := range chanResult {
			// in verbose mode, print all errors and results, with corresponding input values
			if verbose {
				if result.err != nil {
					fmt.Fprintf(os.Stderr, "%s -- %s\n", result.addr, result.err)
				} else {
					fmt.Fprintf(os.Stdout, "%s -- %s\n", result.addr, result.names)
				}
			} else {
				// non-verbose: just print scraped names, one at line
				for _, name := range result.names {
					fmt.Fprintln(os.Stdout, name)
				}
			}
		}
		outputWG.Done()
	}()

	// consume output to start things moving
	if len(flag.Args()) > 0 {
		for _, addr := range flag.Args() {
			processInputItem(addr, chanInput, chanResult)
		}
	} else {
		// every line of stdin is considered as a input
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			addr := strings.TrimSpace(sc.Text())
			processInputItem(addr, chanInput, chanResult)
		}
	}

	// close input channel when input fully consumed
	close(chanInput)

	// wait for processing to finish
	outputWG.Wait()
}

// process input item
// if orrors occur during parsing, they are pushed straight to result channel
func processInputItem(input string, chanInput chan string, chanResult chan *procResult) {
	_ = "STUB: not implemented"
	// initial inputs are skipped
	return
}

// split input to host and port (if specified)

// get ports list to use

// use ports from default list if not specified explicitly

// CIDR?

// expand CIDR

// feed IPs from CIDR to input channel

// feed atomic host to input channel

/*
	connects to addr and grabs certificate information.

returns slice of domain names from grabbed certificate
*/
func grabCert(addr string, dialer *net.Dialer, onlyValidDomainNames bool) ([]string, error) {
	_ = "STUB: not implemented"
	// dial
	return nil, nil
}

// get first certificate in chain

// get CommonName and all SANs into a slice

// append all SANs, excluding one that is equal to CN (if any)
