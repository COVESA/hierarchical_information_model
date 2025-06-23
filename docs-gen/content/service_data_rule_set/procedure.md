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
- The `procedure` node must have a child of node type `attribute` with the name Version.
- The `procedure` node may have children of the node type `iostruct` with Input or Output as the node name.

An example of the specification of a `procedure` node is given below.
```YAML
microserviceName:
  type: procedure
  description: This is an example of a microservice procedure node containing its mandatory metadata.
```

## Version node
The version node contains the version of the procedure.
The node type shall be `attribute`, and the datatype shall be `string`.
The default value shall follow the Semantic version format with Major, Minor, and Patch dot delimited values.

An example of a Version node is given below.
```YAML
Version:
  type: attribute
  datatype: string
  default: 0.0.1
  description: This is an example of a microservice Version node containing its mandatory metadata.
```
