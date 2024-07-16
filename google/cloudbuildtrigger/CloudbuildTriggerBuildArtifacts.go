// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger


type CloudbuildTriggerBuildArtifacts struct {
	// A list of images to be pushed upon the successful completion of all build steps.
	//
	// The images will be pushed using the builder service account's credentials.
	//
	// The digests of the pushed images will be stored in the Build resource's results field.
	//
	// If any of the images fail to be pushed, the build is marked FAILURE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/cloudbuild_trigger#images CloudbuildTrigger#images}
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	// maven_artifacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/cloudbuild_trigger#maven_artifacts CloudbuildTrigger#maven_artifacts}
	MavenArtifacts interface{} `field:"optional" json:"mavenArtifacts" yaml:"mavenArtifacts"`
	// npm_packages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/cloudbuild_trigger#npm_packages CloudbuildTrigger#npm_packages}
	NpmPackages interface{} `field:"optional" json:"npmPackages" yaml:"npmPackages"`
	// objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/cloudbuild_trigger#objects CloudbuildTrigger#objects}
	Objects *CloudbuildTriggerBuildArtifactsObjects `field:"optional" json:"objects" yaml:"objects"`
	// python_packages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/cloudbuild_trigger#python_packages CloudbuildTrigger#python_packages}
	PythonPackages interface{} `field:"optional" json:"pythonPackages" yaml:"pythonPackages"`
}

