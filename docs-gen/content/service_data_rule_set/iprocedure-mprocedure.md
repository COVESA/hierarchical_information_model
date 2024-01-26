---
title: "iprocedure and mprocedure nodes"
weight: 30
---

As described in the [Service Data Model](https://covesa.github.io/hierarchical_information_model/service_data_rule_set/basics#node-types-iprocedure-and-mprocedure) chapter
there are two types of nodes that can be used to represent microservices.

Nodes of the types `iprocedure` or `mprocedure` must have the following mandatory metadata:
- Name
- Type
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

- A parent must be of the node type `branch`.
- Children, if any, must be of the node type `iostruct`.

An example of the specification of a `iprocedure` node is given below. Please observe that the microserviceName should be prefixed with the letter i.
```YAML
imicroserviceName:
  type: iprocedure
  description: This is an example of a microservice iprocedure node containing its mandatory metadata.
```
Changing the type to `mprocedure` and the microserviceName prefix to the letter m would then represent the monitor flavor of the microservice specification.
