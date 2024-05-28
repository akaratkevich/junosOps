package internal

import (
	"encoding/xml"
)

func ParseInterfaceXML(xmlData []byte, nodeName string) ([]InterfaceData, error) {
	var rpcReply RpcReply
	err := xml.Unmarshal(xmlData, &rpcReply)
	if err != nil {
		return nil, err
	}

	var interfaces []InterfaceData
	for _, iface := range rpcReply.InterfaceInformation.PhysicalInterface {
		data := InterfaceData{
			Node:        nodeName,
			Interface:   iface.Name,
			Description: iface.Description,
			LastFlapped: iface.InterfaceFlapped.Text,
		}
		interfaces = append(interfaces, data)
	}
	return interfaces, nil
}
