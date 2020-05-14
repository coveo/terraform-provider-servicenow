package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const scriptedRestApiActive = "active"
const scriptedRestApiConsumes = "consumes"
const scriptedRestApiConsumes_customized = "consumes_customized"
const scriptedRestApiEnforce_acl = "enforce_acl"
const scriptedRestApiName = "name"
const scriptedRestApiProduces = "produces"
const scriptedRestApiProduces_customized = "produces_customized"
const scriptedRestApiServiceId = "service_id"
const scriptedRestApiBaseURI = "base_uri"
const scriptedRestApiNamespace = "namespace"
const scriptedRestApiDocLink = "doc_link"
const scriptedRestApiShortDescription = "short_description"

// ResourceScriptedRestApi manages a System Property in ServiceNow.
func ResourceScriptedRestApi() *schema.Resource {

	return &schema.Resource{
		Create: createResourceScriptedRestApi,
		Read:   readResourceScriptedRestApi,
		Update: updateResourceScriptedRestApi,
		Delete: deleteResourceScriptedRestApi,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			scriptedRestApiName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the API. Appears in API documentation.",
			},
			scriptedRestApiActive: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Activates the API. Inactive APIs cannot serve requests.",
			},
			scriptedRestApiConsumes: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Default supported request formats.",
			},
			scriptedRestApiConsumes_customized: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates the default supported request formats is customized.",
			},
			scriptedRestApiEnforce_acl: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ACLs to enforce when accessing resources. Individual resources may override this value.",
			},
			scriptedRestApiProduces: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Default supported response formats.",
			},
			scriptedRestApiProduces_customized: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates the default supported request formats is customized.",
			},
			scriptedRestApiServiceId: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The API identifier used to distinguish this API in URI paths. Must be unique within API namespace.",
			},
			scriptedRestApiBaseURI: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The base API path (URI) to access this API.",
			},
			scriptedRestApiNamespace: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The namespace the API belongs to. The value depends on the current application scope.",
			},
			scriptedRestApiDocLink: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies a URL that links to static documentation about the API.",
			},
			scriptedRestApiShortDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the API. Appears in API documentation.",
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
			commonScope:            getScopeSchema(),
		},
	}
}

func readResourceScriptedRestApi(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	scriptedRestApi := &client.ScriptedRestApi{}
	if err := snowClient.GetObject(client.EndpointScriptedRestApi, data.Id(), scriptedRestApi); err != nil {
		data.SetId("")
		return err
	}

	resourceFromScriptedRestApi(data, scriptedRestApi)

	return nil
}

func createResourceScriptedRestApi(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	scriptedRestApi := resourceToScriptedRestApi(data)
	if err := snowClient.CreateObject(client.EndpointScriptedRestApi, scriptedRestApi); err != nil {
		return err
	}

	resourceFromScriptedRestApi(data, scriptedRestApi)

	return readResourceScriptedRestApi(data, serviceNowClient)
}

func updateResourceScriptedRestApi(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointScriptedRestApi, resourceToScriptedRestApi(data)); err != nil {
		return err
	}

	return readResourceScriptedRestApi(data, serviceNowClient)
}

func deleteResourceScriptedRestApi(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointScriptedRestApi, data.Id())
}

func resourceFromScriptedRestApi(data *schema.ResourceData, scriptedRestApi *client.ScriptedRestApi) {
	data.SetId(scriptedRestApi.ID)
	data.Set(scriptedRestApiActive, scriptedRestApi.Active)
	data.Set(scriptedRestApiConsumes, scriptedRestApi.Consumes)
	data.Set(scriptedRestApiConsumes_customized, scriptedRestApi.ConsumesCustomized)
	data.Set(scriptedRestApiEnforce_acl, scriptedRestApi.EnforceACL)
	data.Set(scriptedRestApiName, scriptedRestApi.Name)
	data.Set(scriptedRestApiProduces, scriptedRestApi.Produces)
	data.Set(scriptedRestApiProduces_customized, scriptedRestApi.ProducesCustomized)
	data.Set(scriptedRestApiServiceId, scriptedRestApi.ServiceId)
	data.Set(scriptedRestApiBaseURI, scriptedRestApi.BaseURI)
	data.Set(scriptedRestApiNamespace, scriptedRestApi.Namespace)
	data.Set(scriptedRestApiDocLink, scriptedRestApi.DocLink)
	data.Set(scriptedRestApiShortDescription, scriptedRestApi.ShortDescription)
	data.Set(commonProtectionPolicy, scriptedRestApi.ProtectionPolicy)
	data.Set(commonScope, scriptedRestApi.Scope)
}

func resourceToScriptedRestApi(data *schema.ResourceData) *client.ScriptedRestApi {
	scriptedRestApi := client.ScriptedRestApi{
		Active:             data.Get(scriptedRestApiActive).(bool),
		EnforceACL:         data.Get(scriptedRestApiEnforce_acl).(string),
		Name:               data.Get(scriptedRestApiName).(string),
		Produces:           data.Get(scriptedRestApiProduces).(string),
		ProducesCustomized: data.Get(scriptedRestApiProduces_customized).(bool),
		ServiceId:          data.Get(scriptedRestApiServiceId).(string),
		BaseURI:            data.Get(scriptedRestApiBaseURI).(string),
		Namespace:          data.Get(scriptedRestApiNamespace).(string),
		DocLink:            data.Get(scriptedRestApiDocLink).(string),
		ShortDescription:   data.Get(scriptedRestApiShortDescription).(string),
	}

	consumes, consumesOk := data.GetOk(scriptedRestApiConsumes)
	scriptedRestApi.Consumes = consumes.(string)
	//ConsumesCustomized is inferred from Consumes. If Consumes is empty, ConsumesCustomized is false.
	if consumesOk {
		scriptedRestApi.ConsumesCustomized = true
	} else {
		scriptedRestApi.ConsumesCustomized = false
	}

	produces, producesOk := data.GetOk(scriptedRestApiProduces)
	scriptedRestApi.Produces = produces.(string)
	//ProducesCustomized is inferred from Produces. If Produces is empty, ProducesCustomized is false.
	if producesOk {
		scriptedRestApi.ProducesCustomized = true
	} else {
		scriptedRestApi.ProducesCustomized = false
	}

	scriptedRestApi.ID = data.Id()
	scriptedRestApi.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	scriptedRestApi.Scope = data.Get(commonScope).(string)
	return &scriptedRestApi
}
