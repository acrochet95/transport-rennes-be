package skill

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/acrochet95/transport-rennes-be/internal/opendatasoft"
	"gitlab.com/dasjott/alexa-sdk-go"
)

var odsClient *opendatasoft.OpendatasoftClient

func Initialize() {
	config := opendatasoft.ODSConfig{BaseUrl: os.Getenv("OPENDATASOFT_BASE_URL"), ApiKey: os.Getenv("OPENDATASOFT_API_KEY")}
	odsClient = opendatasoft.New(config)
}

var Handlers = alexa.IntentHandlers{
	"LaunchRequest": func(c *alexa.Context) {
		c.Tell(c.T("WELCOME_MSG"))
	},
	"AMAZON.HelpIntent": func(c *alexa.Context) {
		c.Tell(c.T("HELP"))
	},
	"AMAZON.FallbackIntent": func(c *alexa.Context) {
		c.Tell(c.T("ERROR_MSG"))
	},
	"UpcomingBusIntent":   upcomingBus,
	"AMAZON.StopIntent":   bye,
	"AMAZON.CancelIntent": bye,
}

func bye(c *alexa.Context) {
	c.Tell(c.T("GOODBYE_MSG"))
}

func upcomingBus(c *alexa.Context) {
	busStop := c.Slot("busstop")
	busName := c.Slot("bus")

	upcomingBus := odsClient.GetUpcomingBus(busName.Value, busStop.Value, "")

	// If no bus left
	if upcomingBus.NHits == 0 {
		c.Tell(c.T("NO_BUS_AVAILABLE"))
	}

	// sort records by departure time
	sort.SliceStable(upcomingBus.Records, func(i, j int) bool {
		return upcomingBus.Records[i].Information.Departure.Before(upcomingBus.Records[j].Information.Departure)
	})

	// store records by destionation
	x := make(map[string][]opendatasoft.UpcomingBusRecord)
	for _, record := range upcomingBus.Records {
		x[record.Information.Destination] = append(x[record.Information.Destination], record)
	}

	// Generate message
	var message string
	// Loop over destination
	for destination, busList := range x {
		if len(busList) >= 2 {
			message = message + fmt.Sprintf(c.T("UPCOMING_TWO_BUSES_MSG"), busName, busStop, destination,
				getDelay(&busList[0].Information.Departure),
				getDelay(&busList[1].Information.Departure))
		} else if len(busList) == 1 {
			message = message + fmt.Sprintf(c.T("UPCOMING_ONE_BUS_MSG"), busName, busStop, destination, getDelay(&busList[0].Information.Departure))
		}
	}

	// Send the final message
	c.Tell(message)
}

// Return delay before departure in minutes
func getDelay(departure *time.Time) int {
	return int(departure.Sub(time.Now().UTC()).Minutes())
}
