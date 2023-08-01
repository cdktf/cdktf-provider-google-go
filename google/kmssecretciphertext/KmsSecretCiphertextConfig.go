package kmssecretciphertext

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsSecretCiphertextConfig struct {
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
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/kms_secret_ciphertext#crypto_key KmsSecretCiphertext#crypto_key}
	CryptoKey *string `field:"required" json:"cryptoKey" yaml:"cryptoKey"`
	// The plaintext to be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/kms_secret_ciphertext#plaintext KmsSecretCiphertext#plaintext}
	Plaintext *string `field:"required" json:"plaintext" yaml:"plaintext"`
	// The additional authenticated data used for integrity checks during encryption and decryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/kms_secret_ciphertext#additional_authenticated_data KmsSecretCiphertext#additional_authenticated_data}
	AdditionalAuthenticatedData *string `field:"optional" json:"additionalAuthenticatedData" yaml:"additionalAuthenticatedData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/kms_secret_ciphertext#id KmsSecretCiphertext#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/kms_secret_ciphertext#timeouts KmsSecretCiphertext#timeouts}
	Timeouts *KmsSecretCiphertextTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

