package apigeesharedflowdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeSharedflowDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_sharedflow_deployment#environment ApigeeSharedflowDeployment#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The Apigee Organization associated with the Apigee instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_sharedflow_deployment#org_id ApigeeSharedflowDeployment#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Revision of the Sharedflow to be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_sharedflow_deployment#revision ApigeeSharedflowDeployment#revision}
	Revision *string `field:"required" json:"revision" yaml:"revision"`
	// Id of the Sharedflow to be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_sharedflow_deployment#sharedflow_id ApigeeSharedflowDeployment#sharedflow_id}
	SharedflowId *string `field:"required" json:"sharedflowId" yaml:"sharedflowId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_sharedflow_deployment#id ApigeeSharedflowDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The service account represents the identity of the deployed proxy, and determines what permissions it has.
	//
	// The format must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_sharedflow_deployment#service_account ApigeeSharedflowDeployment#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_sharedflow_deployment#timeouts ApigeeSharedflowDeployment#timeouts}
	Timeouts *ApigeeSharedflowDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

