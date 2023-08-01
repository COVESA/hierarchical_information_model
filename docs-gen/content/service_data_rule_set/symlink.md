---
title: "Symlink node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

Data represented by the node type `symlink` have a relationship to any other data having the same `iostruct` parent in that they are all representing either
Input parameters or Output parameters of the service that they are linked to.

An Otput parameter declared as a symlink will lead to that the returned value is read from that data point. 

An Input parameter declared as a symlink will lead to that the provided input value is written to that data point.

Nodes of the type `symlink` must have the following mandatory metadata:
- Name
- Type
- Path
- Domain
- Version
- Description

The Path must be a valid path to a leaf node in the tree declared by the Domain and Version metadata.

The Domain must declare an existing HIM tree of information type resourcedata.

The Version must be a dot delimited triplet declaring a version of the tree eclared by the Domain.
The triplet may a suffix of either "+" or "-", denoting that also any following or previous versions are also valid.

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

This node type must have a node of type `iostruct` as parent, and must not have any children.

An example of the specification of a `symlink` node is given below.

```YAML
VehicleService.GPS.GetPosition.Output.Latitude:
  type: symlink
  path: Vehicle.CurrentLocation.Latitude
  domain: Automotive.Vehicle.Car.ResourceData
  version: 3.0.0+
  description: This is a symlink to a VSS tree.
```
