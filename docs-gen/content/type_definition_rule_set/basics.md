---
title: "Type Definition Model"
weight: 1
---

A `TypeDef` tree contains data type definitions.
These can either be complex datatypes that are not part of the
[primitive datatypes](/hierarchical_information_model/common_rule_set/data_entry/datatypes#primitive-datatypes) defined by HIM,
or it can be common allowed definitions.

If a tree with data or service information contains links to a TypeDef tree then this tree must be available to the server managing the data/service tree.
The server shall use the type definition information to verify that client requests are valid in the case of links as mentioned.

## Root node name
The name of the root node of the type definitions tree shall be "Types".\
If an implementation of a server ecosystem selects to include multiple type definition trees in the defined forest then
[the root node name that the HIM configuration file](/hierarchical_information_model/configuration_rule_set/node_types/direct/#name) maps to these trees shall be "TypesX"
where X can be any sequence of readable characters, with a length from one and up.
X can be a number or more descriptive such as Car, Trailer, etc.\
If the HIM configuration file only contains one type definition tree then the mapped name shall be "Types".

## Node Types
The currently supported node types can be used for either the definition of [structs](https://en.wikipedia.org/wiki/Composite_data_type),
or for the definition of [allowed](/hierarchical_information_model/common_rule_set/data_entry/value_restrictions/#allowed) value restrictions.

The node types for representing type definitions are:
- Branch
- [Struct](/hierarchical_information_model/type_definition_rule_set/struct)
- [Property](/hierarchical_information_model/type_definition_rule_set/property)
- [Typedef](/hierarchical_information_model/type_definition_rule_set/typedef)
- Attribute: only used in the [tree identification](/hierarchical_information_model/common_rule_set/tree_identification) nodes

Please see more information about these node types [here](/hierarchical_information_model/common_rule_set/node_types/).
