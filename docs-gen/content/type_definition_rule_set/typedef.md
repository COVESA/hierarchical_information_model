---
title: "Typedef node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

The node type `typedef` is used to define node metadata that can be referenced from nodes declared in trees of other information types.
This enables one definition to be shared between multiple nodes in other trees.

Nodes of the type `typedef` must have the following mandatory metadata:
- Name
- Type
- Datatype
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used:
- Unit
- Min
- Max
- Default
- Allowed/Enum
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

## Typedef usage rule set
The following rules define how typedef nodes shall be used.
* The effective set of metadata of a node that refers to a typedef is:
  * Name: Referring node
  * Type: Referring node
  * Datatype: Referred node
  * Description: Referring node
  * Any other metadata: If exists in Referring node then Referring node else Referred node.
* VSS-tools exporters shall in the exported format synthesize a node containing the effective set of metadata as defined in the typedef rule set.

## Example using typedef
In the example a vspec file describing a tree of the VehicleData information type contains a node that refers to a typedef node containing an enum definition.
The node in the referring tree:
```
Vehicle.X.Y:
  type: sensor
  datatype: Types.Enum.Example
  description: A sensor node in the Vehicle tree.
```
The node in the reffered (ypedef) tree:
```
Types.Enum.Example:
  description: Enum example type definition.
  type: typedef
  datatype: uint8
  enum:
    A: 0
    B: 1
    C: 2
    D: 3
```
The resulting node after exporting the vspec declarations to YAML:
```
Vehicle.X.Y:
  type: sensor
  datatype: uint8
  enum:
    A: 0
    B: 1
    C: 2
    D: 3
  description: A sensor node in the Vehicle tree.
```
