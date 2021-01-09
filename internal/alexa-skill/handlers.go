package skill

import "gitlab.com/dasjott/alexa-sdk-go"

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
	// call odsClient.GetUpcomingBus
}
