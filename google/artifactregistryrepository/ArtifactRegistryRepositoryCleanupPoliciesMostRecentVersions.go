// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryCleanupPoliciesMostRecentVersions struct {
	// Minimum number of versions to keep.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/artifact_registry_repository#keep_count ArtifactRegistryRepository#keep_count}
	KeepCount *float64 `field:"optional" json:"keepCount" yaml:"keepCount"`
	// Match versions by package prefix. Applied on any prefix match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/artifact_registry_repository#package_name_prefixes ArtifactRegistryRepository#package_name_prefixes}
	PackageNamePrefixes *[]*string `field:"optional" json:"packageNamePrefixes" yaml:"packageNamePrefixes"`
}

