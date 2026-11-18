# HIM example files

This directory contains an example file of each of the information types that HIM currently defines:
- Vehicle data: HIM_VehicleData_example.vspec
- Data: HIM_Data_example.vspec
- Service: HIM_Service_example.vspec
- Type definition data: TypeDefinition_example.vspec

It also contains an example file of a HIM configuration file: HIM_config_example.him

The information type example files shall use the vspec format as the primary format.
However, HIM inherits the pattern from VSS to allow translation from the primary format into other formats such as JSON, etc.

The files units.yaml and quantities.yaml define the units that are used by the example files (a subset of the VSS catalog).
They are found by the exporter when it is run on the example files, see the [HIM tools](../tools/README.md) description.
The example files shall use ASCII characters only, as some platforms do not read files as UTF-8 by default.
