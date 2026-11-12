---
title: "Proxy node"
weight: 30
---

A `proxy` node in a HIM configuration file defines a tree that the server managing this forest does not have a local copy of.
It however knows how to forward requests to the server that manages the tree and has a local copy of it.
An example could be a home automation scenario, where a server deployed in a car knows the Local Area Network access point address bearer type,
so that when in proximity of the LAN, and after the vehicle has connected to it, requests received by the HIM server in the vehicle can then be
routed to the server deployed on the LAN.

Nodes of the type `proxy` must have the following mandatory metadata:
- Name
- Type
- Domain
- Version
- Local
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Public
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

This node type must have a node of type `branch` as parent, and must not have any children.

## Name
The name of this node is the root node name that a client shall use when addressing the tree, i. e. the first segment of the paths of its nodes.
It must be unique in the forest described in this file.
It does not have to be the same as the name of the root node of the tree file that it represents.
The server shall map the name of this node to the tree that it manages, and translate paths accordingly.
This makes it possible to use a different root node name than the one in a "standard" tree,
which then enables that multiple instances of the same standard tree can be present in the same forest.
A scenario where this may happen is when a truck has more than one trailer that each is represented by the same standard tree.
The trailers can then be distinguished by having different root node names in the configuration file.

## Domain
The `domain` metadata is what defines the domain and information type that the tree represents.
A domain name is made up of dot delimited segment names:
- The first segment (the left-most) represents a "top domain".
  - Examples of top domains can be Automotive, RoadInfrastructure, Aviation.
- Following segments, except the last, are sub domains.
  - Examples of sub domains can be Car, Truck, Trailer, ChargingStation.
  - The number of sub domains is not fixed, but it is recommended to stay below five.
  - Each additional sub domain should be a subset of the previous domain.
- The last segment (the right-most) represents the [information type](/hierarchical_information_model/).

## Version
The `version` shall be identical to the version of the tree where it shall be mandatory.
Versioning shall be based on the [Semantic versioning](https://semver.org/spec/v2.0.0.html) principles.

## Local
The `local` metadata shall be a URL, including schema (schema://url), that points either directly to the server that contains the tree, or to a proxy for it.
The schema selection is implementation specific, examples are “uds” if a Unix Domain socket is used, or “wss” if a secure WebSocket is used.
The `local` metadata shall not be supplied to a client requesting the configuration file.

## Public
The `public` metadata shall, if available, be a URL to a public copy of the tree managed by the server at the end-point.
It may be used to provision a client with a copy of the tree instead of e. g. requesting the (end-point) server to provision a copy.

## Example
An example of the specification of a `proxy` node is given below.
The `local` points to a Unix domain socket for communication with a proxy server in the vehicle that then communicates with the home over Zigbee.

```YAML
HIM.Home:
  type: proxy
  domain: HomeAutomation.Zigbee.Service
  version: X.Y.Z
  local: uds://path/to/uds-file
  description: Home automation proxy tree.
```
