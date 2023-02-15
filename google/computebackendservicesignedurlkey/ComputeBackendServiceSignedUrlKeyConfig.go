package computebackendservicesignedurlkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeBackendServiceSignedUrlKeyConfig struct {
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
	// The backend service this signed URL key belongs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service_signed_url_key#backend_service ComputeBackendServiceSignedUrlKey#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
	// 128-bit key value used for signing the URL.
	//
	// The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service_signed_url_key#key_value ComputeBackendServiceSignedUrlKey#key_value}
	KeyValue *string `field:"required" json:"keyValue" yaml:"keyValue"`
	// Name of the signed URL key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service_signed_url_key#name ComputeBackendServiceSignedUrlKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service_signed_url_key#id ComputeBackendServiceSignedUrlKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service_signed_url_key#project ComputeBackendServiceSignedUrlKey#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service_signed_url_key#timeouts ComputeBackendServiceSignedUrlKey#timeouts}
	Timeouts *ComputeBackendServiceSignedUrlKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

