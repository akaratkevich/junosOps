package internal

import "encoding/xml"

// RpcReply was generated 2024-05-28 10:11:29 by https://xml-to-go.github.io/
type RpcReply struct {
	XMLName              xml.Name `xml:"rpc-reply"`
	Text                 string   `xml:",chardata"`
	Junos                string   `xml:"junos,attr"`
	InterfaceInformation struct {
		Text              string `xml:",chardata"`
		Xmlns             string `xml:"xmlns,attr"`
		Style             string `xml:"style,attr"`
		PhysicalInterface []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name"`
			AdminStatus struct {
				Text   string `xml:",chardata"`
				Format string `xml:"format,attr"`
			} `xml:"admin-status"`
			OperStatus            string `xml:"oper-status"`
			LocalIndex            string `xml:"local-index"`
			SnmpIndex             string `xml:"snmp-index"`
			Description           string `xml:"description"`
			LinkLevelType         string `xml:"link-level-type"`
			Mtu                   string `xml:"mtu"`
			SonetMode             string `xml:"sonet-mode"`
			Mru                   string `xml:"mru"`
			SourceFiltering       string `xml:"source-filtering"`
			Speed                 string `xml:"speed"`
			EthSwitchError        string `xml:"eth-switch-error"`
			BpduError             string `xml:"bpdu-error"`
			LdPduError            string `xml:"ld-pdu-error"`
			L2ptError             string `xml:"l2pt-error"`
			Loopback              string `xml:"loopback"`
			IfFlowControl         string `xml:"if-flow-control"`
			IfAutoNegotiation     string `xml:"if-auto-negotiation"`
			IfRemoteFault         string `xml:"if-remote-fault"`
			PadToMinimumFrameSize string `xml:"pad-to-minimum-frame-size"`
			IfDeviceFlags         struct {
				Text        string `xml:",chardata"`
				IfdfPresent string `xml:"ifdf-present"`
				IfdfRunning string `xml:"ifdf-running"`
			} `xml:"if-device-flags"`
			IfdSpecificConfigFlags struct {
				Text          string `xml:",chardata"`
				InternalFlags string `xml:"internal-flags"`
			} `xml:"ifd-specific-config-flags"`
			IfConfigFlags struct {
				Text          string `xml:",chardata"`
				IffSnmpTraps  string `xml:"iff-snmp-traps"`
				InternalFlags string `xml:"internal-flags"`
			} `xml:"if-config-flags"`
			IfMediaFlags struct {
				Text     string `xml:",chardata"`
				IfmfNone string `xml:"ifmf-none"`
			} `xml:"if-media-flags"`
			PhysicalInterfaceCosInformation struct {
				Text                             string `xml:",chardata"`
				PhysicalInterfaceCosHwMaxQueues  string `xml:"physical-interface-cos-hw-max-queues"`
				PhysicalInterfaceCosUseMaxQueues string `xml:"physical-interface-cos-use-max-queues"`
			} `xml:"physical-interface-cos-information"`
			CurrentPhysicalAddress  string `xml:"current-physical-address"`
			HardwarePhysicalAddress string `xml:"hardware-physical-address"`
			InterfaceFlapped        struct {
				Text    string `xml:",chardata"`
				Seconds string `xml:"seconds,attr"`
			} `xml:"interface-flapped"`
			TrafficStatistics struct {
				Text      string `xml:",chardata"`
				Style     string `xml:"style,attr"`
				InputBps  string `xml:"input-bps"`
				InputPps  string `xml:"input-pps"`
				OutputBps string `xml:"output-bps"`
				OutputPps string `xml:"output-pps"`
			} `xml:"traffic-statistics"`
			ActiveAlarms struct {
				Text            string `xml:",chardata"`
				InterfaceAlarms struct {
					Text            string `xml:",chardata"`
					AlarmNotPresent string `xml:"alarm-not-present"`
				} `xml:"interface-alarms"`
			} `xml:"active-alarms"`
			ActiveDefects struct {
				Text            string `xml:",chardata"`
				InterfaceAlarms struct {
					Text            string `xml:",chardata"`
					AlarmNotPresent string `xml:"alarm-not-present"`
				} `xml:"interface-alarms"`
			} `xml:"active-defects"`
			EthernetPcsStatistics struct {
				Text                 string `xml:",chardata"`
				Style                string `xml:"style,attr"`
				BitErrorSeconds      string `xml:"bit-error-seconds"`
				ErroredBlocksSeconds string `xml:"errored-blocks-seconds"`
			} `xml:"ethernet-pcs-statistics"`
			EthernetFecMode struct {
				Text           string `xml:",chardata"`
				Style          string `xml:"style,attr"`
				EnabledFecMode string `xml:"enabled_fec_mode"`
			} `xml:"ethernet-fec-mode"`
			EthernetFecStatistics struct {
				Text             string `xml:",chardata"`
				Style            string `xml:"style,attr"`
				FecCcwCount      string `xml:"fec_ccw_count"`
				FecNccwCount     string `xml:"fec_nccw_count"`
				FecCcwErrorRate  string `xml:"fec_ccw_error_rate"`
				FecNccwErrorRate string `xml:"fec_nccw_error_rate"`
			} `xml:"ethernet-fec-statistics"`
			InterfaceTransmitStatistics string `xml:"interface-transmit-statistics"`
			LogicalInterface            struct {
				Text          string `xml:",chardata"`
				Name          string `xml:"name"`
				LocalIndex    string `xml:"local-index"`
				SnmpIndex     string `xml:"snmp-index"`
				IfConfigFlags struct {
					Text          string `xml:",chardata"`
					IffUp         string `xml:"iff-up"`
					IffSnmpTraps  string `xml:"iff-snmp-traps"`
					InternalFlags string `xml:"internal-flags"`
				} `xml:"if-config-flags"`
				Encapsulation     string `xml:"encapsulation"`
				PolicerOverhead   string `xml:"policer-overhead"`
				TrafficStatistics struct {
					Text          string `xml:",chardata"`
					Style         string `xml:"style,attr"`
					InputPackets  string `xml:"input-packets"`
					OutputPackets string `xml:"output-packets"`
				} `xml:"traffic-statistics"`
				FilterInformation string `xml:"filter-information"`
				AddressFamily     []struct {
					Text               string `xml:",chardata"`
					AddressFamilyName  string `xml:"address-family-name"`
					Mtu                string `xml:"mtu"`
					MaxLocalCache      string `xml:"max-local-cache"`
					NewHoldLimit       string `xml:"new-hold-limit"`
					IntfCurrCnt        string `xml:"intf-curr-cnt"`
					IntfUnresolvedCnt  string `xml:"intf-unresolved-cnt"`
					IntfDropcnt        string `xml:"intf-dropcnt"`
					AddressFamilyFlags struct {
						Text                 string `xml:",chardata"`
						IfffIsPrimary        string `xml:"ifff-is-primary"`
						IfffSendbcastPktToRe string `xml:"ifff-sendbcast-pkt-to-re"`
						InternalFlags        string `xml:"internal-flags"`
					} `xml:"address-family-flags"`
					InterfaceAddress struct {
						Text     string `xml:",chardata"`
						IfaFlags struct {
							Text                 string `xml:",chardata"`
							IfafCurrentPreferred string `xml:"ifaf-current-preferred"`
							IfafCurrentPrimary   string `xml:"ifaf-current-primary"`
						} `xml:"ifa-flags"`
						IfaDestination string `xml:"ifa-destination"`
						IfaLocal       string `xml:"ifa-local"`
						IfaBroadcast   string `xml:"ifa-broadcast"`
					} `xml:"interface-address"`
				} `xml:"address-family"`
			} `xml:"logical-interface"`
		} `xml:"physical-interface"`
	} `xml:"interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
