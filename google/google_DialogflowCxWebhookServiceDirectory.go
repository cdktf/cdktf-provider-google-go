// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DialogflowCxWebhookServiceDirectory struct {
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_webhook#generic_web_service DialogflowCxWebhook#generic_web_service}
	GenericWebService *DialogflowCxWebhookServiceDirectoryGenericWebService `field:"required" json:"genericWebService" yaml:"genericWebService"`
	// The name of Service Directory service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_webhook#service DialogflowCxWebhook#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

