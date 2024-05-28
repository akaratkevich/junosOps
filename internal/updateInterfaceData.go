package internal

import "encoding/xml"

func UpdateInterfaceData(data *InterfaceData, xmlData []byte) error {
	var rpcReply RpcReply
	err := xml.Unmarshal(xmlData, &rpcReply)
	if err != nil {
		return err
	}

	for _, iface := range rpcReply.InterfaceInformation.PhysicalInterface {
		data.Interface = iface.Name
		data.Description = iface.Description
		data.LastFlapped = iface.InterfaceFlapped.Text
	}
	return nil
}
