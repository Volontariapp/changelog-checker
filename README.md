# Changelog Checker

Un outil CLI léger et rapide écrit en Go pour s'assurer que les fichiers `CHANGELOG.md` respectent les standards de Volontariapp.

## Objectif

Cet outil vérifie la conformité du changelog avant une fusion (merge) :
1. **Correspondance de Version :** Le premier en-tête de version du changelog doit correspondre exactement au champ `version` du fichier `package.json`.
2. **Structure du Changelog :** `CHANGELOG.md` doit contenir :
 - `# Changelog`
 - un en-tête de version au format `## x.y.z`
 - `### Patch Changes` (ou Minor/Major)
 - au moins un élément de liste sous cette section (ligne commençant par `- `)

Il **ne valide plus** la présence d'un lien PR en bas de page.

Exemple attendu :

```markdown
# Changelog

## 0.7.1

### Patch Changes

- [#25](https://github.com/Volontariapp/npm-packages/pull/25) [`fb3b3cf`](https://github.com/Volontariapp/npm-packages/commit/fb3b3cf44270654a85ff24517585f888bb78ea48) Merci [@Fyorix](https://github.com/Fyorix)! - ajout d'un nouveau test pour le contrat du package

- test des contrats patch
```

## Utilisation

Vous pouvez lancer le vérificateur en fournissant les chemins vers les fichiers `package.json` et `CHANGELOG.md`.

```bash
# Utilisation par défaut (suppose que les fichiers sont dans le dossier courant)
./changelog-checker

# Spécifier des chemins personnalisés
./changelog-checker chemin/vers/package.json chemin/vers/CHANGELOG.md
```

**Codes de sortie :**
- `0`: Validation réussie.
- `1`: Validation échouée (incohérence de version, format invalide, ou fichiers manquants).

## Développement

### Prérequis
- Go 1.22+

### Lancer les Tests
Une suite complète de tests est incluse pour couvrir les cas particuliers (fichiers valides, sections manquantes, entrées de liste manquantes, et incohérences de version).
```bash
go test -v ./...
```

### Compiler le Binaire
```bash
go build -o changelog-checker main.go
```
