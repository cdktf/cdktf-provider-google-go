package beyondcorpappconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpAppConnectionConfig struct {
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
	// application_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#application_endpoint BeyondcorpAppConnection#application_endpoint}
	ApplicationEndpoint *BeyondcorpAppConnectionApplicationEndpoint `field:"required" json:"applicationEndpoint" yaml:"applicationEndpoint"`
	// ID of the AppConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#name BeyondcorpAppConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of AppConnectors that are authorised to be associated with this AppConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#connectors BeyondcorpAppConnection#connectors}
	Connectors *[]*string `field:"optional" json:"connectors" yaml:"connectors"`
	// An arbitrary user-provided name for the AppConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#display_name BeyondcorpAppConnection#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#gateway BeyondcorpAppConnection#gateway}
	Gateway *BeyondcorpAppConnectionGateway `field:"optional" json:"gateway" yaml:"gateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#id BeyondcorpAppConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#labels BeyondcorpAppConnection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#project BeyondcorpAppConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the AppConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#region BeyondcorpAppConnection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#timeouts BeyondcorpAppConnection#timeouts}
	Timeouts *BeyondcorpAppConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of network connectivity used by the AppConnection. Refer to https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type for a list of possible values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connection#type BeyondcorpAppConnection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

