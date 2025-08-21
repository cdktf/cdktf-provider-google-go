// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaserulesrelease


type FirebaserulesReleaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/firebaserules_release#create FirebaserulesRelease#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/firebaserules_release#delete FirebaserulesRelease#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

