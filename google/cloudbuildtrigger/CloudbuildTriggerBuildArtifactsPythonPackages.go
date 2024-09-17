// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger


type CloudbuildTriggerBuildArtifactsPythonPackages struct {
	// Path globs used to match files in the build's workspace.
	//
	// For Python/ Twine, this is usually dist/*, and sometimes additionally an .asc file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/cloudbuild_trigger#paths CloudbuildTrigger#paths}
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// Artifact Registry repository, in the form "https://$REGION-python.pkg.dev/$PROJECT/$REPOSITORY".
	//
	// Files in the workspace matching any path pattern will be uploaded to Artifact Registry with this location as a prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/cloudbuild_trigger#repository CloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

