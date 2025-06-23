# HIM example files

This directory contains an example file of each of the information types that HIM currently defines:
- Resource data: HIM_Resource_example.v1.0.0.vspec
- Service data: HIM_Service_example.v1.0.0.vspec
- Type definition data: TypeDefinition_example.v1.0.0.vspec

It also contains an example file of a HIM configuration file: HIM_config_example.v1.0.0.him

The resource data example uses the vspec format that is inherited from the VSS project.
The service data and type definition data examples are expressed in YAML as it is the primary format for HIM files.
However, HIM also inherits the pattern from VSS to allow translation from the primary format into other formats such as JSON, etc.
