---
title: "Service Data Model"
date: 2019-08-04T13:05:11+02:00
weight: 1
---

## HIM Service-Data Tree
A HIM service data tree contains a taxonomy of microservice declarations.
The tree structure is used to create a sets of microservices.
Such as set is named `Service`, or `Service Group` for larger sets where a set of Service Groups would then form a Service,
see the figure below, or the [example service data tree](https://github.com/COVESA/hierarchical_information_model/blob/master/examples/HIM_Service.v1.0.0.him).

![HIM service tree structure](/hierarchical_information_model/images/service_tree_structure.png?width=50pc)
*Figure x. HIM service tree structure

## HIM Microservice Data
A microservice is in HIM defined by a procedure signature as shown below:

**(output1,..outputN) microServiceName(input1,..inputM)**

where
- microServiceName is the name of the procedure.
- input1,..,inputM denotes the input parameters of the procedure, which may be zero or more.
- output1,..,outputN denotes the output parameters of the procedure, which may be zero or more.

HIM only specifies the data defining the above procedure signature, it does not specify details of how a call is made, such as how it is serialized, etc.
This is expeted to be defined in an interface specification that uses HIM.

## Microservice tree structure

The general structure of how a microservice is represented in a HIM tree is shown in the figure below.

![HIM microservice tree structure](/hierarchical_information_model/images/microservice_tree_structure.png?width=50pc)
*Figure x. HIM microservice tree structure

- The name of the procedure is the name of the node of the type `procedure`.
- The input and output parameters are respectively represented by a  node of type `iostruct` that must have the names 'Input' and 'Output', respectively.
- An input/output parameter is represented by a node of type `property` or `symlink`.

Any Input/Output nodes, and their associated children, are only present if the procedure has at least one parameter of the respective Input/Output.

## Microservice completion state

A microservice may have a significant temporal duration from it is started and until it completes.
It may therefore be desireable that the state of the microservice can be observed,
as e. g. it may be the case that an ongoing microservice execution does not allow new microservice actuations to be started.
If an observable microservice completion state is desired, then the designer of this HIM service domain must create such an output parameter of the microservice.
The characteristics of this output parameter can be freely chosen, but below is a proposal that if widely used may help to improve interoperability.

The parameter shall be named Completion.

The completion state of the HIM microservice is represented by a uint8 datatype value.

The allowed values of the microservice state are:
- 0-99 : ongoing
- 100 : ready

The rules for how a microservice shall update the state value follows below:

- A fully functioning microservice that is not ongoing shall have the value 100.
- When a valid microservice request is received the state shall be set to zero (0).
- The microservice may update the state value as the microservice execution proceeds.
- When the microservice execution successfully terminates the state value must be set to 100.

If there is a need for microservice specific error codes then these should be defined as another output parameter.
