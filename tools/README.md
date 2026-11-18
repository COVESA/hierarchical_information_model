# HIM tools

HIM trees are defined in the vspec format, which is the primary format of HIM.
To translate a vspec tree into other formats (YAML, JSON, binary, etc.) HIM relies on the
[VSS-tools](https://github.com/COVESA/vss-tools) exporters.

The upstream VSS-tools only supports the HIM Vehicle data profile (the historic VSS rule set).
The HIM profiles Data and Service, and the HIM Type definition rule set,
are supported by an extended version of VSS-tools found on the `him-support` branch of this fork:

**https://github.com/UlfBj/vss-tools/tree/him-support**

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
(.venv)$ pip install git+https://github.com/ulfbj/vss-tools.git@him-support
(.venv)$ vspec --help                        # verify that the "--profile" option is listed
```

When done, leave the virtual environment by `deactivate`. Next time it is sufficient to activate it again.

## Exporting the HIM example trees

The [examples](../examples) folder contains one vspec file for each HIM information type.
The commands below are issued from the root of this repository, with the virtual environment activated.

### Units files

The exporter validates the `unit` values against a units file and a quantities file.
By default it looks for `units.yaml` and `quantities.yaml` in the same folder as the vspec file,
and the examples folder does not contain any.
The Data example uses the units `cm` and `kg`, and the Type definition example uses `percent`,
which are all defined by the VSS catalog.
Download the files from the VSS repository, and point to them by the `-u` and `-q` options:

```
$ curl -O https://raw.githubusercontent.com/COVESA/vehicle_signal_specification/master/spec/units.yaml
$ curl -O https://raw.githubusercontent.com/COVESA/vehicle_signal_specification/master/spec/quantities.yaml
```

Without these files the export fails with the error `'cm' is not a valid unit`.
Trees that do not use the `unit` property do not need them.

### Vehicle data profile

```
(.venv)$ vspec export yaml -s examples/HIM_VehicleData_example.vspec -u units.yaml -q quantities.yaml -o HIM_VehicleData_example.yaml
```

### Data profile

```
(.venv)$ vspec --profile data export yaml -s examples/HIM_Data_example.vspec -u units.yaml -q quantities.yaml -o HIM_Data_example.yaml
```

### Service profile

The service tree refers to the types defined in the type definition tree, which therefore must be provided by the `-t` option.
The exported type definition tree is written to the file given by `--types-output`.

```
(.venv)$ vspec --profile service export yaml -s examples/HIM_Service_example.vspec -t examples/TypeDefinition_example.vspec -u units.yaml -q quantities.yaml -o HIM_Service_example.yaml --types-output TypeDefinition_example.yaml
```

The instances of the Seating service group (`Row[1,2]` x `["DriverSide","PassengerSide"]`) are expanded in the exported tree.

### Type definition tree only

The type definition tree can be exported by itself with any profile selected.

```
(.venv)$ vspec export yaml -s examples/TypeDefinition_example.vspec -u units.yaml -q quantities.yaml -o TypeDefinition_example.yaml
```

### Other export formats

Replace `yaml` with another exporter, e.g. `json` or `binary`, and change the output file name accordingly.
Run `vspec export --help` to list the available exporters and `vspec export <exporter> --help` for the options of an exporter.

```
(.venv)$ vspec --profile service export binary -s examples/HIM_Service_example.vspec -t examples/TypeDefinition_example.vspec -u units.yaml -q quantities.yaml -o HIM_Service_example.binary
```

The binary format is the one used by [VISSR](https://github.com/COVESA/vissr) for its trees.
Parsers for it, in C and Go, that support all HIM node types are available in the `binary` folder of the extended VSS-tools.

### Known limitations

* The HIM configuration file (e.g. [HIM_config_example.him](../examples/HIM_config_example.him)) cannot be exported,
  as the node types of the configuration rule set (`direct`, `proxy`) are not supported by the extended VSS-tools.
* A `symlink` node is exported as it is defined; the target of the link in another tree is neither resolved nor verified.
* The tool does not verify that a `procedure` node has a `Version` child node.
* Chained typedefs, i.e. a typedef whose datatype is another typedef, are not supported.
* On Windows, set the environment variable `PYTHONUTF8=1` if the export fails with a `UnicodeDecodeError`.
  The example files contain non-ASCII characters (typographic quotes), and Python on Windows does not read files as UTF-8 by default.

## Further usage information

The CVED project exports its HIM Data profile tree by use of the same tool.
The usage is described in the [CVED README](https://github.com/UlfBj/cved/tree/main).
The general VSS-tools documentation, which also applies to the extended version,
is found in the [VSS-tools repository](https://github.com/COVESA/vss-tools).
