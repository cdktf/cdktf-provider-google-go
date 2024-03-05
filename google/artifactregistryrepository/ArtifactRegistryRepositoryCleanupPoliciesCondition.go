// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryCleanupPoliciesCondition struct {
	// Match versions newer than a duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/artifact_registry_repository#newer_than ArtifactRegistryRepository#newer_than}
	NewerThan *string `field:"optional" json:"newerThan" yaml:"newerThan"`
	// Match versions older than a duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/artifact_registry_repository#older_than ArtifactRegistryRepository#older_than}
	OlderThan *string `field:"optional" json:"olderThan" yaml:"olderThan"`
	// Match versions by package prefix. Applied on any prefix match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/artifact_registry_repository#package_name_prefixes ArtifactRegistryRepository#package_name_prefixes}
	PackageNamePrefixes *[]*string `field:"optional" json:"packageNamePrefixes" yaml:"packageNamePrefixes"`
	// Match versions by tag prefix. Applied on any prefix match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/artifact_registry_repository#tag_prefixes ArtifactRegistryRepository#tag_prefixes}
	TagPrefixes *[]*string `field:"optional" json:"tagPrefixes" yaml:"tagPrefixes"`
	// Match versions by tag status. Default value: "ANY" Possible values: ["TAGGED", "UNTAGGED", "ANY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/artifact_registry_repository#tag_state ArtifactRegistryRepository#tag_state}
	TagState *string `field:"optional" json:"tagState" yaml:"tagState"`
	// Match versions by version name prefix. Applied on any prefix match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/artifact_registry_repository#version_name_prefixes ArtifactRegistryRepository#version_name_prefixes}
	VersionNamePrefixes *[]*string `field:"optional" json:"versionNamePrefixes" yaml:"versionNamePrefixes"`
}

