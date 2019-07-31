package client

// EndpointExtensionPoint is the endpoint to manage JS Include records.
const EndpointExtensionPoint = "sp_js_include.do"

// ExtensionPoint represents the json response for a Js Include in ServiceNow.
type ExtensionPoint struct {
	BaseResult
	Name          string `json:"name"`
	RestrictScope bool   `json:"restrict_scope,string"`
	Description   string `json:"description"`
	Example       string `json:"example"`
	APIName       string `json:"api_name,omitempty"`
}
