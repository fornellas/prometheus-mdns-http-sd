package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fornellas/prometheus-mdns-http-sd/cli/lib"
	"github.com/fornellas/prometheus-mdns-http-sd/mdns"
)

type Target struct {
	Target string            `json:"target"`
	Labels map[string]string `json:"labels"`
}

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
		if req.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		entries, err := mdns.Query(
			lib.InterfaceStr,
			lib.Service,
			lib.Domain,
			lib.Timeout,
			lib.WantUnicastResponse,
			lib.DisableIPv4,
			lib.DisableIPv6,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error querying mDNS: %v", err)
			return
		}

		var targets []Target
		for _, entry := range entries {

			targets = append(targets, Target{
				Target: fmt.Sprintf("%s:%d", strings.TrimRight(entry.Host, "."), entry.Port),
				Labels: map[string]string{
					"name": entry.Name,
					"info": entry.Info,
				},
			})
		}

		jsonData, err := json.Marshal(targets)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error marshalling JSON: %v", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	return http.Server{
		Addr:    addr,
		Handler: serveMux,
	}
}
