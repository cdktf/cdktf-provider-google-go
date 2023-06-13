package kmskeyring


type KmsKeyRingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/kms_key_ring#create KmsKeyRing#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/kms_key_ring#delete KmsKeyRing#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

