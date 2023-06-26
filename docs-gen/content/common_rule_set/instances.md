---
title: "Instances"
date: 2019-07-31T15:27:36+02:00
weight: 5
---

It is not uncommon that when defining HIM taxonomies there is a need to repeat branches and data entries
when multiple instances of the same object occurs. 
One mechanism to handle that is the [include directive](/hierarchical_information_model/common_rule_set/includes),
but that does not solve how to compactly describe a branch structure that uniquely enumerates multiple includes.
To avoid hard-coded repetitions of branches and data entries in the specification an instance-concept is supported.
Instances remove the need of repeating definitions, by defining at the node itself how often it occurs in
the resulting tree.
They are meant as a short-cut in the specification and can interpreted by tools that expands it to a full size tree.

![Example tree before and after tool expansion](/hierarchical_information_model/images/instances.png?width=50pc)

## Definition

### How can I create instances for my `branch`?

1. An instance can be defined in any branch node.
2. The instantiation is done for every node in the subtree of this branch node.
3. Instances are defined with the key-word `instances`, followed by its
   definition, which can be either:
   * a list of strings, where each element defines a single instance, e.g.
     `['Left','Right']` results into two instances of every following
     data entry in the path, named `Left` and `Right`
   * a string, followed by a range defined through `[n,m]`, with `n,m` as integer and `n <= m`,
     which defines the number of instances.
     `Position[1,4]` results into 4 instances of every following
     data entry in the path, named `Position1`, `Position2`, `Position3`
     and `Position4`. It is in HIM recommended to use `1` as start index for the first instance.
4. If multiple instances occur in one node or on the path to a data entry,
   the instances get combined, by the order of occurrence. Following the example above,
   four position instances will be created for each of the 'Left' and 'Right' instances,
   resulting into a total number of 8 instances of the subtree.

### How can I exclude child-nodes from instantiation?

Often it makes sense to instantiate all child-nodes of a branch.
But there are cases, when nodes are linked more to the general concept of
a subtree, but not to every the single instance of it. With examples from the VSS taxonomy,
this could be the `DoorCount`,
which would rather be `Door.Count`, `WheelDiameter`, which is rather linked
to an axle rather than the wheel itself or `Brake.FluidLevel` which is not
measured for a single break, but rather a system indication.

To exclude a child-node from the instantiation of the *direct* parent node, set the
keyword `instantiate` to `false` (`true` by default). Please check the following
example for details.

## Example

Example from the VSS taxonomy:

```YAML
# Cabin.vspec
Door:
  type: branch
  instances:
    - Row[1,4]
    - ["Left","Right"]
  description: All doors, including windows and switches
#include SingleDoor.vspec Door

Door.Count:
  datatype: uint8
  type: attribute
  default: 4
  instantiate: false
  description: Number of doors in vehicle.
```


```yml
# SingleDoor.vspec

#
# Definition of a single door
#
IsOpen:
  datatype: boolean
  type: actuator
  description: Is door open or closed
```

Results in the following dot-notated output:

```
Vehicle.Cabin.Door
Vehicle.Cabin.Door.Count
Vehicle.Cabin.Door.Row1
Vehicle.Cabin.Door.Row1.Left
Vehicle.Cabin.Door.Row1.Left.IsOpen
Vehicle.Cabin.Door.Row1.Right
Vehicle.Cabin.Door.Row1.Right.IsOpen
Vehicle.Cabin.Door.Row2
Vehicle.Cabin.Door.Row2.Left
Vehicle.Cabin.Door.Row2.Left.IsOpen
Vehicle.Cabin.Door.Row2.Right
Vehicle.Cabin.Door.Row2.Right.IsOpen
Vehicle.Cabin.Door.Row3
Vehicle.Cabin.Door.Row3.Left
Vehicle.Cabin.Door.Row3.Left.IsOpen
Vehicle.Cabin.Door.Row3.Right
Vehicle.Cabin.Door.Row3.Right.IsOpen
Vehicle.Cabin.Door.Row4
Vehicle.Cabin.Door.Row4.Left
Vehicle.Cabin.Door.Row4.Left.IsOpen
Vehicle.Cabin.Door.Row4.Right
Vehicle.Cabin.Door.Row4.Right.IsOpen
```

## Redefinition

The default instantiation provided by a standard HIM tree can be redefined  with
different instantiation information. If multiple definitions of a branch exist with different
instance definitions, then the last found definition will be used.
As an example, if only two rows of doors are needed, then a default HIM instance definition of four
can be overridden by redefining the Door branch as shown in the example below.

```YAML
#Redefinition changing number of rows from 4 to 2
#The redefinition must appear "after" the original definition
Vehicle.Cabin.Door:
  type: branch
  instances:
    - Row[1,2]
    - ["Left","Right"]
  description: All doors, including windows and switches
```

## Recommendations

The HIM design shall support configurability of multiple variants of a "standard" taxonomy.
This means that the default instantiation of a HIM taxonomy may not fit all needs of it.
The following is an example from the VSS taxonomy where parts of the windshield signals defined in `Body.vspec` are shown below.
VSS offers the possibility to control windshield heating separately for front and rear windshield,
and VSS also gives the possibility to report washer fluid level separately for each windshield.
This fits very well for a vehicle that has separate washer fluid containers for front and rear windshield
and that offers heating for both windshields. But that is not the case for all vehicles,
it is not even certain that all vehicles have two windshields. This sections gives recommendations on how
to use VSS for a vehicle if the VSS specification does not offer an exact match of the capabilities of the vehicle.

```YAML
Windshield:
  type: branch
  instances: ["Front", "Rear"]
  description: Windshield signals

Windshield.Heating:
  type: branch
  description: Windshield heater signals

Windshield.Heating.Status:
  datatype: boolean
  type: actuator
  description: Windshield heater status. 0 - off, 1 - on

Windshield.WasherFluid:
  type: branch
  description: Windshield washer fluid signals

Windshield.WasherFluid.LevelLow:
  datatype: boolean
  type: sensor
  description: Low level indication for washer fluid. True = Level Low. False = Level OK.
```

### Recommendation: Instance Mismatch

If a vehicle does not have as many instances as specified in VSS then one
of the following methods are recommended:

- Redefine the branch. If a vehicle for example does not have a rear windshield
then append a redefinition at the end of the VSS:

```YAML
Vehicle.Body.Windshield:
  type: branch
  instances: ["Front"]
  description: Windshield signals
```

- Accept that a `branch Vehicle.Body.Windshield.Rear` will exist in the generated VSS representation,
  use mechanisms outside of VSS to ignore that branch

### Recommendation: Features shared among instances

If a feature is shared among instances, it is recommended to publish that feature for all concerned instances.

Example: In VSS washer fluid can be handled separately for front and rear windshield.
If a vehicle use a common container serving both front and rear windshield,
then it is recommended that the vehicle report information on that container in both
`Vehicle.Body.Windshield.Front.WasherFluid.LevelLow` and `Vehicle.Body.Windshield.Rear.WasherFluid.LevelLow`.

### Recommendation: Features lacking for some instances

Not all instances in a vehicle might have the same features. If e.g. the front windshield
from the example above lack a heater, then it is recommended to use mechanisms outside of VSS
to ignore `Vehicle.Body.Windshield.Front.Heating`.
