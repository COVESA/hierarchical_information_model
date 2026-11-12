---
title: "Vehicle Data model"
date: 2019-08-04T13:05:11+02:00
weight: 1
---

The rule set for vehicle data is inherited from the [VSS rule set](https://covesa.github.io/vehicle_signal_specification/rule_set/).
It is the leaf nodes of a tree that represents and defines the actual vehicle data.
The definition is expressed by metadata describing the data samples associated to the node.

## Node Types

The node types for representing data entries are:
- Sensor: for data that is read-only and may have a dynamically changing value.
- Attribute: for data that is read-only and have a static value.
- Actuator: for data that is read-write.

Please see the respective chapters for more information about these node types.

## Vspec YAML extensions

The vspec format is based on YAML that is extended with two features described in respective chapters shown below.

* [Instances](/hierarchical_information_model/common_rule_set/vspec_extensions/instances)
* [Include](/hierarchical_information_model/common_rule_set/vspec_extensions/include)
