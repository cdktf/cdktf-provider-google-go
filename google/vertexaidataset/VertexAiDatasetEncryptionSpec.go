// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaidataset


type VertexAiDatasetEncryptionSpec struct {
	// Required.
	//
	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/vertex_ai_dataset#kms_key_name VertexAiDataset#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

