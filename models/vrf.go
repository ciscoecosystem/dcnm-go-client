package models

type VRF struct {
	Fabric             string `json:",omitempty"`
	Name               string `json:",omitempty"`
	Id                 string `json:",omitempty"`
	Template           string `json:",omitempty"`
	Config             string `json:",omitempty"`
	ExtensionTemplate  string `json:",omitempty"`
	ServiceVRFTemplate string `json:",omitempty"`
	Source             string `json:",omitempty"`
}

func (vrf *VRF) ToMap() (map[string]interface{}, error) {
	vrfAttributeMap := make(map[string]interface{})
	A(vrfAttributeMap, "fabric", vrf.Fabric)
	A(vrfAttributeMap, "vrfName", vrf.Name)
	A(vrfAttributeMap, "vrfId", vrf.Id)
	A(vrfAttributeMap, "vrfTemplate", vrf.Template)
	A(vrfAttributeMap, "vrfTemplateConfig", vrf.Config)
	if vrf.ExtensionTemplate != "" {
		A(vrfAttributeMap, "vrfExtensionTemplate", vrf.ExtensionTemplate)
	}
	if vrf.ServiceVRFTemplate != "" {
		A(vrfAttributeMap, "serviceVrfTemplate", vrf.ServiceVRFTemplate)
	}
	if vrf.Source != "" {
		A(vrfAttributeMap, "source", vrf.Source)
	}
	return vrfAttributeMap, nil
}
