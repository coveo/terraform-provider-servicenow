package client

// EndpointScriptedRestApi is the endpoint to manage Scripted Rest Api records.
const EndpointScriptedRestApi = "sys_ws_definition.do"

// ScriptInclude is the json response for a Scripted Rest Api in ServiceNow.
type ScriptedRestApi struct {
	BaseResult
	Active             bool   `json:"active,string"`
	Consumes           string `json:"consumes,omitempty"`
	ConsumesCustomized bool   `json:"consumes_customized,string"`
	EnforceACL         string `json:"enforce_acl,omitempty"`
	Name               string `json:"name"`
	Produces           string `json:"produces,omitempty"`
	ProducesCustomized bool   `json:"produces_customized,string"`
}
