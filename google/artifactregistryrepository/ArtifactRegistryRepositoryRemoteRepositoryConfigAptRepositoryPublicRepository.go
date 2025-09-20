// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepository struct {
	// A common public repository base for Apt, e.g. '"debian/dists/stable"' Possible values: ["DEBIAN", "UBUNTU", "DEBIAN_SNAPSHOT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/artifact_registry_repository#repository_base ArtifactRegistryRepository#repository_base}
	RepositoryBase *string `field:"required" json:"repositoryBase" yaml:"repositoryBase"`
	// Specific repository from the base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/artifact_registry_repository#repository_path ArtifactRegistryRepository#repository_path}
	RepositoryPath *string `field:"required" json:"repositoryPath" yaml:"repositoryPath"`
}

