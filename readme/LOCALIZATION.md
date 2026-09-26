# Localisation

La localisation est gérée avec **i18n** et ses fichiers de traduction.

## Organisation des fichiers

Les dossiers sont parcourus récursivement à partir de `Config.Dir`, afin de récupérer tous les fichiers présents sous ce répertoire avant de les filtrer.

Les fichiers peuvent être organisés dans autant de sous-dossiers que nécessaire, tant qu'il respecte à la fois le type et le nom attendus defini en `Config`.

Tout fichier ne correspondant pas à ces deux critères ne sera pas pris en compte.

## Localiser des entrées personnalisées

Si certains éléments créés en jeu nécessitent une localisation, il faudra veiller à respecter les étapes suivantes:

1. Noter les valeurs en jeu sous forme de clés i18n.
   - ⚠️ Il n'est pas encore possible d'éditer une valeur directement en jeu.
2. Mettre à jour les fichiers de localisation correspondants.

### Exemple

Pour un fichier de données `fantasy.yaml`.

#### **Fichier à localiser** : `data/template/hint/fantasy.yaml`

```yaml
name: 'hint.fantasy'
values:
  - hint.fantasy.dragon
  - hint.fantasy.grimoire
  - hint.fantasy.enchanter
  ...
```

Les valeurs du fichier sont directement des **clés i18n** plutôt que du texte à affiché à l'utilisateur.

#### **Dossier** : `locales/hint/fantasy/`

Les traductions peuvent être regroupées dans un dossier dédié, tant que les fichiers respectent les règles de nommage et de format attendues.

#### **Fichiers de traduction** 🇺🇸 : `locales/hint/fantasy/en-US.yaml`

```yaml
hint.fantasy:
  other: fantasy
hint.fantasy.dragon:
  other: dragon
hint.fantasy.grimoire:
  other: grimoire
hint.fantasy.enchanter:
  other: enchanter
```

#### **Fichiers de traduction** 🇫🇷 : `locales/hint/fantasy/fr-FR.yaml`

```yaml
hint.fantasy:
  other: fantasy
hint.fantasy.dragon:
  other: dragon
hint.fantasy.grimoire:
  other: grimoire
hint.fantasy.enchanter:
  other: enchanteur
```

## Configuration i18n

Actuellement, i18n n'est pas configurable, mais a pour objectif de le devenir à terme via la ligne de commande.

### Structure actuelle

```go
type Config struct {
	Default string   // Le code de la langue par défaut
	Dir     string   // Le dossier où aller chercher les traductions
	Format  Format   // JSON | YAML
	Locales []Locale // Les locales de chaque langue supportée
}

type Locale struct {
	Code string // Le code
	ISO  string // L'ISO
	Name string // Le nom de la langue dans celle-ci
	File string // Le nom de fichier attendu
}
```

### Configuration par défaut

```yaml
default: "fr"
dir: "locales"
format: "yaml"
locales:
  - code: en
    iso: "en-US"
    name: "English"
    file: "en-US.yaml"
  - code: fr
    iso: "fr-FR"
    name: "Français"
    file: "fr-FR.yaml"
```