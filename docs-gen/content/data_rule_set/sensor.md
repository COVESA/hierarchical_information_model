---
title: "Sensor node"
date: 2019-08-04T12:37:03+02:00
weight: 30
---

Sensors are signals to read values of properties in a vehicle. Values of sensors typically change over time. Reading a sensor shall return the current actual value of the related property, e.g. the current speed or the current position of the seat.

Nodes of the type `sensor` must have the following mandatory metadata:
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

An example of the specification of a `sensor` node is given below.

```YAML
Speed:
  type: sensor
  description: The vehicle speed.
  comment: For engine speed see Vehicle.Powertrain.CombustionEngine.Engine.Speed.
  datatype: float
  unit: km/h
  min: 0
  max: 300
```
