package main

import (
	"fmt"
	"log"

	mackerelsesame "github.com/Arthur1/mackerel-sesame"
	"github.com/caarlos0/env/v8"
)

// TODO: DeviceName と UUID をセットで複数渡せるように
type config struct {
	MackerelAPIKey            string `env:"MACKEREL_API_KEY"`
	MackerelServiceName       string `env:"MACKEREL_SERVICE_NAME"`
	SesameAPIKey              string `env:"SESAME_API_KEY"`
	SesameDeviceUUID          string `env:"SESAME_DEVICE_UUID"`
	SesameDeviceNameForMetric string `env:"SESAME_DEVICE_NAME_FOR_METRIC"`
}

func main() {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	fetcher := mackerelsesame.NewSesameFetcher(cfg.SesameAPIKey)
	result, err := fetcher.Fetch(cfg.SesameDeviceUUID, cfg.SesameDeviceNameForMetric)
	if err != nil {
		log.Fatal(err)
	}
	exporter := mackerelsesame.NewMackerelExporter(cfg.MackerelAPIKey, cfg.MackerelServiceName)
	if err := exporter.Export(result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}
