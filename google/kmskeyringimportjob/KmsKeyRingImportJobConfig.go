package kmskeyringimportjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsKeyRingImportJobConfig struct {
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
	// It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/kms_key_ring_import_job#import_job_id KmsKeyRingImportJob#import_job_id}
	ImportJobId *string `field:"required" json:"importJobId" yaml:"importJobId"`
	// The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256", "RSA_OAEP_4096_SHA1_AES_256"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/kms_key_ring_import_job#import_method KmsKeyRingImportJob#import_method}
	ImportMethod *string `field:"required" json:"importMethod" yaml:"importMethod"`
	// The KeyRing that this import job belongs to. Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/kms_key_ring_import_job#key_ring KmsKeyRingImportJob#key_ring}
	KeyRing *string `field:"required" json:"keyRing" yaml:"keyRing"`
	// The protection level of the ImportJob.
	//
	// This must match the protectionLevel of the
	// versionTemplate on the CryptoKey you attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/kms_key_ring_import_job#protection_level KmsKeyRingImportJob#protection_level}
	ProtectionLevel *string `field:"required" json:"protectionLevel" yaml:"protectionLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/kms_key_ring_import_job#id KmsKeyRingImportJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/kms_key_ring_import_job#timeouts KmsKeyRingImportJob#timeouts}
	Timeouts *KmsKeyRingImportJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

