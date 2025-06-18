// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingbuild


type FirebaseAppHostingBuildTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/firebase_app_hosting_build#create FirebaseAppHostingBuild#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/firebase_app_hosting_build#delete FirebaseAppHostingBuild#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/firebase_app_hosting_build#update FirebaseAppHostingBuild#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

