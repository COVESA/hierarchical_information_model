---
title: "Branch node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

The interior of a HIM tree is built up of nodes that have the node type `branch`.

Nodes of the type `branch` must have the following mandatory metadata:
- Name
- Type
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Comment
- Instances
- Aggregate

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

- The root node of a HIM tre must be of the node type `branch`.

- A leaf node cannot have the `branch` node type.

- If it has a parent, then it must be of the node type `branch`.

An example of the specification of a `branch` node is given below.
```YAML
NodeName:
  type: branch
  description: This is a branch node containing its mandatory metadata.
```
