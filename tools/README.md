# HIM tools overview

The HIM tools is a supplement tool besides the VSS-tools that can be used to create exporter formats of vspec-defined service trees.

Currently it supports the exporter formats:
* YAML
* JSON
* binary

All the key names of properties used to define a data tree can be used, plus te following:
* procedure  // the name of the procedure that is called to execute the service
* iostruct   // a struct that is dedicated to define the procedure input and output parameters.

The HIM tools does not support the "instances" feature.

## HIM tools generated service tree in a VISSR context
HIM tools can be used to generate a binary formatted tree that can to a limited extent be used by the VISSR server.
The limitation depends on that the current version of VISSR only supports the HIM data profile, not he HIM service profile.
It is herefore not possible to invoke a service on VISSR, i. e. make a call to a procedure node together with its input parameters.
What can be done however, is to issue a "service discovery" request to the server, on which it will return the metadata of the addressed subtree.

To set up a demo to test this the following needs to be done.

1. Clone the HIM repo.
2. Build the HIM tools. Issue "$ go build -o himTools" in the tools directory
3. Create a binary formatted service tree.
This can be done by using the example vspec service tree file VehicleServices/VehicleServiceSpecification.vspec as input to the HIM tools.
This can be done by issuing "$ ./himTools -e binary," resulting in the binary file himExporter.binary.
4. Clone the VISSR repo.
5. Copy the file himExporter.binary to the vissr/server/vissv2server/forest directory.
6. Add the following lines to the vissr/server/vissv2serverviss.him file
HIM.VehicleService:
  type: direct
  domain: Vehicle.Car.Service
  version: 0.0.1
  local: forest/himExport.binary
#  public: https://himrepo.oem.com?instance=Vehicle.Car.ResourceData.X.Y.Z
  description: A HIM service tree.
7. Start the VISSR server plus a feeder. This can be done by issuing "$ ./runstack startme" in the vissr directory.
8. Start a VISR client. This can e. g. be done by clicking on the HTML-based client vissr/client/client-1.0/Javascript/wsclient.html that will start up in a browser tab.
9. Find the IP address of the machine, or use "localhost" if the server runs on the same machine. Enter it in the box and click the "Server IP" button. "Connected" shall then be shown in the tab below.
10. Enter the service discovery request {"action":"get","path":"VehicleService","filter":{"variant":"metadata","parameter":"0"},"requestId":"245"} in the other box and click the Send button.
11. The server response is then shown in the tab below. This contains a JSON formatted representation of the metadata in the service tree.

The above sequence should also be performed for the VehicleServices/TypeDefinition/Type.vspec file to create a binary Type tree
but as the VISSR server currently does not utilize the Type tree for struct verification it may be skipped at this point.

If the COVESA VISS project decides to add support for the HIM service profile in a coming version, then it is likely that VISSR will support that,
so that clients also can invoke execution of services on VISSR.
