package accesscontextmanageregresspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerEgressPolicyConfig struct {
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
	// The name of the Service Perimeter to add this resource to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_egress_policy#egress_policy_name AccessContextManagerEgressPolicy#egress_policy_name}
	EgressPolicyName *string `field:"required" json:"egressPolicyName" yaml:"egressPolicyName"`
	// A GCP resource that is inside of the service perimeter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_egress_policy#resource AccessContextManagerEgressPolicy#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_egress_policy#id AccessContextManagerEgressPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_egress_policy#timeouts AccessContextManagerEgressPolicy#timeouts}
	Timeouts *AccessContextManagerEgressPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

