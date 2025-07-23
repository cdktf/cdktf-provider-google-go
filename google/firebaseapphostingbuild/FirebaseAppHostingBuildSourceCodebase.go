// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingbuild


type FirebaseAppHostingBuildSourceCodebase struct {
	// The branch in the codebase to build from, using the latest commit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firebase_app_hosting_build#branch FirebaseAppHostingBuild#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The commit in the codebase to build from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firebase_app_hosting_build#commit FirebaseAppHostingBuild#commit}
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
}

