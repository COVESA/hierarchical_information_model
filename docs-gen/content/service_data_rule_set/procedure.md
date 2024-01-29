---
title: "procedure node"
weight: 30
---

The `procedure` node type must have the following mandatory metadata:
- Name
- Type
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

- A parent must be of the node type `branch`.
- Children, if any, must be of the node type `iostruct`.

An example of the specification of a `procedure` node is given below.
```YAML
microserviceName:
  type: procedure
  description: This is an example of a microservice procedure node containing its mandatory metadata.
```
