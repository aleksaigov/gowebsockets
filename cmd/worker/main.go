package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"syscall"

	"github.com/vlad2095/wss/pkg/pool"
	"github.com/vlad2095/wss/pkg/websocket"
)

func main() {
	// Increase resources limitations
	if runtime.GOOS == "linux" {
		var rLimit syscall.Rlimit
		if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
			panic(err)
		}

		rLimit.Cur = rLimit.Max
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
			panic(err)
		}
		fmt.Println("linux started")
	}

	// Enable pprof hooks
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	pool := pool.NewPool(128)
	channels := websocket.NewChannelPool()
	echo := func(c *websocket.Channel, op websocket.OpCode, data []byte) {
		// echo
		c.Send(op, data)
	}
	wh, _ := websocket.NewHandler(echo, channels)
	wh.SetPool(pool)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wh.Upgrade(w, r)
	})
	fmt.Println("start")

	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")
	if err := http.ListenAndServeTLS(":11112", certFile, keyFile, nil); err != nil {
		log.Fatal(err)
	}
}
