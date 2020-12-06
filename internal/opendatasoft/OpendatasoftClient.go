package opendatasoft

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const endpoint string = "/api/records/1.0/search/?dataset=tco-bus-circulation-passages-tr" +
	"&q=&facet=nomcourtligne&facet=destination&facet=precision&facet=arrte&facet=nomarret" +
	"&refine.nomcourtligne=%s&refine.nomarret=%s&refine.destination=%s&refine.precision=Temps+réel"

type OpendatasoftClient struct {
	client *http.Client
	config ODSConfig
}

// Create new instance of OpendatasoftClient
func New(config ODSConfig) *OpendatasoftClient {
	httpClient := OpendatasoftClient{}
	httpClient.client = &http.Client{Timeout: 10 * time.Second}
	httpClient.config = config

	return &httpClient
}

func (ods *OpendatasoftClient) GetUpcomingBus(busLineName string, stopName string, destination string) *UpcomingBus {
	var upcomingBus UpcomingBus
	ods.getRequest(fmt.Sprintf(endpoint, busLineName, stopName, destination), &upcomingBus)

	log.Printf("nhits: %d et size: %d", upcomingBus.NHits, len(upcomingBus.Records))
	for i, record := range upcomingBus.Records {
		delay := record.Information.Departure.Sub(time.Now())
		log.Printf("%d: departure at %s (in %s)", i, record.Information.Departure, delay)
	}

	return &upcomingBus
}

// Get request to opendatasoft api
func (ods *OpendatasoftClient) getRequest(request string, target interface{}) error {
	req, _ := http.NewRequest("GET", ods.config.BaseUrl+request, nil)
	req.Header.Set("Authorization", ods.config.ApiKey)
	resp, err := ods.client.Do(req)
	if err != nil {
		log.Panic(err)
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&target)
	return nil
}
