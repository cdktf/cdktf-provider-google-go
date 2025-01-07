// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryCleanupPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/artifact_registry_repository#id ArtifactRegistryRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Policy action. Possible values: ["DELETE", "KEEP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/artifact_registry_repository#action ArtifactRegistryRepository#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/artifact_registry_repository#condition ArtifactRegistryRepository#condition}
	Condition *ArtifactRegistryRepositoryCleanupPoliciesCondition `field:"optional" json:"condition" yaml:"condition"`
	// most_recent_versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/artifact_registry_repository#most_recent_versions ArtifactRegistryRepository#most_recent_versions}
	MostRecentVersions *ArtifactRegistryRepositoryCleanupPoliciesMostRecentVersions `field:"optional" json:"mostRecentVersions" yaml:"mostRecentVersions"`
}

