package main

import (
	"encoding/json"
    "log"
    "net/http"
	"time"
	"os"
)

type server struct{}

type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	resp := HandsOn{
		Time:     time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}

    w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
	//resp := fmt.Sprintf("La hora es %v y el hostname es %v", time.Now(), os.Getenv("HOSTNAME"))
    w.Write(jsonResp)
}

func main() {
    s := &server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(":9090", nil))
}
