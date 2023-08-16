package dataprocautoscalingpolicyiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocAutoscalingPolicyIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_autoscaling_policy_iam_member#member DataprocAutoscalingPolicyIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_autoscaling_policy_iam_member#policy_id DataprocAutoscalingPolicyIamMember#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_autoscaling_policy_iam_member#role DataprocAutoscalingPolicyIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_autoscaling_policy_iam_member#condition DataprocAutoscalingPolicyIamMember#condition}
	Condition *DataprocAutoscalingPolicyIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_autoscaling_policy_iam_member#id DataprocAutoscalingPolicyIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_autoscaling_policy_iam_member#location DataprocAutoscalingPolicyIamMember#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_autoscaling_policy_iam_member#project DataprocAutoscalingPolicyIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

