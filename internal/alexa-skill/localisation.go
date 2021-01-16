package skill

import "gitlab.com/dasjott/alexa-sdk-go"

var Locales = alexa.Localisation{
	"fr-FR": alexa.Translation{
		"WELCOME_MSG":            "Bienvenue sur transport rennes!",
		"UPCOMING_ONE_BUS_MSG":   "Le %s, arrêt %s, direction %s: prochain départ dans %s minutes.",
		"UPCOMING_TWO_BUSES_MSG": "Le %s, arrêt %s, direction %s: prochain départ dans %s minutes, le suivant dans %s minutes.",
		"NO_BUS_AVAILABLE":       "Aucun bus n'est actuellement prévu",
		"HELP_MSG":               "Je peux vous indiquer quand arrive le prochain bus. Dites simplement 'quand arrive le prochain bus'.",
		"GOODBYE_MSG":            "Au revoir!",
		"ERROR_MSG":              "Désolé, je n'ai pas compris. Pouvez-vous reformuler?",
	},
}
