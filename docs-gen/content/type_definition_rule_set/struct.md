---
title: "Struct node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

A `struct` node type may be used to represent multiple data points that have the relationship that they are always sampled in an "atomic" operation,
i. e. the set of samples have the same timestamp.
An example of such a set of data points are the GPS dta points latitude, longitude, height.

This node type is only allowed in trees of the information type `typedefinition`.
For more information, please see respective parts of this document.

Nodes of the type `struct` must have the following mandatory metadata:
- Name
- Type
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

- A leaf node cannot have the `struct` node type.
- A parent must be of the node type `branch`.
- Children must be of the node type `property`.
- It must have at least one child.

An example of the specification of a `struct` node is given below.
```YAML
NodeName:
  type: struct
  description: This is a struct node containing its mandatory metadata.
```
