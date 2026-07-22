## Outil CLI Go — validation de CHANGELOG.md

Vérifie, avant merge, qu'un `CHANGELOG.md` (généré par Changesets dans `npm-packages`) respecte le format attendu :
- le premier en-tête de version correspond au champ `version` de `package.json`
- présence de `# Changelog`, d'un en-tête `## x.y.z`, d'une section `### Patch Changes` (ou Minor/Major), et d'au moins une entrée `- ` dessous

Ne valide plus la présence d'un lien PR en bas de page (règle retirée). Code dans `main.go`, tests dans `main_test.go`.

## 🚀 RTK - Rust Token Killer (Optimized)
All shell commands (`git`, `npm`, `jest`, etc.) are automatically proxied via `rtk` for 80% token savings.
- **Direct Usage:** `rtk gain` (analytics), `rtk discover` (missed savings).
- **Files:** Use `rtk read <file>`, `rtk ls`, `rtk find`, `rtk grep` for compressed agent output.
