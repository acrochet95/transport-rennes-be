package skill

import "github.com/dasjott/alexa-sdk-go"

var Locales = alexa.Localisation{
	"fr-FR": alexa.Translation{
		"WELCOME_MSG":            "Bienvenue sur transport rennes! Vous pouvez me demander quand arrive le prochain bus ?",
		"UPCOMING_ONE_BUS_MSG":   "Le {bus}, arrêt {busstop}, direction {destination}: prochain départ dans {dep1} minutes. ",
		"UPCOMING_TWO_BUSES_MSG": "Le {bus}, arrêt {busstop}, direction {destination}: prochain départ dans {dep1} minutes, le suivant dans {dep2} minutes. ",
		"TOO_MANY_BUSES":         "Beaucoup de bus sont disponibles à l'arrêt {busstop}. Veuillez préciser votre bus dans la question.",
		"NO_BUS_AVAILABLE":       "Aucun bus {bus} n'est actuellement prévu à l'arrêt {busstop}",
		"FAVORITE_SAVED":         "L'arrêt de bus {busstop} a été ajouté en favori",
		"FAVORITE_DELETED":       "Votre favori a été supprimé",
		"NO_FAVORITE":            "Aucun arrêt de bus n'est enregistré en favori",
		"FAVORITE_UNAVAILABLE":   "Le service de favori est momentanément indisponible",
		"HELP_MSG":               "Je peux vous indiquer quand arrive le prochain bus. Dites simplement: quand arrive le prochain bus ?",
		"GOODBYE_MSG":            "Au revoir!",
		"ERROR_MSG":              "Désolé, je n'ai pas compris. Pouvez-vous reformuler?",
	},
}
