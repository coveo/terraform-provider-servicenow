package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const applicationName = "name"
const applicationScope = "scope"
const applicationVersion = "version"

// DataSourceApplication reads an Application in ServiceNow.
func DataSourceApplication() *schema.Resource {
	return &schema.Resource{
		Read: readDataSourceApplication,

		Schema: map[string]*schema.Schema{
			applicationName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Application to retrieve from the ServiceNow instance.",
			},
			applicationScope: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique scope of the application. Normally in the format x_[companyCode]_[shortAppId].",
			},
			applicationVersion: {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func readDataSourceApplication(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	application, err := client.GetApplicationByName(data.Get(applicationName).(string))
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromApplication(data, application)

	return nil
}

func resourceFromApplication(data *schema.ResourceData, application *client.Application) {
	data.SetId(application.Id)
	data.Set(applicationName, application.Name)
	data.Set(applicationScope, application.Scope)
	data.Set(applicationVersion, application.Version)
}
