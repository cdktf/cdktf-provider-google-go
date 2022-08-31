// Prebuilt google Provider for Terraform CDK (cdktf)
package google

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleServiceAccountJwtConfig struct {
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
	// A JSON-encoded JWT claims set that will be included in the signed JWT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/service_account_jwt#payload DataGoogleServiceAccountJwt#payload}
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/service_account_jwt#target_service_account DataGoogleServiceAccountJwt#target_service_account}.
	TargetServiceAccount *string `field:"required" json:"targetServiceAccount" yaml:"targetServiceAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/service_account_jwt#delegates DataGoogleServiceAccountJwt#delegates}.
	Delegates *[]*string `field:"optional" json:"delegates" yaml:"delegates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/service_account_jwt#id DataGoogleServiceAccountJwt#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

