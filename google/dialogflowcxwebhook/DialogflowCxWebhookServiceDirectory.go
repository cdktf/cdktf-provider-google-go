package dialogflowcxwebhook


type DialogflowCxWebhookServiceDirectory struct {
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#generic_web_service DialogflowCxWebhook#generic_web_service}
	GenericWebService *DialogflowCxWebhookServiceDirectoryGenericWebService `field:"required" json:"genericWebService" yaml:"genericWebService"`
	// The name of Service Directory service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#service DialogflowCxWebhook#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

