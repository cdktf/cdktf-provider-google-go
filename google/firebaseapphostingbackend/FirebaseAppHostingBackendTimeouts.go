// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingbackend


type FirebaseAppHostingBackendTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/firebase_app_hosting_backend#create FirebaseAppHostingBackend#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/firebase_app_hosting_backend#delete FirebaseAppHostingBackend#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/firebase_app_hosting_backend#update FirebaseAppHostingBackend#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

