package dialogflowfulfillment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowFulfillmentConfig struct {
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
	// The human-readable name of the fulfillment, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#display_name DialogflowFulfillment#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Whether fulfillment is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#enabled DialogflowFulfillment#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#features DialogflowFulfillment#features}
	Features interface{} `field:"optional" json:"features" yaml:"features"`
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#generic_web_service DialogflowFulfillment#generic_web_service}
	GenericWebService *DialogflowFulfillmentGenericWebService `field:"optional" json:"genericWebService" yaml:"genericWebService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#id DialogflowFulfillment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#project DialogflowFulfillment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#timeouts DialogflowFulfillment#timeouts}
	Timeouts *DialogflowFulfillmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

