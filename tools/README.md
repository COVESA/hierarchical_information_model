# HIM tools

HIM trees are defined in the vspec format, which is the primary format of HIM.
To translate a vspec tree into other formats (YAML, JSON, binary, etc.) HIM relies on the
[VSS-tools](https://github.com/COVESA/vss-tools) exporters.

The upstream VSS-tools only supports the HIM Vehicle data profile (the historic VSS rule set).
The HIM profiles Data and Service, and the HIM Type definition rule set,
are supported by an extended version of VSS-tools, available as the tagged release `him-v1.0.0` at:

**https://github.com/UlfBj/vss-tools/releases/tag/him-v1.0.0**

The extended exporter has the same CLI as the upstream exporter, plus a top-level `--profile` option
that selects which HIM profile the input vspec file is validated against.
All exporters (yaml, json, binary, etc.) respect the option.

| `--profile` value | HIM profile | Node types |
|-------------------|-------------|------------|
| `vehicle-data` (default) | Vehicle data | `branch`, `sensor`, `actuator`, `attribute`, `struct`, `property` |
| `data` | Data | `branch`, `ro`, `rw`, `struct`, `property` |
| `service` | Service | `branch`, `procedure`, `iostruct`, `symlink`, `attribute`, `struct`, `property` |

The node types `branch`, `struct`, `property`, and `typedef` are accepted by all profiles,
so a Type definition tree can be exported regardless of which profile is selected.

The older, Go-based `himTools` that previously was located in this folder has been removed as it is superseded by the above.

## Installation

Python 3.10 or later is required. Install the extended VSS-tools in a Python virtual environment.
It must not be installed together with the standard `vss-tools` package as both provide the `vspec` command.

```
$ python -m venv .venv
$ source .venv/bin/activate                  # Windows: .venv\Scripts\activate
(.venv)$ pip uninstall vss-tools -y          # only needed if the standard vss-tools was installed in this venv
(.venv)$ pip install git+https://github.com/UlfBj/vss-tools.git@him-v1.0.0
(.venv)$ vspec --help                        # verify that the "--profile" option is listed
```

When done, leave the virtual environment by `deactivate`. Next time it is sufficient to activate it again.

## Exporting the HIM example trees

The [examples](../examples) folder contains one vspec file for each HIM information type.
The commands below are issued from the root of this repository, with the virtual environment activated.

The exporter validates the `unit` values against a units file and a quantities file.
By default it looks for `units.yaml` and `quantities.yaml` in the same folder as the vspec file.
The examples folder contains these files, covering the units used by the examples.
A tree in another folder that uses the `unit` property needs its own files, e.g. the ones of the
[VSS catalog](https://github.com/COVESA/vehicle_signal_specification/tree/master/spec),
which can be pointed to by the `-u` and `-q` options.

### Vehicle data profile

```
(.venv)$ vspec export yaml -s examples/HIM_VehicleData_example.vspec -o HIM_VehicleData_example.yaml
```

### Data profile

```
(.venv)$ vspec --profile data export yaml -s examples/HIM_Data_example.vspec -o HIM_Data_example.yaml
```

### Service profile

The service tree refers to the types defined in the type definition tree, which therefore must be provided by the `-t` option.
The exported type definition tree is written to the file given by `--types-output`.

```
(.venv)$ vspec --profile service export yaml -s examples/HIM_Service_example.vspec -t examples/TypeDefinition_example.vspec -o HIM_Service_example.yaml --types-output TypeDefinition_example.yaml
```

The instances of the Seating service group (`Row[1,2]` x `["DriverSide","PassengerSide"]`) are expanded in the exported tree.

### Type definition tree only

The type definition tree can be exported by itself with any profile selected.

```
(.venv)$ vspec export yaml -s examples/TypeDefinition_example.vspec -o TypeDefinition_example.yaml
```

### Other export formats

Replace `yaml` with another exporter, e.g. `json` or `binary`, and change the output file name accordingly.
Run `vspec export --help` to list the available exporters and `vspec export <exporter> --help` for the options of an exporter.

```
(.venv)$ vspec --profile service export binary -s examples/HIM_Service_example.vspec -t examples/TypeDefinition_example.vspec -o HIM_Service_example.binary --types-output TypeDefinition_example.binary
```


## Using an exported service tree in a VISSR server

This walkthrough puts the HIM service example tree into a [VISSR](https://github.com/COVESA/vissr) server,
discovers the service, and invokes it.
It was verified with the VISSR branch `v3.2.5`, using the HIM example files and the extended VSS-tools of this document.
Support for the HIM Service profile is at the time of writing not available on the `master` branch of VISSR,
so check the README of VISSR for the branch that currently has it.

1. Export the service tree and its type definition tree to the binary format, see the previous section.
2. Clone VISSR and switch to a branch that supports the Service profile:
   ```
   $ git clone --branch v3.2.5 https://github.com/COVESA/vissr.git
   ```
3. Copy the two binary files to the `server/vissv2server/forest` directory of VISSR.
4. Replace the content of `server/vissv2server/viss.him` with the entries below, or add them to it.
   The key of the first entry, `HIM.VehicleService`, must be identical to the root node name of the service tree,
   as clients address the tree by that name. The `domain` of a service tree must end with `.Service`.
   ```
   HIM:
     type: branch
     description: Defines the set of trees that are managed as one virtual domain.

   HIM.VehicleService:
     type: direct
     domain: Vehicle.Car.Service
     version: 1.0.0
     local: forest/HIM_Service_example.binary
     description: The HIM service example tree.

   HIM.Types:
     type: direct
     domain: Vehicle.Car.TypeDefinition
     version: 0.5
     local: forest/TypeDefinition_example.binary
     description: The HIM type definition example tree.
   ```
5. Build and start the server (Go is required). The server reads `viss.him` from its working directory.
   `-s none` starts it without a state storage database, which is not needed for services.
   ```
   $ cd server/vissv2server
   $ go build
   $ ./vissv2server -s none
   ```
   The `runstack.sh` script in the VISSR root directory can be used to start the server together with a feeder,
   but it also requires the tools described in the VISSR README.
6. Connect a client to the WebSocket interface of the server, with the sub-protocol `VISS-noenc`, on port 8080.
   The HTML client `client/client-1.0/Javascript/wsclient.html` can be used for this.
   Enter the IP address of the server (`localhost` if it runs on the same machine) and click the "Server IP" button.
7. Send a service discovery request. The depth value `"0"` returns the complete subtree.
   ```
   {"action":"discover","path":"VehicleService.Seating.Row1.DriverSide.MoveSeat","depth":"0","requestId":"1"}
   ```
   The response contains the metadata of the MoveSeat procedure, i.e. its `Input` and `Output` iostructs and their properties.
8. Invoke the service.
   ```
   {"action":"invoke","path":"VehicleService.Seating.Row1.DriverSide.MoveSeat","input":{"MovementType":"longitudinal","Position":"3"},"filter":{"variant":"all"},"requestId":"2"}
   ```
   The server first responds with the service status `ONGOING` and a `serviceId`. It then sends `monitoring` messages,
   one per second in this case, with the progress in `outdata`, and ends with a message with the status `SUCCESSFUL`.

The paths of a service that manages multiple resources, like the Seating service group in the example,
contain the resource instances (`Row1.DriverSide`) as they are expanded by the exporter.
When no external service process has registered for a procedure, the VISSR server executes a built-in simulation of
the MoveSeat, ActivateMassage, and GetCapabilities services of the example tree.
A real service is implemented as an external process that registers its procedure with the server,
see the VISSR documentation for the service profile.

## The HIM configuration file

The HIM configuration file (e.g. [HIM_config_example.him](../examples/HIM_config_example.him))
is not exported or translated by the tools described here.
It is intended to be read as it is by the server that manages the forest of trees,
which uses it to find the trees and to create its internal data structures.
The node types of the configuration rule set (`direct`, `proxy`) are therefore not accepted by the exporter.

## Known limitations

The extended VSS-tools has the following limitations, i.e. the corresponding HIM rules are not verified by the tool:

* A `symlink` node is exported as it is defined; the target of the link in another tree is neither resolved nor verified.
* The tool does not verify that a `procedure` node has a `Version` child node.
* Chained typedefs, i.e. a typedef whose datatype is another typedef, are not supported.

These limitations are also noted in the [Service rule set](https://covesa.github.io/hierarchical_information_model/service_rule_set/basics/#tool-support)
and the [typedef rule set](https://covesa.github.io/hierarchical_information_model/type_definition_rule_set/typedef/).

## Further usage information

The CVED project exports its HIM Data profile tree by use of the same tool.
The usage is described in the [CVED README](https://github.com/UlfBj/cved/tree/main).
The general VSS-tools documentation, which also applies to the extended version,
is found in the [VSS-tools repository](https://github.com/COVESA/vss-tools).
