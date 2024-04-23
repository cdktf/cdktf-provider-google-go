// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger


type CloudbuildTriggerBuildAvailableSecretsSecretManager struct {
	// Environment variable name to associate with the secret.
	//
	// Secret environment
	// variables must be unique across all of a build's secrets, and must be used
	// by at least one build step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/cloudbuild_trigger#env CloudbuildTrigger#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// Resource name of the SecretVersion. In format: projects/* /secrets/* /versions/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/cloudbuild_trigger#version_name CloudbuildTrigger#version_name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
}

