---
title: "Hierarchical Information Model"
---
# Hierarchical Information Model

The Hierarchical Information Model (HIM) is an initiative by [COVESA](https://www.covesa.global/) to define a syntax for how to define taxonomies containing different types of information.
The documentation, source code and releases can be found in the [HIM github repository](https://github.com/COVESA/hierarchical_information_model).

The information of a domain that is described using HIM is represented in a graph made up of a tree structure with parent-child relationships,
as shown in Figure 1.
![HIM graph structure](/hierarchical_information_model/images/him_graph_structure.png?width=25pc)
*Figure 1. HIM graph structure

The model provides a structured solution to a scenario where an entity, e. g. a server, manages multiple domain taxonomies.

A domain can in the HIM context be defined with the help of two dimensions:
- a 'coherent' dimension which represents information related to something that is logically coherent.
  - Examples are a car, a truck, a trailer, an airplane, etc.,
- an 'information type' dimension that is used in the description of a coherent dimension using a specific information type.
  - The currently defined information types are listed below.
    - Vehicle data
    - Data
    - Service
    - Type definitions

HIM specifies rule sets for the different information types that can then be used to define taxonomies for different coherent domains.
A HIM taxonomy must conform to the rule set of one information type.
Use case scenarios where multiple information types are needed must define these in separate taxonomies.

The creation of taxonomies for different domains is not done within the HIM project, it is expected to be done by other projects using the HIM rule sets.
One example of such a project is the [Vehicle Signal Specification](https://github.com/COVESA/vehicle_signal_specification) taxonomy.

The documentation is structured in the different rule sets shortly described below.

## HIM Vehicle Data Rule Set
[Rules](/hierarchical_information_model/vehicle_data_rule_set/) for describing data produced/consumed by vehicles that can be represented by static or dynamically changing data values.

## HIM Data Rule Set
[Rules](/hierarchical_information_model/data_rule_set/) for describing data that can be characterized as having either read-write or read-only properties.

## HIM Service Rule Set
[Rules](/hierarchical_information_model/service_rule_set/) for describing services that can be represented by procedure signatures.

## HIM Type Definition Rule Set
[Rules](/hierarchical_information_model/type_definition_rule_set/) for describing complex datatype definitions, specifically struct definitions.

## HIM Configuration Rule Set
[Rules](/hierarchical_information_model/configuration_rule_set/) for how a set of domain taxonomies is defined.

## HIM Common Rule Set
[Rules](/hierarchical_information_model/common_rule_set/) that are commonly used in the other rule sets.

## HIM profiles
The abstraction level of a model describing 'vehicledata' (or 'data') is different from the abstraction level of a model describing 'services'.
This is reflected in the different HIM rule sets.
The API exposed by a server supporting access to 'vehicledata'/'data' or 'services' will also differ due to the abstraction level differencies.
To support use case scenarios for the different abstraction levels HIM defines the following profiles:
* HIM VehicleData Profile
* HIM Data Profile
* HIM Service Profile

The profiles are described below.
HIM does not restrict use case scenarios / server implementations to support only one of the profiles,
supporting multiple profiles in parallel is a valid option.

### HIM VehicleData Profile
The HIM vehicle data profile excludes use of taxonomies that represent the Service information type.
This profile is for example compatible with the [VSS](https://github.com/COVESA/vehicle_signal_specification) taxonomy that represents vehicle data.

### HIM Data Profile
The HIM data profile excludes use of taxonomies that represent the Service information type.
The difference to the vehicle data profile is that the node types 'attribute' and 'sensor' are renamed to 'ro' read-only)
and the 'actuator' node type is renamed to 'rw' (read-write).
For domains that are not technically oriented the terminology read-only/read-write is more natural than the attribute/sensor/actuator names.
A taxonomy representing 'user data', e. g. data of the driver of a vehicle, could use the HIM Data information type.

### HIM Service Profile
The HIM Service profile excludes use of trees that contain the Vehicle data or Data information types.
This profile can for example be used in "pure" SOA architectures.

## HIM enablers
The common syntax model that is used in all profiles makes it feasible to develop a server that can expose APIs
for access to both data and services.

HIM makes it possible to abstract both 'vehicledata'/'data' and 'service' information as data in an API.
This eliminates the need to update the API when updating the respective taxonomies as the information is carried in the API payloads
and do not require any updates of the static parts of the API.

This is also the case if entire new taxonomies (domains) is introduced as the HIM model supports scenarios with multiple taxonomies.
The [HIM configuration](/hierarchical_information_model/configuration_rule_set/)
file enables these updates to be implemented without requiring server restart if that is desireable.

## Heritage
The ideas behind HIM originated in the [COVESA VSS](https://github.com/COVESA/vehicle_signal_specification) project,
when interest started to be raised for using it for of not only Data but also service data,
and for different domains than the legacy VSS passenger car domain.
This may explain why examples in this documentation are often taken from that domain.
