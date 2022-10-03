package datagooglekmscryptokey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleKmsCryptoKeyConfig struct {
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
	// The KeyRing that this key belongs to. Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/kms_crypto_key#key_ring DataGoogleKmsCryptoKey#key_ring}
	KeyRing *string `field:"required" json:"keyRing" yaml:"keyRing"`
	// The resource name for the CryptoKey.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/kms_crypto_key#name DataGoogleKmsCryptoKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/kms_crypto_key#id DataGoogleKmsCryptoKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

