package config

import (
	"fmt"
	"solopg/app/services/t"
	"solopg/app/services/yaml"
)

type config struct {
	Name      string            `yaml:"name"`
	Verbose   bool              `yaml:"verbose"`
	I18n      t.Config          `yaml:"i18n"`
	Commands  commands          `yaml:"commands"`
	Documents structureDocument `yaml:"documents"`
}

type commands struct {
	Config  commandParams `yaml:"config"`
	Root    commandParams `yaml:"root"`
	Run     commandParams `yaml:"run"`
	Try     commandParams `yaml:"try"`
	Version commandParams `yaml:"version"`
}

type commandParams struct {
	Use     string `yaml:"use"`
	Short   string `yaml:"short"`
	Long    string `yaml:"long"`
	Example string `yaml:"example"`
	Args    []arg  `yaml:"args"`
}

type structureDocument struct {
	Folders structureFolders `yaml:"folders"`
	Files   structureFiles   `yaml:"files"`
}

type structureFolders struct {
	Characters string `yaml:"characters"`
	Locations  string `yaml:"locations"`
	Oracle     string `yaml:"oracle"`
	Encounters string `yaml:"encounters"`
}

type structureFiles struct {
	Classes         string `yaml:"classes"`
	Dice            string `yaml:"dice"`
	ObjectsCategory string `yaml:"objects_category"`
	Races           string `yaml:"races"`
	Rarity          string `yaml:"rarity"`
	Slots           string `yaml:"slots"`
	Stats           string `yaml:"stats"`
	Variety         string `yaml:"variety"`
}

type arg struct {
	Name     string `yaml:"name"`
	Required bool   `yaml:"required"`
}

var Current config

func init() {
	loaded, err := Load()
	if err != nil {
		panic(err)
	}
	Current = *loaded
}

func Load() (*config, error) {
	current, err := yaml.LoadFromFile[config]("config/.config.yaml")
	if err != nil {
		return nil, err
	}

	if err := current.I18n.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return current, nil
}

/*
TODO
gerer le cas post buil avec un truc genre:
	getPath -> if isDev = path actuell
		else constructeur de path
aussi faire en sorte de embeded config.yaml part defaut
et de laisser un chemin pour configurer avec un yaml externe.
creer une commande genre
	CLI config --init
et si je fait
	CLI config --set [...args]
je passe automatiquement par l'equivalent de
	isInit ? -> si oui rien, puis config
		-> si non init, puis config

permtet de set les config direct une a une ou en batch.

a voir si la config est editable en fichier .yaml ou embarque é editable avec Vim ect.
je crois que ce serais le mieux

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
