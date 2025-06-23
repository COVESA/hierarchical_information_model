---
title: "Service Data Model"
date: 2019-08-04T13:05:11+02:00
weight: 1
---

## HIM Service-Data Tree
A HIM service data tree contains a taxonomy of microservice declarations.
The tree structure is used to create a sets of microservices.
Such as set is named `Service`, or `Service Group` for larger sets where a set of Service Groups would then form a Service,
see the figure below, or the [example service data tree](https://github.com/COVESA/hierarchical_information_model/blob/master/examples/HIM_Service_example.v1.0.0.yaml).

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

Service response or event messages shall therefore contain a parameter called Status, an enum with the following definition:
```
enum {
	ONGOING = 1     // in execution of latest call
	SUCCESSFUL = 0  // terminated successfully in latest call
	FAILED = -1      // terminated due to failure in latest call
}
```

The rules for how a microservice shall update the state value follows below:

- A fully functioning microservice that is not ongoing shall have the value SUCCESSFUL.
- When a valid microservice request is received and the service execution is started the state shall be set to ONGOING.
- When the microservice execution successfully terminates the state value shall be set to SUCCESSFUL.
- If the microservice fails at any point during its execution the state value shall be set to FAILED.

If there is a need for microservice specific error codes then these should be defined as another output parameter.
A microservice may include in other output parameters that carry information about its execution state.

## Service group common properties
A service group may have common configuration data, here referred to as properties.
This could for example be the seating locations, i. e. the location names for the seats in the vehicle.
Another example is the axle/wheel configuration of the vehicle, information about the number of axles and the number of wheels on each.
If a service group shares common property data then the service group branch shall have a `branch` type child node with the name Properties.
This node shall then have children nodes of `attribute` type with default values representing the configuration data.

An example of a Properties node and its SeatingLayout child is given below.
```YAML
Properties:
  type: branch
  description: This is an example of a service group Properties node.
```

```YAML
SeatingLayout:
  type: attribute
  datatype: string
  default: [{"Row1": {"Left", "Right"}}, {"Row2": {""Left", "Middle", "Right""}}]
  description: This is an example of the configuration of a seating layout.
```

