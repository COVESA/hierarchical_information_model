---
title: "Value Restrictions"
date: 2019-08-04T12:37:12+02:00
weight: 50
---

HIM supports the following metadata for restricting the valid values of data.
- min
- max
- allowed

## Min
The minimum value, within the interval of the given datatype, that the data entry can be assigned.
If omitted, the minimum value will be the “Min” value for the given datatype.
The `min` restrictions can only be used data having a "number"datatype (intx, uintx, float, double).
Must not be specified if `allowed` is defined in the same node.

## Max
The maximum value, within the interval of the given datatype, that the data entry can be assigned.
If omitted, the maximum value will be the “Max” value for the given datatype.
The `max` restrictions can only be used data having a "number"datatype (intx, uintx, float, double).
Must not be specified if `allowed` is defined in the same node.

## Min and Max Example
```YAML
Powertrain.TractionBattery.StateOfCharge.Current:
  type: sensor
  unit: percent
  datatype: float
  min: 0
  max: 100.0
  description: Physical state of charge of the high voltage battery, relative to net capacity.
```

## Enumerated values
Definition of enumerated set of values can be done using one of two different syntax models.
The two models are referred to by the key values `enum` or `allowed`.
If one is used in a node of a tree the other must not also be used in the same tree.

### Enum
The `enum` restriction defines a array of accepted data values where each value is associated with a symbolic value in a key-value declaration as shown below.
```YAML
Vehicle.Exterior.RoadSurfaceCondition:
  datatype: uint8
  type: sensor
  enum:
    UNKNOWN: 0
    DRY: 1
    WET: 2
    SNOW: 3
    ICE: 4
    SLUSH: 5
    WET_ICE: 6
    LOOSE_GRAVEL: 7
  description: Assessed condition of the road surface beneath or ahead of the vehicle.
```
The datatype must be of number type, i. e. one of [uint8, uint16, uint32, uint64, int8, int16, int32, int64] or arrays of these.

### Allowed
The `allowed` restriction defines a array of accepted data values, defined as a comma separated list of values confined within square brackets, see example below.
It is expected, that any value not mentioned in the array is considered an error and the implementation of the specification shall react accordingly.
The datatype of the array elements shall be compatible with the `datatype` defined for the data entry itself.
```YAML
SteeringWheel.Position:
  datatype: string
  type: attribute
  default: 'FRONT_LEFT'
  allowed: ['FRONT_LEFT', 'FRONT_RIGHT']
  description: Position of the steering wheel on the left or right side of the vehicle.

```
If `allowed` is set, `min` or `max` cannot be defined.

`allowed` is valid for all datatypes.

### Recommendation on String values

For string values used for `allowed` statements it is recommended to start with `A-Z` and then use only `A-Z`, `0-9` and underscore (`_`).
It is recommended to use single quotes (`'`) as delimiter before and after the string value.
It is not recommended to specify a dedicated value corresponding to "unknown" or "undefined" as data values in general are not expected to have unique values for this.

### Allowed values for array types

The `allowed` keyword can also be used for array datatypes. In that case, `allowed` specifies the only valid values for all the array elements, see example below.
```YAML
DogBreeds:
  datatype: string[]
  type: attribute
  allowed: ['AKITA', 'BOXER', 'DACHSHUND', 'PAPILLON', 'PUG', 'VIZSLA']
  description: Brief list of dog breeds.
```
Examples of valid arrays:
```
  [] # Empty array
  ['BOXER']
  ['PAPILLON', 'VIZSLA', 'BOXER', 'AKITA', 'DACHSHUND']
  ['PUG', 'PUG'] # duplication is allowed
```
Example of an invalid array:
```
  ['PAPILLON', 'VIZSLA', 'LOBSTER'] # LOBSTER is not in the allowed value list
```

### Allowed for struct types

Please see [struct]({{< ref "structs#allowed-values" >}} ) documentation.
