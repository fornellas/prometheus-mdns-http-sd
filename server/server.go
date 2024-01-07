package server

import (
	"encoding/json"
	"fmt"
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
	disableIPv4 bool,
	disableIPv6 bool,
) http.Server {
	m := mdns.NewMDNS()

	proto := mdns.ProtoAny
	if lib.DisableIPv4 {
		proto = mdns.ProtoInet6
	}
	if lib.DisableIPv6 {
		proto = mdns.ProtoInet
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/targets", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		services, err := m.BrowseServices(
			lib.InterfaceStr,
			proto,
			lib.Service,
			lib.Domain,
			lib.Timeout,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error querying mDNS: %v", err)
			return
		}

		var targetsList []Targets
		for _, service := range services {
			targetsList = append(targetsList, Targets{
				Targets: []string{
					fmt.Sprintf("%v:%d", service.IP, service.Port),
				},
				Labels: map[string]string{
					"mdns_host":         service.Host,
					"mdns_port":         fmt.Sprintf("%d", service.Port),
					"mdns_name":         service.Name,
					"mdns_service_type": service.Type,
					"mdns_interface":    service.Interface,
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
