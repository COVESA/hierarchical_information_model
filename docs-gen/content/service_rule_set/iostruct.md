---
title: "IO-struct node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

An `iostruct` node type may be used to encapsulate multiple data points that have the relationship that they are all representing either
Input parameters or Output parameters of the microservice that they are linked to.

This node type is only allowed in trees of the information type `service`.
For more information, please see respective parts of this document.

An `iostruct` node must have one of the two names `Input` or `Output`.

Nodes of the type `iostruct` must have the following mandatory metadata:
- Name
- Type
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

- A leaf node cannot have the `iostruct` node type.
- A parent must be of the node type `procedure`.
- Children must be of the node types `property` or `symlink`.
- It must have at least one child.

An example of the specification of a `struct` node is given below.
```YAML
Input:
  type: iostruct
  description: This is an iostruct node representing Input parameters of a microservice.
```
