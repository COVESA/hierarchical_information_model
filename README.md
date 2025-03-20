![Status - Incubating](https://img.shields.io/static/v1?label=Status&message=Incubating&color=FEFF3A&style=for-the-badge)

# COVESA Hierarchical Information Model (HIM)

[![License](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](https://opensource.org/licenses/MPL-2.0)

Please find the official documentation generated from this repository at: [HIM Documentation](https://covesa.github.io/hierarchical_information_model/)

### Maintainers
Ulf Björkengren - Ford Motor Company



## Governance

<!-- Give a short rationale for the repository -->

This repository contains the COVESA Hierarchical Information Model (HIM) - an initiative from COVESA to create a model
that can be used to define tree structures representing information (see below) defining a domain (e.g. Truck, Trailer, or Car domains).
[VSS](https://github.com/COVESA/vehicle_signal_specification) is an example of one domain that is planned to be possible to model in HIM.

<!-- Give a short description of which group/project in COVESA that manages this repository and link to their wiki -->

This repository is managed by the [COVESA Data Expert Group](https://wiki.covesa.global/display/WIK4/Data+Expert+Group),
together with the [COVESA CVIS project](https://wiki.covesa.global/display/WIK4/CVIS+Meeting+Topics+and+Meeting+Notes).

## HIM Overview

The information of a domain that is described using HIM is represented in a graph made up of a tree structure with parent-child relationships,
as shown in Figure 1.
![HIM graph structure](docs-gen/static/images/him_graph_structure.png?raw=true)<br>
*Figure 1. HIM graph structure<br>

HIM does not have the ambition to define the content of any domain,
it stops at defining the rules for how a hierarchical representation of a domain is to be expressed.
This rule set aims at being domain agnostic and can thus be used to define the taxonomies for different domains.

Another ambition of HIM s that it shall support definition of taxonomies of different types of information,
and not only the type here called "resource data".
HIM currently supports the following "information types":
- Resource data: Information defining the data that a resource produces or consumes.
- Service data: Information defining the details of a "service" in the form of a procedure (name, input, output).
- Type definition data: Information defining complex/composite data types that are used by the other information types.

HIM also provides a rule set for the definition of a "configuration file" that can be used by a HIM enabled data server in its management of
a "forest" of multiple trees. This information can also by the server be provided to a client wanting to know what data or services the server
has to offer.

What is not an ambition of HIM is to define how the information is transported between agents,
which is left to projects defining transport protocols, interfaces, and the like.
Neither is it aiming for defining the content of any tree, which is also left to other projects.

## HIM origin

The HIM model traces its roots to the COVESA
[Vehicle Signal Specification (VSS)](https://github.com/COVESA/vehicle_signal_specification) project.

When the adoption of the VSS data model started to take off,
also requests for additional data and other functionality started to be asked for.
Instead of trying to accomodate all of these requests in the single tree that defined the VSS domain,
it was decided that an extended model was to be developed, the Hierarchical Information Model.


<!-- For status and roadmap of this repository please visit ... -->

*(There should for each COVESA repository be a page on COVESA wiki giving information on roadmaps, meetings, maintainers and so on...)*

## Pre-commit checks
The repository has [configuration file](.pre-commit-config.yaml) with pre-commits hooks.
It executes a number of checks that typically must pass for a new Pull Request to be accepted and merged.
You must manually configure pre-commit to use the provided hooks by running `pre-commit install` from the
respository top folder.

```bash
~/vehicle_signal_specification$: pip install pre-commit
~/vehicle_signal_specification$: pre-commit install
```

## Compatibility policy
`Model` refers to the HIM model as it is defined on the [HIM repository](https://github.com/COVESA/hierarchical_information_model).

Non-backwards compatible (NBC) changes to the `Model` shall be handled as follows.

* The period from that the NBC idea is introduced until it is accepted and merged shall be at least three months, including at least three WG meetings where it has been discussed.
* Any participant in these WG discussions can request a vote whether to reject the idea or not. At least three regular WG participants, or regular participants in other projects that use HIM, must vote. A majority of at least 51% is required for decision.
 
Further, NBC changes to particularly important features of the `Model` in the Axiom list below must be handled as follows.

* The period from that the NBC idea is introduced until it is accepted and merged shall be at least six months, including at least three WG meetings where it has been discussed.
* Any participant in these WG discussions can request a vote whether to reject the idea or not. At least five regular WG participants, or regular participants in other projects that use HIM, must vote. A majority of at least 80% is required for decision.
 
Axioms:

* The tree structure, where the tree root node name is used for identifying one tree in a forest, is a fundamental part of the `Model`.
* The metadata defining a tree node shall not include references to other nodes, whether it is in the same tree or in another tree.
* Low complexity and ease of understanding have priority in the evolution of the `Model`.
* As an example of keeping complexity low, the `Model` shall not include metadata related to the transport of data or the capture of data.
Adding metadata that is not part of the `Model` may be done in an implementation of the `Model`.
 
The merging of an NBC change must lead to a new major release. The release notes must clearly state the NBC change,
and if it involves an Axiom then also the rationale behind it.

If a vote is requested, then all other projects known to use the `Model` must be notified in due time about the coming vote.
