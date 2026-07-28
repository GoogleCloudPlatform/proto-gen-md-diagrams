# Release Notes

## July 28th, 2026 - Parser & Diagram Generation Improvements

- **Supported Output Format Documentation**: Illustrated both Default (HTML-Enhanced) and Pure Markdown (`-md`) modes alongside output Mermaid class diagrams in README.md.
- **JSON Name Annotations**: Added `(→ json_name)` alias rendering to Mermaid class diagram attributes and `JSON: <name>` labels to Markdown description tables when `json_name` is present.
- **Mermaid Class Diagram Generics**: Updated map and list collection type rendering to use valid Mermaid tilde syntax (`Map~K, V~` and `List~T~`).
- **Ordinal Field Ordering**: Enforced protobuf field ordinal ordering for class diagram attributes and markdown description tables.
- **Documentation & Test Sync**: Regenerated README.md diagrams and updated test fixtures to align with parser enhancements.

## Feb 8th - Bazel Module Additions

Reverted to standard Bazel build and added support
for the Bazel module. Since there are no guarantees that
the Bazel module will continue to work in it's existing form.
