---
title: "Type Definition Model"
weight: 1
---

A `type_definition` tree contains data type definitions.
These can either be complex datatypes that are not part of the
[primitive datatypes](/hierarchical_information_model/common_rule_set/data_entry/datatypes#primitive-datatypes) defined by HIM,
or it can be common allowed definitions.

If a tree with data or service information contains links to a typedefinition tree then this tree must be available to the server managing the data/service tree.
The server shall use the type definition information to verify that client requests are valid in the case of links as mentioned.

## Root node name
The name of the root node of the type definitions tree shall be "TypesX" where X can be any sequence of readable characters, with a length from zero and up.
With an empty X the name will be "Types" and thus conformant with the COVESA VSS rule set.
If an implementation of a server ecosystem selects to use multiple type definition trees then X must be used to distinguish between the different trees.
X can then be a number or more descriptive such as Car, Trailer, etc.

## Node Types
The currently supported node types can be used for either the definition of [structs](https://en.wikipedia.org/wiki/Composite_data_type),
or for the definition of [allowed](/hierarchical_information_model/common_rule_set/data_entry/value_restrictions/#allowed) value restrictions.

The node types for representing type definitions are:
- Branch
- [Struct](/hierarchical_information_model/type_definition_rule_set/struct)
- [Property](/hierarchical_information_model/type_definition_rule_set/property)
- [Typedef](/hierarchical_information_model/type_definition_rule_set/typedef)
- Attribute // must only be used in tree version nodes

Please see more information about these node types [here](/hierarchical_information_model/common_rule_set/node_types/).
