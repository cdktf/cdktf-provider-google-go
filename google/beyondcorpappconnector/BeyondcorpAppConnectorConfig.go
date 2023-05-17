package beyondcorpappconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpAppConnectorConfig struct {
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
	// ID of the AppConnector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#name BeyondcorpAppConnector#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// principal_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#principal_info BeyondcorpAppConnector#principal_info}
	PrincipalInfo *BeyondcorpAppConnectorPrincipalInfo `field:"required" json:"principalInfo" yaml:"principalInfo"`
	// An arbitrary user-provided name for the AppConnector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#display_name BeyondcorpAppConnector#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#id BeyondcorpAppConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#labels BeyondcorpAppConnector#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#project BeyondcorpAppConnector#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the AppConnector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#region BeyondcorpAppConnector#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#timeouts BeyondcorpAppConnector#timeouts}
	Timeouts *BeyondcorpAppConnectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

