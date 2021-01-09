package skill

import "gitlab.com/dasjott/alexa-sdk-go"

var Locales = alexa.Localisation{
	"fr-FR": alexa.Translation{
		"WELCOME_MSG":     "Bienvenue sur transport rennes!",
		"UPCOMINGBUS_MSG": "Le prochain bus arrivera à {{busstop}} dans 5 minutes.",
		"HELP_MSG":        "Je peux vous indiquer quand arrive le prochain bus. Dites simplement 'quand arrive le prochain bus'.",
		"GOODBYE_MSG":     "Au revoir!",
		"ERROR_MSG":       "Désolé, je n'ai pas compris. Pouvez-vous reformuler?",
	},
}
