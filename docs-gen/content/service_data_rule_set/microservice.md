---
title: "Microservice node"
weight: 30
---

Nodes of the type `microservice` must have the following mandatory metadata:
- Name
- Type
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

- A parent must be of the node type `branch`.
- Children, if any, must be of the node type `iostruct`.

An example of the specification of a `microservice` node is given below.
```YAML
microserviceName:
  type: microservice
  description: This is a microservice node containing its mandatory metadata.
```

## Microservice State

A `microservice` node has a mandatory datatype of uint8. As it cannot be set to any other datatype than uint8,
the datatype metadata shall not be explicitly shown in the node.
The main purpose of the state information is to provide a "percentage" state of the execution to clients.

The microservice must set i to zero (0) at the beginning of the execution, and it must set it to hundred (100) at the end of a successful execution.
It may further update it one or more times during the execution.

If an error occurs during execution the state shall be set to one of the predefined error values.
This error value is not replacing either a microservice specific error value which the microservice may return as an Output parameter,
nor with the error value that a transport protocol may report.
