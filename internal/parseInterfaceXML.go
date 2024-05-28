package internal

import (
	"encoding/xml"
	"fmt"
)

func ParseInterfaceXML(xmlData []byte, nodeName string) ([]InterfaceData, error) {
	var rpcReply RpcReply
	err := xml.Unmarshal(xmlData, &rpcReply)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %v", err)
	}

	var interfaceDataList []InterfaceData
	for _, iface := range rpcReply.InterfaceInformation.PhysicalInterface {
		data := InterfaceData{
			Node:        nodeName,
			Interface:   iface.Name,
			Description: iface.Description,
			Status:      iface.OperStatus,
			LastFlapped: iface.InterfaceFlapped.Text,
		}
		interfaceDataList = append(interfaceDataList, data)
	}

	return interfaceDataList, nil
}
