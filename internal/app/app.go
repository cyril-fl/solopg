package app

import (
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/mongo"
	"solopg/internal/infrastructure/t"
	"solopg/internal/platform/log"
)

func Start() error {
	if err := t.Init(config.Current.I18n, ""); err != nil {
		return err
	}

	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	return runSession(db)
}

func Try() error {

	log.ParseJson(config.Current)

	return nil
}

/*
	TODO
	LOW
	# Étendre la CLI avec des commande comme:

	> Cli edit —config —locales [lang]
		- d’autre argument pouvant être passé en fonction de la config par défaut.
	> Cli edit —rulesset [rules]
		- rules étant les class, races et autre document yaml
	> Cli reset —config …
		- Prend les meme argument que edit
		- ajoute —all, —save
			- trouver un truc genre —all-remember pour effacer tout sauf les save

	Adapter un système comme « mole » avec des commande qui font les action rapide . Et une commande qui ouvre un genre de menu:

	> Cli config

	En vu de ça créer un doublon des data / config files avec justement un dossier comme « default », un dossier genre « user » et mette les deux dans une approche à la VS code config
*/
