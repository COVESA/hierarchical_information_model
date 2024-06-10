---
title: "Data model"
date: 2019-08-04T13:05:11+02:00
weight: 1
---

The rule set for resource data is inherited from the [VSS rule set](https://covesa.github.io/vehicle_signal_specification/rule_set/).
It is the leaf nodes of a tree that represents and defines the actual data.
The definition is expressed by metadata describing the data associated to the node.

## Node Types

The node types for representing data entries are:
- Sensor
- Actuator
- Attribute

Please see the respective chapters for more information about these node types.

## Vspec YAML extensions

The vspec format is based on YAML that is extended with two features described in respective chapters shown below.

* [Instances](/hierarchical_information_model/resource_data_rule_set/vspec_extensions/instances)
* [Includes](/hierarchical_information_model/resource_data_rule_set/vspec_extensions/includes)
