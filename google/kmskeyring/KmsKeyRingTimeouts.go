package kmskeyring


type KmsKeyRingTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/kms_key_ring#create KmsKeyRing#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/kms_key_ring#delete KmsKeyRing#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

