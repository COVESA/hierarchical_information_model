---
title: "Attribute node"
date: 2019-08-04T12:37:31+02:00
weight: 40
---

Nodes of the types `attribute` must have the following mandatory metadata:
- Name
- Type
- Datatype
- Default
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Unit
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

This node type must have a node of type `branch` as parent, and must not have any children.

An example of the specification of an `attribute` node is given below.
```YAML
MaxPower:
  datatype: uint16
  type: attribute
  default: 0
  unit: kW
  description: Peak power, in kilowatts, that the engine can generate.
```
## Default
An attribute must define a `default` value, which is the persistent value assigned to the node.
Although categorized as persistent, it may be updated, but it should typically not change more than once per ignition cycle.

It is possible to set default values also to array elements. In this case square brackets shall be used.
The value for each element in the array must be specified. The size of the array is given by the number of elements specified within the square brackets.

Example 1: Empty Array

```YAML
  default: []
```

Example 2: Array with 3 elements, first element has value 1, second element value 2, third element value 0

```YAML
  default: [1, 2, 0]
```

Full example, array with two elements, first with value2, second with value 3:

```YAML
SeatPosCount:
  datatype: uint8[]
  type: attribute
  default: [2, 3]
  description: Number of seats across each row from the front to the rear
```

Using default values for structs is not allowed!
