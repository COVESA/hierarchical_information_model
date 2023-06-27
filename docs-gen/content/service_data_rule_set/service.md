---
title: "Service node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

Nodes of the type `service` must have the following mandatory metadata:
- Name
- Type
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

- A parent must be of the node type `branch`.
- Children, if any, must be of the node type `struct`.

An example of the specification of a `service` node is given below.
```YAML
NodeName:
  type: service
  description: This is a service node containing its mandatory metadata.
```

## Service State

A `service` node has a mandatory datatype of uint8. As it cannot be set to any other datatype than uint8,
the datatype metadata shall not be explicitly shown in the node.
