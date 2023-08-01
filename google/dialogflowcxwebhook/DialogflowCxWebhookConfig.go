package dialogflowcxwebhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxWebhookConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The human-readable name of the webhook, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#display_name DialogflowCxWebhook#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Indicates whether the webhook is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#disabled DialogflowCxWebhook#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#enable_spell_correction DialogflowCxWebhook#enable_spell_correction}
	EnableSpellCorrection interface{} `field:"optional" json:"enableSpellCorrection" yaml:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#enable_stackdriver_logging DialogflowCxWebhook#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#generic_web_service DialogflowCxWebhook#generic_web_service}
	GenericWebService *DialogflowCxWebhookGenericWebService `field:"optional" json:"genericWebService" yaml:"genericWebService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#id DialogflowCxWebhook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The agent to create a webhook for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#parent DialogflowCxWebhook#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#security_settings DialogflowCxWebhook#security_settings}
	SecuritySettings *string `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// service_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#service_directory DialogflowCxWebhook#service_directory}
	ServiceDirectory *DialogflowCxWebhookServiceDirectory `field:"optional" json:"serviceDirectory" yaml:"serviceDirectory"`
	// Webhook execution timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#timeout DialogflowCxWebhook#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_webhook#timeouts DialogflowCxWebhook#timeouts}
	Timeouts *DialogflowCxWebhookTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

