package apigeeflowhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeFlowhookConfig struct {
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
	// The resource ID of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_flowhook#environment ApigeeFlowhook#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Where in the API call flow the flow hook is invoked.
	//
	// Must be one of PreProxyFlowHook, PostProxyFlowHook, PreTargetFlowHook, or PostTargetFlowHook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_flowhook#flow_hook_point ApigeeFlowhook#flow_hook_point}
	FlowHookPoint *string `field:"required" json:"flowHookPoint" yaml:"flowHookPoint"`
	// The Apigee Organization associated with the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_flowhook#org_id ApigeeFlowhook#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Id of the Sharedflow attaching to a flowhook point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_flowhook#sharedflow ApigeeFlowhook#sharedflow}
	Sharedflow *string `field:"required" json:"sharedflow" yaml:"sharedflow"`
	// Flag that specifies whether execution should continue if the flow hook throws an exception.
	//
	// Set to true to continue execution. Set to false to stop execution if the flow hook throws an exception. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_flowhook#continue_on_error ApigeeFlowhook#continue_on_error}
	ContinueOnError interface{} `field:"optional" json:"continueOnError" yaml:"continueOnError"`
	// Description of the flow hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_flowhook#description ApigeeFlowhook#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_flowhook#id ApigeeFlowhook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_flowhook#timeouts ApigeeFlowhook#timeouts}
	Timeouts *ApigeeFlowhookTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

