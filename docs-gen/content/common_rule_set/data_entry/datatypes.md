---
title: "Datatypes"
date: 2019-08-04T11:11:48+02:00
weight: 10
---

HIM supports the datatype concepts:
- primitive datatypes, 
- composite datatypes, 
- arrays.

This chapter also mentions the related concepts of [Timestamps](#timestamps), and [Data streams](#data-streams).

## Primitive datatypes

The list below shows the primitive datatypes that are supported.

Name       | Description                       | Min  | Max
:----------|:---------------------------|:-----|:---
uint8      | unsigned 8-bit integer     | 0    | 255
int8       | signed 8-bit integer       | -128 | 127
uint16     | unsigned 16-bit integer    |  0   | 65535
int16      | signed 16-bit integer      | -32768 | 32767
uint32     | unsigned 32-bit integer    | 0 | 4294967295
int32      | signed 32-bit integer      | -2147483648 | 2147483647
uint64     | unsigned 64-bit integer    | 0    | 2^64 - 1
int64      | signed 64-bit integer      | -2^63 | 2^63 - 1
boolean    | boolean value              | 0/false | 1/true
float      | floating point number      | -3.4e -38 | 3.4e 38
double     | double precision floating point number | -1.7e -300 | 1.7e 300
string     | character string           | n/a  | n/a

## Composite datatypes

HIM also supports the 'struct' composite datatype that groups a list of data variables into one object.
The data variables in the list may of be primitive or composite type.

The recommended usage of structs is to have a separate tree with the information type 'TypeDefinition' in which the struct is defined.
The node that want to declare this struct as datatype then use the path to a node of type struct (node z in the example below) in the TypeDefinition tree as reference.
```YAML
datatype: TypeDefinition.x.y.z
```
This pattern enables efficient reuse of the struct definition by multiple nodes.

HIM struct support is further described in the [Structs chapter](/hierarchical_information_model/common_rule_set/data_entry/structs).

## Arrays

The primitive and composite datatypes above all define singleton data elements, but these can all be extended to define a sequence of elements of the same datatype, an `array`.

By default the size of the array is undefined, but by use of the optional metadata `arraysize: x` the size of the array can be specified (x elements in this case).

The syntax to declare an array is to concatenate the datatype with a pair of square brackets,
as shown in the example from the VSS taxonomy by the data `Vehicle.OBD.DTCList` which contains a list of Diagnostic Trouble Code (DTC) string elements.
```YAML
DTCList:
  datatype: string[]
  type: sensor
  description: List of currently active DTCs formatted according OBD II (SAE-J2012DA_201812) standard ([P|C|B|U]XXXXX )
```

## Timestamps

Timestamps are in HIM typically represented as strings, formatted according to ISO 8601.
Timestamps shall be expressed in UTC (Coordinated Universal Time), with special UTC designator ("Z").
Time resolution SHALL at least be seconds, with subsecond resolution as an optional degree of precision when desired.
The time and date format shall be as shown below, where the sub-second data and delimiter is optional.

```
YYYY-MM-DDTHH:MM:SS.ssssssZ
```

## Data Streams

Data Entries, which describe resources offering binary streams (e.g. cameras), are not supported directly by HIM with a
dedicated data type. Instead, they are described through the meta data about the sensor itself and how to retrieve the
corresponding data stream.

A camera can be a good example of it. The Data Entry for the camera and the corresponding video stream could look like:

```YAML
Camera:
  type: branch
  description: Information about the camera and how to connect to the video stream

Camera.IsActive:
  type: actuator
  datatype: boolean
  description: If the camera is active, the client is able to retrieve the video stream

Camera.URI:
  type: sensor
  datatype: string
  description: URI for retrieving the video stream, with information on how to access the stream (e.g. protocol,  data format, encoding, etc.)

```

In this example, it shows the usage of meta data about the status of the camera. The camera can be set to active through
the same data entry (`actuator`). A dynamic data entry (`sensor`) is used for the URI of the video stream,
which is expected to provide information on how to access the stream.
