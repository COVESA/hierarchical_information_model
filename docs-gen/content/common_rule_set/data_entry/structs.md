---
title: "Structs Explained"
date: 2019-08-04T11:11:48+02:00
weight: 15
---

As mentioned in the chapter [Composite datatypes](/hierarchical_information_model/common_rule_set/data_entry/datatypes#composite-datatypes)
a struct is defined using the node types 'struct' and 'property' in trees having the information type TypeDefinition, and it is referred to by using its path.

The struct node type can also be used in trees having the information type Service, which is described in [HIM Services Rule Set](/hierarchical_information_model/services_rule_set/).

A struct definition in a Typedefinition tree can contain a member that has a different struct as its datatype.
This other struct must then be defined elsewhere in the Typedefinition tree and referred to by its path.

## Intended usage

The struct support in HIM is introduced to facilitate logical binding/grouping of data that originates from the same source.
It is intended to be used only when it is important that the data is read or written in an atomic operation.
It is not intended to be used to specify how data shall be packaged and serialized when transported.

The order of elements in a struct is from a HIM perspective considered as arbitrary.
The HIM project will for this reason not publish guidelines on how to order items in the struct to minimize size,
and no concept for introducing padding will exist.

Structs shall be used in a HIM based catalog only when considered to give a significant advantage compared to using only primitive types.

## Structs vs. Aggregate

HIM supports the keyword `aggregate` that can be used on [branches](/hierarchical_information_model/common_rule_set/basics#branch-node-type)
to indicate that the branch preferably shall be read and written in atomic operations.

There have been criticism that `aggregate` changes the semantic meaning of branches and signals, i.e. that a data variable is no longer handed as an independent object.
The exact meaning of `aggregate` is furthermore not well defined.
Shall for example a write request (or update of sensor values) be rejected by an implementation
if not all signals in the branch are updated in the same operation?
Semantic interpretation is also ambiguous if the branch contains a mix of e.g. data objects with read-only and read-write permissions.
Using structs as datatype is better aligned with the view that HIM data objects are independent objects regardless of whether they are of primitive or composite datatypes,
and the semantic ambiguities related to `aggregate` are not present for structs.

Aggregate could however be useful as information on deployment level.
It gives the possibility to indicate that in this particular deployment the signals in the branch shall be treated as an aggregate.
Exact meaning of the `aggregate` keyword is then deployment specific.
With this view, aggregate shall never be used in a standard HIM catalog, but can be added in a subsequent process for deployment-specific purposes.

## General Idea and Basic Semantics

A signal of struct type shall be defined in the same way as other VSS signals,
the only difference would be that instead of using a primitive type there shall be a reference to a struct datatype.
This means that structs can be used for all types of VSS signals (i.e. sensor, attribute and actuator).
If a signal of struct type is sent or received, VSS expects all included items to have valid values, i.e. all items are mandatory.
For example, if a struct contains the items A, B and C - then it is expected that the sent signal contains value for all items.
If some items are considered optional then the value range of the items must be adapted to include values indicating "not available" or "undefined",
or additional items needs to be added to indicate which items that have valid values.

HIM makes no assumption on how structs are transferred or stored by implementations.
It is however expected that they are read and written by atomic operations.
This means that the data storage shall be "locked" while the items of the struct are read, preventing changes to happen while reading/writing the items.

Structs shall be defined in a separate tree. This means that signal definitions and types cannot exist in the same files.
IS THE ABOVE CORRECT? STRUCTS CANNOT ALTERNATVELY BE DEFINED AS A CHILD SUBTREE? 

## Naming Restrictions

The HIM syntax shall not enforce any restrictions on naming for the datatype definition tree.
It may even use the same branch structure as the data tree.
This means that it theoretically at the same time could exist both a data object `A.B.C` and a struct definition `A.B.C`.
This is not a problem as it always from context is clear whether a name refers to a data object or a datatype definition.

## Simple Definition and Usage

This could be a hypothetical content of a HIM datatype definition file

```
Type:
  type: branch
  description: Root node for the datatype definition tree

Type.DeliveryInfo:
  type: struct
  description: A struct type containing info for each delivery

Type.DeliveryInfo.Address:
  datatype: string
  type: property
  description: Destination address

Type.DeliveryInfo.Receiver:
  datatype: string
  type: property
  description: Name of receiver
```

This struct definition could then be referenced from a HIM data tree.
```
Delivery:
  datatype: Type.DeliveryInfo
  type: sensor
```
The datatype definition file may contain sub-branches and `#include`-statements just like regular HIM files.

```
Type:
  type: branch
  description: Root node for the datatype definition tree

Type.Powertrain:
  type: branch
  description: Powertrain types.
#include Powertrain/Powertrain.him Types.Powertrain

```

## Name resolution

Two ways of referring to a type are considered correct:

In Datatype Definition Tree:
* Reference by absolute path
* Reference by (leaf) name to a struct definition within the same branch

In Signal Tree:
* Reference by absolute path

Relative paths (e.g. `../Powertrain.SomeStruct`) are not allowed.
Structs in parent branches will not be visible, in those cases absolute path needs to be used instead.

*The reference by leaf name is applicable only for structs referencing other structs!*

## Array Support

It is allowed to use a struct datatype in an array.
```
DeliveryList:
  datatype: Types.DeliveryInfo[]
  type: sensor
  description: List of deliveries
```

By default the array has an arbitrary number of element and may be empty.
If a fixed size array is wanted the keyword `arraysize` can be used to specify size:

```
DeliveryList:
  datatype: Types.DeliveryInfo[]
  arraysize: 5
  type: sensor
  description: List of deliveries
```

## Structure in Structure

It is allowed to refer to a structure type from within a structure as shown in the example below.

```
Type:
  type: branch
  description: Root node for the datatype definition tree

Type.OpenHours:
  type: struct
  description: A struct type containing information on open hours

Type.OpenHours.Open:
  datatype: uint8
  type: property
  max: 24
  description: Time the address opens

Type.OpenHours.Close:
  datatype: uint8
  type: property
  max: 24
  description: Time the address close

Type.DeliveryInfo:
  type: struct
  description: A struct type containing info for each delivery

Type.DeliveryInfo.Address:
  datatype: string
  type: property
  description: Destination address

Type.DeliveryInfo.Receiver:
  datatype: string
  type: property
  description: Name of receiver

Type.DeliveryInfo.Open:
  datatype: Type.OpenHours
  type: property
  description: When is receiver available

```

## Order of declaration/definition

The order of declaration/definition shall not matter.
As signals and types are defined in different trees this is a topic only for struct definitions referring to other struct definitions.
A hypothetical example is shown below. An item in the struct `DeliveryInfo` can refer to the struct `OpenHours` even if that struct
is defined further down in the same file.
If using `-vt < file>` multiple times all files except the first will be treated similar to overlays.
This means that is allowed to define `A.B.C` in multiple files, but then subsequent (re-)definitions will overwrite
what has been defined previously.

```
DeliveryInfo:
  type: struct
  description: A struct type containing info for each delivery

...

DeliveryInfo.Open:
  datatype: OpenHours
  type: property
  description: When is receiver available

OpenHours:
  type: struct
  description: A struct type containing information on open hours
...
```


## Inline Struct

The node type 'struct' must only be used in trees of the information types 'Typedefinition' and 'Service'.
Trees of any other information type that contains nodes having a datatype that is a struct
must refer to the struct definition in a 'Typedefinition' tree by use of the path to that struct.

## Default Values

Members of a struct may have a default value defined.

## Allowed Values

A member of a struct may use the metadata `allowed` to restrict its allowed values (if `allowed` is supported for the used datatype).

