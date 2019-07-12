package main

import (
	"fmt"
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
		fmt.Println("\tEverything as expected.")
	}
}

func runTests() {

	testDefaultBridges()
	testDirAuths()
	testDomains()
}

func main() {

	runTests()
}
