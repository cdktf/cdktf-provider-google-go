// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger


type CloudbuildTriggerBuildArtifactsNpmPackages struct {
	// Path to the package.json. e.g. workspace/path/to/package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.36.0/docs/resources/cloudbuild_trigger#package_path CloudbuildTrigger#package_path}
	PackagePath *string `field:"optional" json:"packagePath" yaml:"packagePath"`
	// Artifact Registry repository, in the form "https://$REGION-npm.pkg.dev/$PROJECT/$REPOSITORY".
	//
	// Npm package in the workspace specified by path will be zipped and uploaded to Artifact Registry with this location as a prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.36.0/docs/resources/cloudbuild_trigger#repository CloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

