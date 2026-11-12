---
title: "Procedure node"
weight: 30
---

The `procedure` node type must have the following mandatory metadata:
- Name
- Type
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Instances
- Comment
- NativeRate
- TimeToLive
- Deprecation

The Instances directive can be used when a service manages multiple resources,
see [Multiplexed microservice tree structure](/hierarchical_information_model/service_rule_set/basics#multiplexed-microservice-tree-structure).
Its definition follows the same rules as for its usage in branch type nodes as described [here](/hierarchical_information_model/common_rule_set/vspec_extensions/instances)
with the extension that it can be used in procedure nodes.

The NativeRate shall be an integer that indicates the typical duration in milliseconds between two subsequent microservice event messages.
If set to zero then the service terminates instantaneously.

The TimeToLive indicates a maximum time in milliseconds that the server shall allow a service to execute.
When it is exceeded the server may terminate the service and issue a final event message with the Status set to FAILED.

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

- A parent must be of the node type `branch`.
- The `procedure` node must have a child of node type `attribute` with the name Version.
- The `procedure` node may have children of the node type `iostruct` with Input or Output as the node name.

An example of the specification of a `procedure` node is given below.
```YAML
microserviceName:
  type: procedure
  instances: Resource[1, 3]
  description: This is an example of a microservice procedure node that manages three resources.
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
