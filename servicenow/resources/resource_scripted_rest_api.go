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
			scriptedRestApiActive: {
				Type:        schema.TypeBool,
				Optional:    true,
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
				Description: "Override default supported request formats.",
			},
			scriptedRestApiEnforce_acl: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ACLs to enfore when accessing resources. Individual resources may override this value.",
			},
			scriptedRestApiName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the API. Appears in API documentation.",
			},

			scriptedRestApiProduces: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Default supported response formats.",
			},
			scriptedRestApiProduces_customized: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Override default supported response formats.",
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
	data.Set(commonProtectionPolicy, scriptedRestApi.ProtectionPolicy)
	data.Set(commonScope, scriptedRestApi.Scope)
}

func resourceToScriptedRestApi(data *schema.ResourceData) *client.ScriptedRestApi {
	scriptedRestApi := client.ScriptedRestApi{
		Active:             data.Get(scriptedRestApiActive).(bool),
		Consumes:           data.Get(scriptedRestApiConsumes).(string),
		ConsumesCustomized: data.Get(scriptedRestApiConsumes_customized).(bool),
		EnforceACL:         data.Get(scriptedRestApiEnforce_acl).(string),
		Name:               data.Get(scriptedRestApiName).(string),
		Produces:           data.Get(scriptedRestApiProduces).(string),
		ProducesCustomized: data.Get(scriptedRestApiProduces_customized).(bool),
	}
	scriptedRestApi.ID = data.Id()
	scriptedRestApi.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	scriptedRestApi.Scope = data.Get(commonScope).(string)
	return &scriptedRestApi
}
