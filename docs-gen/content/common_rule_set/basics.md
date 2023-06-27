---
title: "Basic Rules"
date: 2019-08-04T13:05:11+02:00
weight: 1
---
## Specification format

The syntax used to define the content of a HIM taxonomy is based on YAML and must comply to the 
[YAML specification](https://yaml.org/).

A HIM taxonomy may contain subtrees that are identical copies of each other, for example descriptions of doors in a vehicle.
In order not to have to repeat the content of the subtree multiple times, the subtree can be defined in a separate file, 
which can then be included into the main tree by the directive
```
#include <filename> [prefix]
```
The optional [prefix] specifies a branch name to be prepended to all node names in the included file. 
The include directive may be used in files that are themselves included.
Please note that, from a YAML perspective, the include directive is just another comment.

A tool is neeed to generate the expanded version of the tree that contains instances of the included subtree,
such as the tools developed by the VSS community - [COVESA/VSS-Tools](https://github.com/COVESA/vss-tools).

The files defining a HIM taxonomy may have the file extension '.him', but the project defining the domain taxonomy may define and use a file extension of their own choice.
The format of the files may be transformed from YAML to another format like JSON, in which case it may use a file extension of that format.
This also pplies to the file extension of the HIM configuration file.

HIM is in itself case sensitive.
This means that keywords, signal names, types and values normally shall be given with the case specified.
It is however recommended not to take advantage of this and reuse the same name with different case,
as some implementations may treat VSS identifiers as case insensitive.

## Mandatory metadata

It is mandatory for all HIM nodes to contain the following metadata:
- Name
- Type
- Description

Some HIM node types also require the following metadata:
- Datatype
- Default

### Node name

A HIM node shall have the syntax of a YAML mapping block with the key name of the mapping block representing the node name as shown below.
```
NodeName:
  type: x
  description: abc
```

A node name is recommeded to be unique with the scope of the tree it is used in. The name shall follow the [Naming Conventions](#naming-conventions)

The qualified node name, which must be unique within the scope of the tree, are defined left-to-right of the concatenation of the node names, 
starting from the root node of the tree and traversing the tree nodes to the node in focus, with a period (.) as delimiter between node names.
An example of a qualified node name, which is called a path, is
```
a.b.c
```
where a is the name of the root node, b is a child of a, 
and the node in focus is c, a child of b.

### Node types

The node types described here are used by more than one information type,
while the node types specifi to a single information type are found in respective parts of this document.
- [Branch](/hierarchical_information_model/common_rule_set/node_types/branch)
- [Struct](/hierarchical_information_model/common_rule_set/node_types/struct)
- [Property](/hierarchical_information_model/common_rule_set/node_types/property)

#### Struct node type

The 'struct' node type is used to define a structure data type that groups a number of fields. 
An example is shown below.
```YAML
NodeName:
  type: struct
  description: This is a struct node containing its mandatory metadata.
```
This node type must have at least one child of node type property, and must not have children of any other node type. 

#### Property Node Type

The 'property' node type is used to define the fields of a struct.
An example is shown below.
```YAML
NodeName:
  type: property
  datatype: "any supported datatype, including a reference to a struct definition"
  description: This is a property node containing its mandatory metadata.
```
The 'datatype' metadata is mandatory for a property node. It may also have optional metadata such as unit, min, max, and allowed. 

A property node can only have a struct node as its parent, and must not have any children.

### Description

Describes the meaning and content of the node. Recommended to start with a capital letter and end with a dot (.).

## Optional Metadata

Different information types may then specify further optional metadata that may be used by their supported node types:
1. Unit
2. Min
3. Max
4. Allowed
5. Comment
6. Instances
7. Aggregate
8. Deprecation

For more information on which information/node types this metadata can be used,
please see the rule set documentation for these.

For oversight a few general rules are shown here:
- The entries 1 through 4 cannot be used in `branch` or `service` node types.
- The entries 6 and 7 can only be used in `branch` type nodes.

### Datatype
See the [Datatypes chapter](/hierarchical_information_model/common_rule_set/data_entry/datatypes)

### Unit
See the [Units chapter](/hierarchical_information_model/common_rule_set/data_entry/units)

### Min/Max/Allowed
See the [Value restrictions chapter](/hierarchical_information_model/common_rule_set/data_entry/value_restrictions)

### Comment

A comment can be used to provide additional informal information on a branch. This could include background information on the rationale for the branch, references to related branches, standards and similar. Recommended to start with a capital letter and end with a dot (.).

### Instances
See the [Instances chapter](/hierarchical_information_model/common_rule_set/instances).
The `instances` metadata is only allowed in `branch` type nodes.

### Aggregate
Defines whether or not this branch is an aggregate.
If not defined, this defaults to ```false```.
An aggregate is a collection of data nodes that make sense to handle together in a system.
A typical example could be GNSS location, where latitude and longitude make sense to read
and write together. This is supposed to be deployment specific,
and for that reason it is recommendedd not to be defined in a standard HIM tree,
but rather be added in a deployment preparation of he tree.
For branches that both have `instances` defined and `aggregate: true`, then aggregate refers to the signals for
individual instances, i.e. signals for different instances can be handled separately.
The `aggregate` metadata is only allowed in `branch` type nodes.

### Deprecation

During the development of a taxonomy nodes might be moved or deleted. Giving users of the taxonomy a chance to adopt to the
changes, the original nodes are marked as deprecated with the following rules.

* Nodes, which are moved in the tree or are intended to be removed from the specification are marked with the deprecation keyword.
* The string following the deprecation keyword shall start with the version, when the node was deprecated starting with `V` (e.g. `V2.1`) followed by the reason for deprecation.
* If the node was moved, it shall be indicated by `moved to` followed by the new node name in dot notation as deprecation reason.
This keyword shall be used only if the meta-data of the moved node hasn't changed.
* If the node is intended to be removed from the specification or the meta data changed, it shall be indicated by `removed` and optionally the reason for the removal as deprecation reason.
* Nodes which are deprecated will be removed from the specification, either in the second minor update or, if earlier, the next major update.

### Example
```YAML
Vehicle.Navigation.CurrentLocation:
  type: branch
  description: The current latitude and longitude of the vehicle.
  deprecation: V2.1 moved to Vehicle.CurrentLocation
```

It is recommended for servers, which are implementing protocols for the vehicle signal specification, to serve old and new nodes during the deprecation period described above.

## Style Guide

The HIM specification must adhere to YAML syntax. To keep different domain taxonomy specifications consistent the following style guide is provided.

### Naming Conventions

The recommended naming convention for node elements is to use camel case notation starting with a capital letter. It is recommended to use only
`A-Z`, `a-z` and `0-9` in node names. For boolean signals it is recommended to start the name with `Is`.

Examples:

```
SomeBranch.AnotherBranch.MySignalName
```
Naming convention for string literals can be found in the [chapter](/hierarchical_information_model/common_rule_set/data_entry/value_restrictions) for specifying allowed values.

### Line Length

It is recommended that line length in files containing HIM information shall not exceed 120 characters.
This is not a strict limit, it is e.g. not recommended to split long URLs present in files over multiple lines.
