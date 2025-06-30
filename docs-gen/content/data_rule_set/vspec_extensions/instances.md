---
title: "Instances"
weight: 5
---

The data model resembles primarily the physical structure of the vehicle, so
quite often there is a need to repeat branches and data entries
(e.g. doors, axles, etc). To avoid hard-coded repetitions of
branches and data entries in the specification an instance-concept is supported.
Instances remove the need of repeating definitions, by defining at the node itself how often it occurs in
the resulting tree. They are meant as a short-cut in the specification and
interpreted by the tools.

![Example tree](/hierarchical_information_model/images/instances.png?width=60pc)



## Definition

### How can I create instances for my `branch`?

1. An instance can be defined in any branch.
2. The instantiation is done for every node in the following path.
3. Instances are defined with the key-word `instances`, followed by its
   definition, which can be either:
   * a list of strings, where each element defines a single instance, e.g.
     `['Left','Right']` results into two instances of every following
     data entry in the path, named `Left` and `Right`
   * a string, followed by a range defined through `[n,m]`, with `n,m` as integer and `n <= m`,
     which defines the number of instances.
     `Position[1,4]` results into 4 instances of every following
     data entry in the path, named `Position1`, `Position2`, `Position3`
     and `Position4`. It is recommended to use `1` as start index for the first row/axle/position/...
4. If multiple instances occur in one node or on the path to a data entry,
   the instances get combined, by the order of occurrence. Following the example above,
   four position instances will be created for each of the 'Left' and 'Right' instances,
   resulting into a total number of 8 instances.

### How can I exclude child-nodes from instantiation?

Often it makes sense to instantiate all child-nodes of a branch.
But there are cases, when nodes are linked more the general concept of
a branch, but not to the single instance. This could be the `DoorCount`,
which would rather be `Door.Count`, `WheelDiameter`, which is rather linked
to an axle rather than the wheel itself or `Brake.FluidLevel` which is not
measured for a single break, but rather a system indication.

To exclude a child-node from the instantiation of the *direct* parent node, set the
keyword `instantiate` to `false` (`true` by default). Please check the following
example for details.

## Example

The example from above in the specification:

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

It is possible to override the default instantiation provided by redefining the branch with
different instantiation information. If multiple definitions of a branch exist with different
instance definitions, then the last found definition will be used.
As an example, if only two rows of doors are needed, then the default instance definition
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

The instantiation feature is designed to be able to be configured for a wide range of vehicles.
This means that the default instantiation represented in the standard may not fit every vehicle.
An example can be seen in the windshield signals defined in `Body.vspec`, parts of them are shown below.
It offers the possibility to control windshield heating separately for front and rear windshield,
and it also gives the possibility to report washer fluid level separately for each windshield.
This fits very well for a vehicle that has separate washer fluid containers for front and rear windshield
and that offers heating for both windshields. But that is not the case for all vehicles,
it is not even certain that all vehicles have two windshields. This sections gives recommendations on how
to use it for a vehicle if the default specification does not offer an exact match of the capabilities of the vehicle.

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

If a vehicle does not have as many instances as specified by default then one
of the following methods are recommended:

- Redefine the branch. If a vehicle for example does not have a rear windshield
then append a redefinition at the end of the vspec file:

```YAML
Vehicle.Body.Windshield:
  type: branch
  instances: ["Front"]
  description: Windshield signals
```

- Accept that a `branch Vehicle.Body.Windshield.Rear` will exist in the generated tree representation,
  use other mechanisms to ignore that branch

### Recommendation: Features shared among instances

If a feature is shared among instances, it is recommended to publish that feature for all concerned instances.

Example: A washer fluid can be handled separately for front and rear windshield.
If a vehicle use a common container serving both front and rear windshield,
then it is recommended that the vehicle report information on that container in both
`Vehicle.Body.Windshield.Front.WasherFluid.LevelLow` and `Vehicle.Body.Windshield.Rear.WasherFluid.LevelLow`.

### Recommendation: Features lacking for some instances

Not all instances in a vehicle might have the same features. If e.g. the front windshield
from the example above lack a heater, then it is recommended to use other mechanisms
to ignore `Vehicle.Body.Windshield.Front.Heating`.
