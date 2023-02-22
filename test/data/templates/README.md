# Templates

These templates represent the execution chain.
Start with "base.tmpl" and work your way down the parse tree.

* Base > Package + Services + Enums + Messages
* Enums > Enum Attributes
* Services > RPCs + Enums + Messages
* Messages > Attributes + Messages + Enums