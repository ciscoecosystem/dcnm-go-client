package models

type ServicePolicy struct {
	PolicyName             		    string `json:",omitempty"`
	FabricName               		string `json:",omitempty"`
	AttachedFabricName              string `json:",omitempty"`
	DestinationNetwork           	string `json:",omitempty"`
	DestinationNetworkName          string `json:",omitempty"`
	DestinationVRFName 				string `json:",omitempty"`
	Enabled 						bool `json:",omitempty"`
	LastUpdate             			string `json:",omitempty"`
	NextHopIP          				string `json:",omitempty"`
	PeeringName 					string `json:",omitempty"`
	PolicyTemplateName 				string `json:",omitempty"`
	ReverseEnabled             		bool `json:",omitempty"`
	ReverseNextHopIP             	string `json:",omitempty"`
	RouteMapName               		string `json:",omitempty"`
	ServiceNodeName              	string `json:",omitempty"`
	ServiceNodeType           		string `json:",omitempty"`
	SourceNetwork          			string `json:",omitempty"`
	SourceNetworkName 				string `json:",omitempty"`
	SourceVRFName        			string `json:",omitempty"`
	Status 							string `json:",omitempty"`
	StatusDetails           		interface `json:",omitempty"`
	AttachDetails          			interface `json:",omitempty"`
	DestinationInterfaces 			interface `json:",omitempty"`
	SourceInterfaces        		interface `json:",omitempty"`
	NVPairs							interface `json:",omitempty"`

}

func (servicepolicy *ServicePolicy) ToMap() (map[string]interface{}, error) {
	servicepolicyAttributeMap := make(map[string]interface{})
	A(servicepolicyAttributeMap, "policyName", servicepolicy.PolicyName)
	A(servicepolicyAttributeMap, "fabricName", servicepolicy.FabricName)
	A(servicepolicyAttributeMap, "attachedFabricName", servicepolicy.AttachedFabricName)
	A(servicepolicyAttributeMap, "destinationNetwork", servicepolicy.DestinationNetwork)
	A(servicepolicyAttributeMap, "destinationNetworkName", servicepolicy.DestinationNetworkName)
	A(servicepolicyAttributeMap, "destinationVRFName", servicepolicy.DestinationVRFName)
	A(servicepolicyAttributeMap, "enabled", servicepolicy.Enabled)
	A(servicepolicyAttributeMap, "lastUpdate", servicepolicy.LastUpdate)
	A(servicepolicyAttributeMap, "nextHopIP", servicepolicy.NextHopIP)
	A(servicepolicyAttributeMap, "peeringName", servicepolicy.PeeringName)
	A(servicepolicyAttributeMap, "policyTemplateName", servicepolicy.PolicyTemplateName)
	A(servicepolicyAttributeMap, "reverseEnabled", servicepolicy.ReverseEnabled)
	A(servicepolicyAttributeMap, "reverseNextHopIP", servicepolicy.ReverseNextHopIP)
	A(servicepolicyAttributeMap, "routeMapName", servicepolicy.RouteMapName)
	A(servicepolicyAttributeMap, "serviceNodeName", servicepolicy.ServiceNodeName)
	A(servicepolicyAttributeMap, "serviceNodeType", servicepolicy.ServiceNodeType)
	A(servicepolicyAttributeMap, "sourceNetwork", servicepolicy.SourceNetwork)
	A(servicepolicyAttributeMap, "sourceNetworkName", servicepolicy.SourceNetworkName)
	A(servicepolicyAttributeMap, "sourceVRFName", servicepolicy.SourceVRFName)
	A(servicepolicyAttributeMap, "status", servicepolicy.Status)
	A(servicepolicyAttributeMap, "statusDetails", servicepolicy.StatusDetails)
	A(servicepolicyAttributeMap, "attachDetails", servicepolicy.AttachDetails)
	A(servicepolicyAttributeMap, "destinationInterfaces", servicepolicy.DestinationInterfaces)
	A(servicepolicyAttributeMap, "sourceInterfaces", servicepolicy.SourceInterfaces)
	A(servicepolicyAttributeMap, "nVPairs", servicepolicy.NVPairs)

	return servicenodeAttributeMap,nil
}	
