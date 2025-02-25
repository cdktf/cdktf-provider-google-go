// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firestoreindex


type FirestoreIndexFieldsVectorConfig struct {
	// The resulting index will only include vectors of this dimension, and can be used for vector search with the same dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/firestore_index#dimension FirestoreIndex#dimension}
	Dimension *float64 `field:"optional" json:"dimension" yaml:"dimension"`
	// flat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/firestore_index#flat FirestoreIndex#flat}
	Flat *FirestoreIndexFieldsVectorConfigFlat `field:"optional" json:"flat" yaml:"flat"`
}

