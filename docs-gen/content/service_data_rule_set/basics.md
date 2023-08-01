---
title: "Service Data Model"
date: 2019-08-04T13:05:11+02:00
weight: 1
---

## HIM Service Data
A service is in HIM defined by a procedure signature as shown below:

**(output1,..outputN) serviceName(input1,..inputM)**

where
- serviceName is the name of the procedure.
- input1,..,inputM denotes the input parameters of the procedure, which may be zero or more.
- output1,..,outputN denotes the output parameters of the procedure, which may be zero or more.

HIM only specifies the data defining the above procedure signature, it does not specify details of how a call is made, such as how it is serialized, etc.

## Service data tree structure

The general structure of how service data is represented in a HIM tree is shown in the figure below.

![HIM service data graph structure](/hierarchical_information_model/images/service_graph_representation.png?width=50pc)
*Figure x. HIM service data graph representation

- The name of the procedure is the name of the node of `service` type.
- The input and output parameters are respectively represented by a  node of type `iostruct` that must have the names 'Input' and 'Output', respectively.
- An input/output parameter is represented by a node of type `property` or `symlink`.

Any Input/Output nodes, and their associated children, are only present if the procedure has at least one parameter of the respective Input/Output.

## Service state

A service may have a temporal duration from it is started and until it completes.
It is therefore desireable that the state of the service can be observed,
as it may be the case that an ongoing service execution does not allow new service actuations to be started.
The state of a HIM service is represented by a uint8 datatype value that all `service` nodes have as a mandatory metadata.

The allowed values of the service state are:
- 0-99 : ongoing
- 100 : ready
- 101 : unavailable
- 102 : broken

The rules for how a service shall update the state value follows below:

- A fully functioning that is not ongoing shall have the value 100.
- When a valid service request is received the state shall be set to zero (0).
- The service may update the state value as the service execution proceeds.
- When the service execution successfully terminates the state value mus be set to 100.
- If a service is functional but temporarily unavailable it shall be set to 101.
- If a service requires repair or other type of critical maintenance to become functioning it shall be set to 102.

The service state does not represent the error code that is typically part of a transport protocol.
