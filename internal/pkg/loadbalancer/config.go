package loadbalancer

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// load balancer configuration
type config struct {
	ProxyPort string     `json:"proxyPort"`
	Backends  []*backend `json:"backends"`
}

func ReadConfig() *config {
	// reading config from JSON file
	var config config
	data, err := ioutil.ReadFile("./lb_config.json")
	if err != nil {
		log.Fatalf("error reading load balancer config: %v", err.Error())
	}
	json.Unmarshal(data, &config)

	return &config
}
