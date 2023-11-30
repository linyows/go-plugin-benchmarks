package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		os.Exit(1)
	}
	m.Run()
}

func wait(addr string) {
	for {
		conn, _ := net.DialTimeout("tcp", addr, time.Second)
		if conn != nil {
			conn.Close()
			break
		}
	}
}

func setup() error {
	go ServeSOPlugin()
	wait("localhost" + SOPluginServerAddr)

	go ServeAPI()
	wait("localhost" + APIServerAddr)

	go ServeAPIPlugin()
	wait("localhost" + APIPluginServerAddr)

	go ServeCLIPlugin()
	wait("localhost" + CLIPluginServerAddr)

	return nil
}

func req(addr string) (error, string) {
	res, err := http.Get("http://localhost" + addr)
	if err != nil {
		return err, ""
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err, ""
	}

	return nil, string(b)
}

func BenchmarkSO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err, _ := req(SOPluginServerAddr + "/so")
		if err != nil {
			fmt.Printf("shared object plugin server error: %s", err)
		}
	}
}

func BenchmarkAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err, _ := req(APIPluginServerAddr + "/api")
		if err != nil {
			fmt.Printf("api plugin server error: %s", err)
		}
	}
}

func BenchmarkCLI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err, _ := req(CLIPluginServerAddr + "/cli")
		if err != nil {
			fmt.Printf("cli plugin server error: %s", err)
		}
	}
}
