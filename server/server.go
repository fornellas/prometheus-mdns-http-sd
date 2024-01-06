package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/fornellas/prometheus-mdns-http-sd/cli/lib"
	"github.com/fornellas/prometheus-mdns-http-sd/mdns"
)

type Targets struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
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

		var targetsList []Targets
		for _, entry := range entries {
			var ip *net.IP
			if entry.AddrV4 != nil {
				ip = &entry.AddrV4
			}
			if entry.AddrV6 != nil {
				ip = &entry.AddrV6
			}
			if ip == nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Discovered service has no IP: %v", entry)
				return
			}

			targetsList = append(targetsList, Targets{
				Targets: []string{
					fmt.Sprintf("%v:%d", ip, entry.Port),
				},
				Labels: map[string]string{
					"mdns_host": entry.Host,
					"mdns_port": fmt.Sprintf("%d", entry.Port),
					"mdns_name": entry.Name,
					"mdns_info": entry.Info,
				},
			})
		}

		jsonData, err := json.Marshal(targetsList)
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
