package beyondcorpappgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpAppGatewayConfig struct {
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
	// ID of the AppGateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#name BeyondcorpAppGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An arbitrary user-provided name for the AppGateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#display_name BeyondcorpAppGateway#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The type of hosting used by the AppGateway. Default value: "HOST_TYPE_UNSPECIFIED" Possible values: ["HOST_TYPE_UNSPECIFIED", "GCP_REGIONAL_MIG"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#host_type BeyondcorpAppGateway#host_type}
	HostType *string `field:"optional" json:"hostType" yaml:"hostType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#id BeyondcorpAppGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#labels BeyondcorpAppGateway#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#project BeyondcorpAppGateway#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the AppGateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#region BeyondcorpAppGateway#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#timeouts BeyondcorpAppGateway#timeouts}
	Timeouts *BeyondcorpAppGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of network connectivity used by the AppGateway. Default value: "TYPE_UNSPECIFIED" Possible values: ["TYPE_UNSPECIFIED", "TCP_PROXY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_gateway#type BeyondcorpAppGateway#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

