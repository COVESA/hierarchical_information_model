/**
* (C) 2025 Ford Motor Company
*
* All files and artifacts in the repository at https://github.com/covesa/hierarchical_information_model
* are licensed under the provisions of the license provided by the LICENSE file in this repository.
*
**/

package main

import (
	"fmt"
	"os"
//	"os/exec"
	"strings"
//	"strconv"
	"encoding/json"
	"bufio"
//	"io/fs"
	"path/filepath"
	"github.com/akamensky/argparse"
	"gopkg.in/yaml.v3"
)

type PropertyData struct {
	Name string
	NodeType string
	Datatype string
	Allowed []string
	Min string
	Max string
	Unit string
	Default string
	Description string
	Uuid string
	Validate uint8
	Children uint8
}

type VspecContext struct {
	Depth int
	BasePath string
	Path string
	FName string
	Exporter func(VspecContext, PropertyData)
	ExporterFp *os.File
	TmpStorage *[]PropertyData
	FileWrite bool
}


func getNodeName(line string, nodeName *string) {
	if len(line) > 0 && line[0] != ' ' && line[0] != '#' && line[len(line)-1] == ':' {
		*nodeName = line[:len(line)-1]
	}
}


func clearPropertyNode(nextNodeName string) PropertyData {
	var propertyNode PropertyData
	propertyNode.Name = nextNodeName
	propertyNode.NodeType = ""
	propertyNode.Datatype = ""
	propertyNode.Allowed = nil
	propertyNode.Min = ""
	propertyNode.Max = ""
	propertyNode.Unit = ""
	propertyNode.Default = ""
	propertyNode.Description = ""
	return propertyNode
}

func getNode(scanner *bufio.Scanner, nextNodeName string) (string, PropertyData, bool) {
	var line string
	continueScan := true
	thisNode := clearPropertyNode(nextNodeName)
	nextLine := ""
	nodeComplete := false
	for continueScan && !nodeComplete {
		if len(nextLine) == 0 {
			continueScan = scanner.Scan()
			line = scanner.Text()
		} else {
			line = nextLine
			nextLine = ""
		}
		key, value := analyzeLine(line)
		switch key {
			case "name":
				if len(thisNode.Name) == 0 {
					thisNode.Name = value
				} else{
					nextLine = line
					nodeComplete = true
				}
			case "type":
				thisNode.NodeType = value
			case "datatype":
				thisNode.Datatype = value
			case "allowed":
				thisNode.Allowed, nextLine, continueScan = getAllowedValues(scanner)
			case "min":
				thisNode.Min = value
			case "max":
				thisNode.Max = value
			case "unit":
				thisNode.Unit = value
			case "default":
				thisNode.Default = value
			case "description":
				thisNode.Description = value
			case "skipline":
				nextLine = line
				nodeComplete = true
		}
	}
	return nextLine, thisNode, continueScan
}

func analyzeLine(line string) (string, string) {
	if len(line) > 0 && line[len(line)-1] == ':' && line[0] != ' ' {
		return "name", line[:len(line)-1]
	}
	if strings.Contains(line, "datatype:") {
		return "datatype", extractValue(line)
	}
	if strings.Contains(line, "type:") {
		return "type", extractValue(line)
	}
	if strings.Contains(line, "allowed:") {
		return "allowed", ""
	}
	if strings.Contains(line, "min:") {
		return "min", extractValue(line)
	}
	if strings.Contains(line, "max:") {
		return "max", extractValue(line)
	}
	if strings.Contains(line, "unit:") {
		return "unit", extractValue(line)
	}
	if strings.Contains(line, "default:") {
		return "default", extractValue(line)
	}
	if strings.Contains(line, "description:") {
		return "description", extractValue(line)
	}
	return "skipline", ""
}

func getAllowedValues(scanner *bufio.Scanner) ([]string, string, bool) {
	var line string
	continueScan := true
	var allowedValues []string
	for continueScan {
		continueScan = scanner.Scan()
		line = scanner.Text()
		if strings.Contains(line, "- ") {
			allowedValues = append(allowedValues, line)
		} else {
			return allowedValues, line, continueScan
		}
	}
	return allowedValues, "", continueScan
}

func extractValue(line string) string {
	colonIndex := strings.Index(line, ":")
	return strings.TrimSpace(line[colonIndex+1:])
}

func exportYaml(ctx VspecContext, node PropertyData) {
	indent := make([]byte, (ctx.Depth - 1)*2)
	for i := 0; i < len(indent); i++ {
		indent[i] = ' '
	}
	if len(ctx.BasePath) == 0 {
		ctx.ExporterFp.Write([]byte(string(indent) + node.Name + ":\n"))
	} else {
		ctx.ExporterFp.Write([]byte(string(indent) + ctx.BasePath + "." + node.Name + ":\n"))
	}
	ctx.ExporterFp.Write([]byte(string(indent) + `  type: ` + node.NodeType + "\n"))
	if len(node.Datatype) > 0 {
		ctx.ExporterFp.Write([]byte(string(indent) + `  datatype: ` + node.Datatype + "\n"))
	}
	if len(node.Allowed) > 0 {
		ctx.ExporterFp.Write([]byte(string(indent) + `  allowed: [` + node.Allowed[0] + "]\n"))
	}
	if len(node.Min) > 0 {
		ctx.ExporterFp.Write([]byte(string(indent) + `  min: ` + node.Min + "\n"))
	}
	if len(node.Max) > 0 {
		ctx.ExporterFp.Write([]byte(string(indent) + `  max: ` + node.Max + "\n"))
	}
	if len(node.Unit) > 0 {
		ctx.ExporterFp.Write([]byte(string(indent) + `  unit: ` + node.Unit + "\n"))
	}
	if len(node.Default) > 0 {
		ctx.ExporterFp.Write([]byte(string(indent) + `  default: ` + node.Default + "\n"))
	}
	ctx.ExporterFp.Write([]byte(string(indent) + `  description: ` + node.Description + "\n"))
	ctx.ExporterFp.Write([]byte("\n"))
}

func exportBinary(ctx VspecContext, thisNode PropertyData) {
	if !ctx.FileWrite {
		*(ctx.TmpStorage) = append(*(ctx.TmpStorage), thisNode)
		return
	}
    ctx.ExporterFp.Write(serializeUInt((uint8)(len(thisNode.Name))))
    ctx.ExporterFp.Write([]byte(thisNode.Name))

    ctx.ExporterFp.Write(serializeUInt((uint8)(len(thisNode.NodeType))))
    ctx.ExporterFp.Write([]byte(thisNode.NodeType))

    ctx.ExporterFp.Write(serializeUInt((uint8)(len(thisNode.Uuid))))
    ctx.ExporterFp.Write([]byte(thisNode.Uuid))

    ctx.ExporterFp.Write(serializeUInt((uint16)(len(thisNode.Description))))
    ctx.ExporterFp.Write([]byte(thisNode.Description))

    ctx.ExporterFp.Write(serializeUInt((uint8)(len(thisNode.Datatype))))
    if (len(thisNode.Datatype) > 0) {
        ctx.ExporterFp.Write([]byte(thisNode.Datatype))
    }

    ctx.ExporterFp.Write(serializeUInt((uint8)(len(thisNode.Min))))
    if (len(thisNode.Min) > 0) {
        ctx.ExporterFp.Write([]byte(thisNode.Min))
    }

    ctx.ExporterFp.Write(serializeUInt((uint8)(len(thisNode.Max))))
    if (len(thisNode.Max) > 0) {
        ctx.ExporterFp.Write([]byte(thisNode.Max))
    }

    ctx.ExporterFp.Write(serializeUInt((uint8)(len(thisNode.Unit))))
    if (len(thisNode.Unit) > 0) {
        ctx.ExporterFp.Write([]byte(thisNode.Unit))
    }

    allowedStrLen := calculatAllowedStrLen(thisNode.Allowed)
    ctx.ExporterFp.Write(serializeUInt((uint16)(allowedStrLen)))
    if len(thisNode.Allowed) > 0 {
	for i := 0; i < len(thisNode.Allowed); i++ {
	    allowedWrite(ctx.ExporterFp, thisNode.Allowed[i])
	}
    }

    ctx.ExporterFp.Write(serializeUInt((uint8)(len(thisNode.Default))))
    if len(thisNode.Default) > 0 {
        ctx.ExporterFp.Write([]byte(thisNode.Default))
    }

    Validate := ValidateToString(thisNode.Validate)
    ctx.ExporterFp.Write(serializeUInt((uint8)(len(Validate))))
    if len(Validate) > 0 {
        ctx.ExporterFp.Write([]byte(Validate))
    }

    ctx.ExporterFp.Write(serializeUInt((uint8)(thisNode.Children)))

//    fmt.Printf("exportBinary: %s:%d children\n", thisNode.Name, thisNode.Children)
}

func ValidateToString(validate uint8) string {
    var validation string
    if (validate%10 == 1) {
        validation = "write-only"
    }
    if (validate%10 == 2) {
        validation = "read-write"
    }
    if (validate/10 == 1) {
        validation = "+consent"
    }
    return validation
}

func serializeUInt(intVal interface{}) []byte {
    switch intVal.(type) {
      case uint8:
        buf := make([]byte, 1)
        buf[0] = intVal.(byte)
        return buf
      case uint16:
        buf := make([]byte, 2)
        buf[1] = byte((intVal.(uint16) & 0xFF00)/256)
        buf[0] = byte(intVal.(uint16) & 0x00FF)
        return buf
      case uint32:
        buf := make([]byte, 4)
        buf[3] = byte((intVal.(uint32) & 0xFF000000)/16777216)
        buf[2] = byte((intVal.(uint32) & 0xFF0000)/65536)
        buf[1] = byte((intVal.(uint32) & 0xFF00)/256)
        buf[0] = byte(intVal.(uint32) & 0x00FF)
        return buf
      default:
        fmt.Println(intVal, "is of an unknown type")
        return nil
    }
}

func calculatAllowedStrLen(allowed []string) int {
    strLen := 0
    for i := 0 ; i < len(allowed) ; i++ {
        strLen += len(allowed[i]) + 2
    }
    return strLen
}

func allowedWrite(fp *os.File, allowed string) {
    fp.Write(intToHex(len(allowed)))
//fmt.Printf("allowedHexLen: %s\n", string(intToHex(len(allowed))))
    fp.Write([]byte(allowed))
}

func intToHex(intVal int) []byte {
    if (intVal > 255) {
        return nil
    }
    hexVal := make([]byte, 2)
    hexVal[0] = hexDigit(intVal/16)
    hexVal[1] = hexDigit(intVal%16)
    return hexVal
}

func hexDigit(value int) byte {
    if (value < 10) {
        return byte(value + '0')
    }
    return byte(value - 10 + 'A')
}

func createExporterFile(vspecDir string, exporter string, exportFileName string) {
	var ctx VspecContext
	files, _ := os.ReadDir(vspecDir)
	for _, file := range files {
//		match, _ := filepath.Match("*Specification.vspec", file.Name())
		match, _ := filepath.Match("*.vspec", file.Name())
		if match {
			var exporterFName string
			ctx.Path = vspecDir
			ctx.FName = file.Name()
			ctx.BasePath = ""
			// always create a YAML exporter file first
			ctx.Exporter = exportYaml
			exporterFName = exportFileName + ".yaml"
			exporterFp, err := os.Create(exporterFName)
			if err != nil {
				fmt.Printf("Could not create %s\n", exporterFName)
				return
			}
			ctx.ExporterFp = exporterFp

			writeExporter(ctx)

			exporterFp.Close()
			fmt.Printf("Exporter file %s is created\n", exporterFName)
			exporterMap := createExporterMap(exporterFName)
			var binaryStorage []PropertyData
			switch exporter {
				case "binary":
					ctx.BasePath = ""
					ctx.Exporter = exportBinary
					exporterFName = exportFileName + ".binary"
					ctx.TmpStorage = & binaryStorage
				case "json":
					data, err := json.Marshal(exporterMap)
					if err != nil {
						fmt.Printf("JSON marshall error=%s\n", err)
						return
					}
					exporterFName = exportFileName + ".json"
					ctx.ExporterFp, err = os.Create(exporterFName)
					if err != nil {
						fmt.Printf("Could not create %s\n", exporterFName)
						return
					}
					ctx.ExporterFp.Write(data)
					ctx.ExporterFp.Close()
					fmt.Printf("Exporter file %s is created\n", exporterFName)
//					fmt.Printf("JSON=%s\n", string(data))
					return
				default:
					fmt.Printf("Unsupporter exporter=%s\n", exporter)
					return
			}
			ctx.ExporterFp, err = os.Create(exporterFName)
			if err != nil {
				fmt.Printf("Could not create %s\n", exporterFName)
				return
			}
			mapToExporter(ctx, exporterMap)
			if exporter == "binary" {
				ctx.FileWrite = true
				for i := len(*(ctx.TmpStorage)); i > 0; i-- {
					ctx.Exporter(ctx, (*(ctx.TmpStorage))[i-1])
				}
			}
			ctx.ExporterFp.Close()
			fmt.Printf("Exporter file %s is created\n", exporterFName)
		}
	}
}

func writeExporter(ctx VspecContext) {
	sourceFp, err := os.Open(ctx.Path + ctx.FName)
	if err != nil {
		fmt.Printf("writeExporter:Error reading %s: %s\n", ctx.Path + ctx.FName, err)
		return
	}
	scanner := bufio.NewScanner(sourceFp)
	scanner.Split(bufio.ScanLines)
	var text string
	var nextLine string
	continueScan := true
	var thisNode PropertyData
	var nodeName string
	for continueScan {
		nodeName = ""
		continueScan = scanner.Scan()
		if len(nextLine) > 0 {
			text = nextLine
		} else {
			text = scanner.Text()
		}
		getNodeName(text, &nodeName)
		if len(nodeName) > 0 {
			ctx.Depth = calculateDepth(nodeName, ctx.BasePath)
//fmt.Printf("Node name:%s\n", nodeName)
//fmt.Printf("ctx.BasePath:%s\n", ctx.BasePath)
//fmt.Printf("ctx.Depth:%d\n", ctx.Depth)
			nextLine, thisNode, continueScan = getNode(scanner, nodeName)
			ctx.Exporter(ctx, thisNode)
		} else if len(text) > 8 && strings.Contains(text[:8], "#include") {  // example: #include Vehicle/Vehicle.vspec Vehicle
			var includeCtx VspecContext
			incFields := strings.Fields(text)
			dir, file := filepath.Split(incFields[1])
			includeCtx.FName = file
			if strings.Contains(dir, "include/") && dir[0] == 'i' { // ugly fix...
				dir = "../" + dir
			}
			includeCtx.Path = ctx.Path + dir
			includeCtx.BasePath = ctx.BasePath
			includeCtx.Exporter = ctx.Exporter
			includeCtx.ExporterFp = ctx.ExporterFp
			if len(incFields) > 2 {
				includeCtx.BasePath = ctx.BasePath + "." + incFields[2]
				if includeCtx.BasePath[0] == '.' {
					includeCtx.BasePath = includeCtx.BasePath[1:]
				}
			}
			writeExporter(includeCtx)
		}
	}
	sourceFp.Close()
}

func calculateDepth(nodeName string, basePath string) int {
	dotCount1 := strings.Count(nodeName, ".") + 1
	if len(nodeName) == 0 {
		dotCount1 = 0
	}
	dotCount2 := strings.Count(basePath, ".") + 1
	if len(basePath) == 0 {
		dotCount2 = 0
	}
	return dotCount1 + dotCount2
}

func createExporterMap(exporterFName string) map[string]interface{} {
	var exporterMap map[string]interface{}

	data, err := os.ReadFile(exporterFName)
	if err != nil {
		fmt.Printf("error reading %s\n", exporterFName)
		return nil
	}
	err = yaml.Unmarshal([]byte(data), &exporterMap)
	if err != nil {
		fmt.Printf("yaml unmarshal error: %v\n", err)
		return nil
	}
	return exporterMap
}

func mapToExporter(ctx VspecContext, exporterMap interface{}) {
	switch v := exporterMap.(type) {
	case interface{}:
//		fmt.Println(v, "is interface{}")
		mapToExporterL1(ctx, "", v.(map[string]interface{}))
	default:
		fmt.Println(v, "is of an unknown type")
	}
}

func mapToExporterL1(ctx VspecContext, nodeName string, exporterMapL1 map[string]interface{}) {
	var nodeData PropertyData
	nodeData.Name = nodeName
	for k, v := range exporterMapL1 {
		switch vv := v.(type) {
		case interface{}:
//			fmt.Println(vv, "is interface{}")
			switch vvv := vv.(type) {
			case string:
//				fmt.Println(k,": ", vvv)
				setNodeData(&nodeData, k ,vvv)
			case interface{}:
				mapToExporterL1(ctx, k, vvv.(map[string]interface{}))
				nodeData.Children++
			}
		default:
			fmt.Println(vv, "is of an unknown type")
		}
	}
//fmt.Printf("Node=%s, Children=%d\n", nodeData.Name, nodeData.Children)
	if len(nodeName) > 0 {
		ctx.Exporter(ctx, nodeData)
	}
}

func setNodeData(nodeData *PropertyData, key string,value string) {
	switch key {
		case "type":
			nodeData.NodeType = value
		case "datatype":
			nodeData.Datatype = value
		case "allowed":
			nodeData.Allowed[0] = value
		case "min":
			nodeData.Min = value
		case "max":
			nodeData.Max = value
		case "unit":
			nodeData.Unit = value
		case "default":
			nodeData.Default = value
		case "description":
			nodeData.Description = value
	}
}

func main() {
	parser := argparse.NewParser("print", "HIM tools")
	exporter := parser.Selector("e", "exporter", []string{"yaml", "json", "binary"}, &argparse.Options{Required: false,
		Help: "Exporter parameter must be either: yaml, json, or binary", Default: "yaml"})
//	confFName := parser.String("c", "configfile", &argparse.Options{Required: false, Help: "configuration file name", Default: "himConfig-truck.json"})
	vspecDir := parser.String("r", "rootdir", &argparse.Options{Required: false, Help: "path to vspec root directory", Default: "VehicleServices/"})
//	sConf := parser.Flag("s", "vspecsave", &argparse.Options{Required: false, Help: "Saves the configured .vspec2 files with extension .vspec"})
//	preProcessOnly := parser.Flag("p", "preprocess", &argparse.Options{Required: false, Help: "Pre-process only, save configured vspec files. Do not run VSS-tools."})
//	enumSubst := parser.Flag("n", "noEnumSubst", &argparse.Options{Required: false, Help: "No substitution of enum links to Datatype tree with actual datatypes"})
//	ConfigValueFName := parser.String("v", "configvaluesfile", &argparse.Options{Required: false, Help: "Config values file name"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	createExporterFile(*vspecDir, *exporter, "himExport")
}
