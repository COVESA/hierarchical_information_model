---
title: "Actuator node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

Actuators are used to control the desired value of a property. Some properties in a vehicle cannot change instantly. A typical example is position of a seat or a window. Reading a value of an actuator shall return the current actual value, e.g. the current position of the seat, rather than the wanted/desired position. A typical example could be if someone wants to change the position of a seat from 0 to 100. This can be changed by setting the corresponding actuator to 100. If the actuator is read directly after the set request it will still return 0 as it might take some seconds before the seat reaches the wanted position of 100. If the seat by some reason is blocked or cannot be moved due to safety reasons it might never reach the wanted position. It is up to the vehicle to decide how long time it shall try to reach the desired value and what to do if it needs to give up.

Nodes of the type `actuator` must have the following mandatory metadata:
- Name
- Type
- Datatype
- Description

For more information, see the [Common Rule Set: Mandatory Metadata](/hierarchical_information_model/common_rule_set/basics#mandatory-metadata).

Besides the mandatory metadata mentioned above, the following optional metadata may be used
- Unit
- Min
- Max
- Allowed
- Comment

For more information, please see the [Common Rule Set: Optional Metadata](/hierarchical_information_model/common_rule_set/basics#optional-metadata).

This node type must have a node of type `branch` as parent, and must not have any children.

An example of the specification of a `actuator` node is given below.

```YAML
TripMeterReading:
  datatype: float
  type: actuator
  unit: km
  description: Trip meter reading.
  comment: The trip meter is an odometer that can be manually reset by the driver
```
