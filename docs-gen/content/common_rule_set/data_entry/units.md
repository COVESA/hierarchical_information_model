---
title: "Units"
date: 2019-08-04T12:36:12+02:00
weight: 20
---

## Introduction to Units

HIM keeps a list of the units of measurements that can be used as the value of the `unit` metadata, see the example below.
```
Vehicle.Speed:
  datatype: float
  type: sensor
  unit: km/h
  description: Vehicle speed.
```
Data that has an associated unit from the list shall declare it using the `unit` metadata. 
It is allowed for a domain to specify an additional list with domain specific units that complement the common units list.
Units shall if possible be based on [SI-units](https://www.iso.org/standard/30669.html), but exceptions like the speed example above are acceptable.
The list contains entries of a unit with a prefix, e. g. `mm` for millimeter, which are the only prefixes that are valid.
The unit must be written exactly as shown in the list.

In some cases it is natural to omit the data unit type. This concerns typically signals where data type `string` is used,
but also signals where the value just represents a number, like in the example below:
```
Vehicle.Cabin.DoorCount:
  datatype: uint8
  type: attribute
  default: 4
  description: Number of doors in vehicle.
```
HIM also supports `percent` and `ratio` as data units.

## List of supported Units

This list is mainly derived from the [International System of Units (SI)](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.330-2019.pdf).


Unit          | Domain                          | Description
:-------------|:--------------------------------|:-------------
mm            | Distance                        | Distance measured in millimeters
cm            | Distance                        | Distance measured in centimeters
m             | Distance                        | Distance measured in meters
km            | Distance                        | Distance measured in kilometers
inch          | Distance                        | Distance measured in inches
km/h          | Speed                           | Speed measured in kilometers per hours
m/s           | Speed                           | Speed measured in meters per second
m/s^2         | Acceleration                    | Acceleration measured in meters per second squared
cm/s^2        | Acceleration                    | Acceleration measured in centimeters per second squared
ml            | Volume                          | Volume measured in milliliters
l             | Volume                          | Volume measured in liters
cm^3          | Volume                          | Volume measured in cubic centimeters
celsius       | Temperature                     | Temperature measured in degree celsius
degrees       | Angle                           | Angle measured in degrees
degrees/s     | Angular Speed                   | Angular speed measured in degrees per second
W             | Power                           | Power measured in watts
kW            | Power                           | Power measured in kilowatts
PS            | Power                           | Power measured in horsepower
kWh           | Energy Consumption              | Energy consumption measured in kilowatt hours
g             | Weight                          | Mass measured in grams
kg            | Weight                          | Mass measured in kilograms
lbs           | Weight                          | Mass measured in pounds
V             | Electric Potential              | Electric potential measured in volts
A             | Electric Current                | Electric current measured in amperes
Ah            | Electric Charge                 | Electric charge measured in ampere hours
ms            | Time                            | Time measured in milliseconds
s             | Time                            | Time measured in seconds
min           | Time                            | Time measured in minutes
h             | Time                            | Time measured in hours
day           | Time                            | Time measured in days
weeks         | Time                            | Time measured in weeks
months        | Time                            | Time measured in months
years         | Time                            | Time measured in years
UNIX Timestamp| Time                            | Unix time is a system for describing a point in time. It is the number of seconds that have elapsed since the Unix epoch, excluding leap seconds.
mbar          | Pressure                        | Pressure measured in millibars
Pa            | Pressure                        | Pressure measured in pascal
kPa           | Pressure                        | Pressure measured in kilopascal
stars         | Rating                          | Rating measured in stars
g/s           | Mass per time                   | Mass per time measured in grams per second
g/km          | Mass per distance               | Mass per distance measured in grams per kilometers
kWh/100km     | Energy Consumption per distance | Energy consumption per distance measured in kilowatt hours per 100 kilometers
ml/100km      | Volume per distance             | Volume per distance measured in milliliters per 100 kilometers
l/100km       | Volume per distance             | Volume per distance measured in liters per 100 kilometers
l/h           | Flow                            | Flow measured in liters per hour
mpg           | Distance per Volume             | Distance per volume measured in miles per gallon
N             | Force                           | Force measured in newton
Nm            | Torque                          | Torque measured in newton meters
rpm           | Rotational Speed                | Rotational speed measured in revolutions per minute
Hz            | Frequency                       | Frequency measured in hertz
ratio         | Relation                        | Relation measured as ratio
percent       | Relation                        | Relation measured in percent
... | ... | ...
