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
The name of the root node of the type definitions tree shall be "Types".

## Node Types
The currently supported node types can be used for either the definition of [structs](https://en.wikipedia.org/wiki/Composite_data_type),
or for the definition of [allowed](/hierarchical_information_model/common_rule_set/data_entry/value_restrictions/#allowed) value restrictions.

The node types for representing type definitions are:
- Branch
- Struct
- Property
- Attribute

Please see more information about these node types [here](/hierarchical_information_model/common_rule_set/node_types/).

## Attribute node used for allowed definition
In this usage the attribute node type must have a node of type `branch` as parent, and must not have any children.
When used for allowed definition, the metadata types Unit, Min, Max or Default must not be used.

An example of the specification of a `attribute` node for an allowed definition is given below.

```YAML
Types.Cabin.DriverPositionValues:
  type: attribute
  datatype: string
  allowed: ['LEFT', 'MIDDLE', 'RIGHT']
  description: DriverPosition allowed values.
```
