---
title: "Tree Identification"
weight: 1
---
## Tree Identification

As defined in [Configuration Overview](../../configuration_rule_set/basics/#configuration-overview) a HIM tree is identified by the ‘domain’ and ‘version’ properties. The tree that is identified by this shall contain the same information. This can be achieved in two ways.

* For legacy tree compatibility reasons the identity information may be present in a legacy format. It must be possible from the legacy identity information to derive both a unique domain specifier and version data.
  * An example is a [VSS]( https://covesa.github.io/vehicle_signal_specification/) tree that contains a branch named VersionVSS, having children containing Major, Minor, Patch (and Label) data.
  The node name VersionVSS, instead of Version, is used to derive that this is a tree belonging to the 'VSS domain'.
* In all other cases the identity information shall be represented by the following structure:
  * A branch with the name Identifier shall exist directly under the root node. This branch shall contain:
   1. A node with the name Domain with the node type set to ‘attribute’, or ‘ro’ for the Data profile, and a ‘default’ property set to the domain value in the HIM configuration file.
   2. A branch named Version that shall have the three children nodes Major, Minor, Patch, each having a 'default' property set to the respective version number in the HIM configuration file.

The HIM configuration file shall not for itself contain any tree identification data.
