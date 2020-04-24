package main

import (
	"bufio"
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"time"
)

var buf bytes.Buffer
var l = log.New(&buf, "", 0)

func testDefaultBridges() {
	l.Println("Testing default bridges:")

	for _, addrTuple := range defaultBridges {
		r := IsTCPPortReachable(addrTuple)
		l.Printf(r.String())
	}
}

func testWebsites() {
	l.Print("Testing websites:")

	for url, substring := range websites {
		r := DoesWebsiteContainStr(url, substring)
		l.Print(r.String())
	}
}

func testDirAuths() {
	l.Println("Testing directory authorities.")

	for _, addrTuple := range dirAuths {
		r := IsTCPPortReachable(addrTuple)
		l.Print(r.String())
	}
}

func main() {

	log.Println("Starting analysis.")
	var useStdout bool
	flag.BoolVar(&useStdout, "stdout", false, "Write results to stdout instead of a file.")
	flag.Parse()

	var f *os.File
	var err error
	if useStdout {
		f = os.Stdout
	} else {
		f, err = ioutil.TempFile("", "emma-")
		if err != nil {
			log.Fatal(err)
		}
	}

	l.Printf("Time: %s\n", time.Now())
	testDefaultBridges()
	testWebsites()
	testDirAuths()

	if _, err := f.WriteString(buf.String()); err != nil {
		os.Remove(f.Name())
		log.Fatalf("Failed writing output to %q because: %s", f.Name(), err)
	} else {
		log.Printf("Wrote output to: %s", f.Name())
	}
	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

	// On Windows, we don't want the terminal window to close after the tool is
	// done; otherwise the user won't know where the output file is.
	if runtime.GOOS == "windows" {
		log.Printf("Press enter to exit.")
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}
}
