// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firestorefield


type FirestoreFieldTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/firestore_field#create FirestoreField#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/firestore_field#delete FirestoreField#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/firestore_field#update FirestoreField#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

