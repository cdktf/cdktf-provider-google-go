package vertexaitensorboard


type VertexAiTensorboardEncryptionSpec struct {
	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	//
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_tensorboard#kms_key_name VertexAiTensorboard#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

