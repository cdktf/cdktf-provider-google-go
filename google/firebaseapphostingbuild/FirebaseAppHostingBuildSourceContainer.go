// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingbuild


type FirebaseAppHostingBuildSourceContainer struct {
	// A URI representing a container for the backend to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/firebase_app_hosting_build#image FirebaseAppHostingBuild#image}
	Image *string `field:"required" json:"image" yaml:"image"`
}

