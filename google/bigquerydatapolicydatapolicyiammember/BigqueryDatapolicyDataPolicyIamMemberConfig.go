package bigquerydatapolicydatapolicyiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryDatapolicyDataPolicyIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_datapolicy_data_policy_iam_member#data_policy_id BigqueryDatapolicyDataPolicyIamMember#data_policy_id}.
	DataPolicyId *string `field:"required" json:"dataPolicyId" yaml:"dataPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_datapolicy_data_policy_iam_member#member BigqueryDatapolicyDataPolicyIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_datapolicy_data_policy_iam_member#role BigqueryDatapolicyDataPolicyIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_datapolicy_data_policy_iam_member#condition BigqueryDatapolicyDataPolicyIamMember#condition}
	Condition *BigqueryDatapolicyDataPolicyIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_datapolicy_data_policy_iam_member#id BigqueryDatapolicyDataPolicyIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_datapolicy_data_policy_iam_member#location BigqueryDatapolicyDataPolicyIamMember#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_datapolicy_data_policy_iam_member#project BigqueryDatapolicyDataPolicyIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

