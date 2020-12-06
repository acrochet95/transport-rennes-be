package main

import (
	"log"
	"time"

	"github.com/acrochet95/transport-rennes-be/internal/opendatasoft"
)

func main() {
	config := opendatasoft.ReadConfigFile("config.json")
	log.Printf("%s", config.BaseUrl)

	odsClient := opendatasoft.New(*config)
	upcomingBus := odsClient.GetUpcomingBus("C1", "Metz Volney", "Chantepie")

	log.Printf("nhits: %d et size: %d", upcomingBus.NHits, len(upcomingBus.Records))
	for i, record := range upcomingBus.Records {
		delay := record.Information.Departure.Sub(time.Now())
		log.Printf("%d: departure at %s (in %s)", i, record.Information.Departure, delay)
	}
}
