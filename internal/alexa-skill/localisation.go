package skill

import "gitlab.com/dasjott/alexa-sdk-go"

var Locales = alexa.Localisation{
	"fr-FR": alexa.Translation{
		"WELCOME_MSG":            "Bienvenue sur transport rennes! Que voulez vous savoir ?",
		"UPCOMING_ONE_BUS_MSG":   "Le {bus}, arrêt {busstop}, direction {destination}: prochain départ dans {dep1} minutes.",
		"UPCOMING_TWO_BUSES_MSG": "Le {bus}, arrêt {busstop}, direction {destination}: prochain départ dans {dep1} minutes, le suivant dans {dep2} minutes.",
		"NO_BUS_AVAILABLE":       "Aucun bus n'est actuellement prévu",
		"HELP_MSG":               "Je peux vous indiquer quand arrive le prochain bus. Dites simplement 'quand arrive le prochain bus'.",
		"GOODBYE_MSG":            "Au revoir!",
		"ERROR_MSG":              "Désolé, je n'ai pas compris. Pouvez-vous reformuler?",
	},
}
