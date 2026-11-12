---
title: HIM Data Rule Set
weight: 20
chapter: true
---

# HIM Data Rule Set

The rules that are specific for the information type `Data` is described in this part of the document,
and together with the rules in the [Common rule set part](/hierarchical_information_model/common_rule_set) it forms the overall rules for this information type.

This information type is functionally equivalent with the `VehicleData` information type.
However, for use cases where the vehicle specific information type may appear as not fitting for the context this information type can be used.
This could e. g. be the case when data related to a driver, or passenger, of a vehicle is to be described.
Instead of using node types like `sensor` or `actuator` the node types `ro` and `rw` can be used.
