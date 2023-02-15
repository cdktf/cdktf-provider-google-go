package accesscontextmanagerserviceperimeters

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerServicePerimetersConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The AccessPolicy this ServicePerimeter lives in. Format: accessPolicies/{policy_id}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeters#parent AccessContextManagerServicePerimeters#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeters#id AccessContextManagerServicePerimeters#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// service_perimeters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeters#service_perimeters AccessContextManagerServicePerimeters#service_perimeters}
	ServicePerimeters interface{} `field:"optional" json:"servicePerimeters" yaml:"servicePerimeters"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeters#timeouts AccessContextManagerServicePerimeters#timeouts}
	Timeouts *AccessContextManagerServicePerimetersTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

