---
title: "Struct node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

The node type `struct` is only allowed in trees of the information types `typedefinition` or `service`.
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
- A parent must be of the node type `branch` or `service`.
- Children must be of the node type `property`.
- It must have at least one child.

An example of the specification of a `struct` node is given below.
```YAML
NodeName:
  type: struct
  description: This is a struct node containing its mandatory metadata.
```
