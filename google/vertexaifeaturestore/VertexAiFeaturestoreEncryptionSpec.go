// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeaturestore


type VertexAiFeaturestoreEncryptionSpec struct {
	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	//
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/vertex_ai_featurestore#kms_key_name VertexAiFeaturestore#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

