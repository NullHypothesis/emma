package main

import (
	"fmt"
)

const (
	SuccessMessage = "everything as expected"
)

func testDefaultBridges() {
	fmt.Println("Testing TCP port of default bridges.")
	for _, addrTuple := range defaultBridges {
		reachable, err := IsTCPPortReachable(addrTuple)
		if reachable {
			fmt.Printf("\treachable: %s\n", addrTuple)
		} else {
			fmt.Printf("\tunreachable: %s (%s)\n", addrTuple, err)
		}
	}
}

func testDirAuths() {
	fmt.Println("Testing TCP port of directory authorities.")
	for _, addrTuple := range dirAuths {
		reachable, err := IsTCPPortReachable(addrTuple)
		if reachable {
			fmt.Printf("\treachable: %s\n", addrTuple)
		} else {
			fmt.Printf("\tunreachable: %s (%s)\n", addrTuple, err)
		}
	}
}

func testDomains() {
	fmt.Println("Testing *.torproject.org domains.")
	var allGood = true
	for domain, addrMap := range domains {
		err := DoesDomainResolve(domain, addrMap)
		if err != nil {
			fmt.Printf("\tunexpected: %s (%s)\n", domain, err)
			allGood = false
		}
	}
	if allGood {
		fmt.Printf("\t%s\n", SuccessMessage)
	}
}

func testWebsites() {
	fmt.Println("Testing websites.")
	var allGood = true
	for url, substring := range websites {
		success, err := DoesWebsiteContainStr(url, substring)
		if err != nil {
			fmt.Println(err)
			allGood = false
		}
		if !success {
			fmt.Printf("\tno substring %q in url %q.\n", substring, url)
			allGood = false
		}
	}
	if allGood {
		fmt.Printf("\t%s\n", SuccessMessage)
	}
}

func main() {

	testDefaultBridges()
	testDirAuths()
	testDomains()
	testWebsites()
}
