package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"
	"strings"
	"net"
	"strconv"
	"os"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getCurrentIpAddress() (string) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		ip := addr.String()
		if strings.Contains(ip, "/") {
			if !strings.HasPrefix(ip, "127") {
				return ip[:strings.Index(ip, "/")]
			}
		}
	}
	return ""

}

var (
	cont = 0
	status = 200
)
var server string = ""

func main() {

	version := os.Getenv("SYS_VERSION")

	http.HandleFunc("/counter/", func (w http.ResponseWriter, r *http.Request) {
		if len(server) == 0 { server = getCurrentIpAddress()  }
		cont+=1
		fmt.Printf("req=counter, server=%s, req=%d, version=%s\n", server,  cont, version)
		fmt.Fprintf(w, "req=counter, server=%s, req=%d, version=%s", server,  cont, version)
	})

	http.HandleFunc("/health/", func (w http.ResponseWriter, r *http.Request) {

		secondsStr := r.URL.Query().Get("wait-time-seconds")
		if len(secondsStr) == 0 {
			secondsStr = os.Getenv("WAIT_TIME_SECONDS")
		}
		seconds, err := strconv.Atoi(secondsStr)
		fmt.Printf("req=health, server=%s, req=%d, seconds=%d, secondsErr=%v, version=%s\n", server,  cont, seconds, err, version)
		time.Sleep(time.Duration(seconds) * time.Second)
		w.WriteHeader(status)
		fmt.Fprintf(w, "status=%d, version=%s", status, version)

	})

	http.HandleFunc("/status/", func (w http.ResponseWriter, r *http.Request) {
		rawStatus := r.URL.Query().Get("status")
		status, _ = strconv.Atoi(rawStatus)
		fmt.Printf("req=status, server=%s, req=%d, rawStatus=%s, status=%d\n", server,  cont, rawStatus, status)
	})
	http.ListenAndServe(":8080", nil)
}
