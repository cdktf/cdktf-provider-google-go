package apigeeinstanceattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeInstanceAttachmentConfig struct {
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
	// The resource ID of the environment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_instance_attachment#environment ApigeeInstanceAttachment#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The Apigee instance associated with the Apigee environment, in the format 'organisations/{{org_name}}/instances/{{instance_name}}'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_instance_attachment#instance_id ApigeeInstanceAttachment#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_instance_attachment#id ApigeeInstanceAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_instance_attachment#timeouts ApigeeInstanceAttachment#timeouts}
	Timeouts *ApigeeInstanceAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

