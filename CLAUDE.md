## Outil CLI Go — validation de CHANGELOG.md

Vérifie, avant merge, qu'un `CHANGELOG.md` (généré par Changesets dans `npm-packages`) respecte le format attendu :
- le premier en-tête de version correspond au champ `version` de `package.json`
- présence de `# Changelog`, d'un en-tête `## x.y.z`, d'une section `### Patch Changes` (ou Minor/Major), et d'au moins une entrée `- ` dessous

Ne valide plus la présence d'un lien PR en bas de page (règle retirée). Code dans `main.go`, tests dans `main_test.go`.

## 🚀 RTK - Rust Token Killer (Optimized)
All shell commands (`git`, `npm`, `jest`, etc.) are automatically proxied via `rtk` for 80% token savings.
- **Direct Usage:** `rtk gain` (analytics), `rtk discover` (missed savings).
- **Files:** Use `rtk read <file>`, `rtk ls`, `rtk find`, `rtk grep` for compressed agent output.

## graphify

This project has a knowledge graph at graphify-out/ with god nodes, community structure, and cross-file relationships.

Rules:
- For codebase questions, first run `graphify query "<question>"` when graphify-out/graph.json exists. Use `graphify path "<A>" "<B>"` for relationships and `graphify explain "<concept>"` for focused concepts. These return a scoped subgraph, usually much smaller than GRAPH_REPORT.md or raw grep output.
- If graphify-out/wiki/index.md exists, use it for broad navigation instead of raw source browsing.
- Read graphify-out/GRAPH_REPORT.md only for broad architecture review or when query/path/explain do not surface enough context.
- After modifying code, run `graphify update .` to keep the graph current (AST-only, no API cost).
