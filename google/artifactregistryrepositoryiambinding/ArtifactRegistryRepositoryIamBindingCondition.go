// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepositoryiambinding


type ArtifactRegistryRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/artifact_registry_repository_iam_binding#expression ArtifactRegistryRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/artifact_registry_repository_iam_binding#title ArtifactRegistryRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/artifact_registry_repository_iam_binding#description ArtifactRegistryRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

