---
title: "Property node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

Data represented by the node type `property` have a relationship to any other data having the same `struct` parent in that they are all sampled in an "atomic" operation,
i.e. the set of data points have the same timestamp representing the sample time.

Nodes of the type `property` must have the following mandatory metadata:
- Name
- Type
- Datatype
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Unit
- Min
- Max
- Allowed
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

This node type must have a node of type `struct` as parent, and must not have any children.

An example of the specification of a `property` node is given below.

```YAML
Type.OpenHours.Open:
  datatype: uint8
  type: property
  max: 24
  description: Time the address opens
```
