// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseappcheckserviceconfig


type FirebaseAppCheckServiceConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/firebase_app_check_service_config#create FirebaseAppCheckServiceConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/firebase_app_check_service_config#delete FirebaseAppCheckServiceConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/firebase_app_check_service_config#update FirebaseAppCheckServiceConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

