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
    - Resource data
    - Service data
    - Type definitions

HIM specifies rule sets for the different information types that can then be used to define taxonomies for different coherent domains.

The creation of taxonomies for different domains is not done within the HIM project, it is expected to be done by other projects using the HIM rule sets.
One example of such a project is the [Vehicle Signal Specification](https://github.com/COVESA/vehicle_signal_specification) project.

The documentation is structured in the different rule sets shortly described below.

## HIM Resource Data Rule Set
[Rules](/hierarchical_information_model/resource_data_rule_set/) for describing resources that can be represented by static or dynamically changing data values.

## HIM Service Data Rule Set
[Rules](/hierarchical_information_model/service_data_rule_set/) for describing services that can be represented by procedure signatures.

## HIM Type Definition Rule Set
[Rules](/hierarchical_information_model/type_definition_rule_set/) for describing complex datatype definitions, specifically struct definitions.

## HIM Configuration Rule Set
[Rules](/hierarchical_information_model/configuration_rule_set/) for how a set of domain taxonomies is defined. 

## HIM Common Rule Set
[Rules](/hierarchical_information_model/common_rule_set/) that are commonly used in the other rule sets.

## Heritage
The ideas behind HIM originated in the [COVESA VSS](https://github.com/COVESA/vehicle_signal_specification) project,
when interest started to be raised for using it for of not only resource data but also service data,
and for different domains than the legacy VSS passenger car domain.
This may explain why examples in this documentation are often taken from that domain.
