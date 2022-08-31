// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeInstanceTemplateDiskDiskEncryptionKey struct {
	// The self link of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#kms_key_self_link ComputeInstanceTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"required" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
}

