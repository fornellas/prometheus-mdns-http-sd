package server

import (
	"net/http"
	"time"
)

func NewServer(
	addr string,
	interfaceStr string,
	service string,
	domain string,
	timeout time.Duration,
	wantUnicastResponse bool,
	disableIPv4 bool,
	disableIPv6 bool,
) http.Server {

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/targets", func(w http.ResponseWriter, req *http.Request) {

	})

	return http.Server{
		Addr:    addr,
		Handler: serveMux,
	}
}
