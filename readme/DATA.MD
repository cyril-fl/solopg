# Data

**Solog** est conçu comme un moteur extensible de JDR solo. Le dossier **data** contient les données du jeu, stockées au format `.yaml` afin d’être facilement éditables, remplaçables et surtout extensibles.

Actuellement aucun moyen de charger un set de règles différent afin de passer d’un univers à l’autre au sein du même moteur n’a encore été prévu.

## Organisation
### Système

Cette partie contient les fichiers nécessaires au fonctionnement du jeu. Ils définissent les règles de l’univers, les classes, les races, les jets de dés et d’autres éléments. C’est ici qu’est défini si le monde est un monde de fantasy, science-fiction ou réaliste.

Bien que **solog** soit conçu comme un moteur et que la majorité des fichiers soient optionnels, les fichiers ou dossiers situés aux chemins suivants sont obligatoires. Leur contenu reste toutefois modifiable, et leurs chemins peuvent être configurés dans `config.yaml` dans la section *documents*.

#### Fichiers
- `dice`: `"data/systems/rules/dices/dice.yaml"`
- `objects_category`: `"data/systems/attributes/category.yaml"`
- `classes`: `"data/systems/attributes/class.yaml"`
- `races`: `"data/systems/attributes/races.yaml"`
- `rarity`: `"data/systems/attributes/rarirty.yaml"`
- `slots`: `"data/systems/attributes/slots.yaml"`
- `stats`: `"data/systems/attributes/stats.yaml"`
- `variety`: `"data/systems/attributes/variety.yaml"`

#### Dossiers
- `oracle`: `"data/systems/rules/oracles"`
- `encounters`: `"data/systems/rules/oracles/encounter"`
- `characters`: `"data/template/characters"`
- `locations`: `"data/template/locations"`

```
⚠️⚠️⚠️ Attention : étant en cours de développement, cette liste peut être devenue obsolète. Les données à jour sont situées dans `config.yaml`.
```

### Templates

Contient des éléments préchargés, comme des personnes, des objets, des quêtes, etc... bien que entièrement optionnels.

Ces éléments pourrait également être créés en jeu et enregistrés en base de données mais *les templates* permettent actuellement une gestion plus fine, en ajoutant par exemple des équipements à un PNJ.

```
⚠️⚠️⚠️ Attention : ces réglages sont actuellement globaux. Ainsi, un personnage nommé Jean ou une classe Orc définis dans les données seront communs à toutes vos parties.
```
