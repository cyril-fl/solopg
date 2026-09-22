package src

import (
	"solopg/app/domain/card/characters"
	"solopg/app/services/mongo"
	"solopg/app/services/t"
	"solopg/app/utils/log"
	"solopg/app/utils/session"
	"solopg/config"
)

func Start() error {
	if err := t.Init(config.Current.I18n, ""); err != nil {
		return err
	}

	db := mongo.NewMongo()
	db.Connect()
	if db.HasErrors() {
		return db.GetErrors()
	}

	defer db.Disconnect()

	return session.RunSession(db)
}

func Try() error {

	// oracle.Log()
	list := characters.List()

	log.ParseJson(list)

	return nil
}

/*
	BACKLOG Étendre la CLI avec des commande comme:

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
