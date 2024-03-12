// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepository struct {
	// A common public repository base for Apt, e.g. '"debian/dists/buster"' Possible values: ["DEBIAN", "UBUNTU"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/artifact_registry_repository#repository_base ArtifactRegistryRepository#repository_base}
	RepositoryBase *string `field:"required" json:"repositoryBase" yaml:"repositoryBase"`
	// Specific repository from the base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/artifact_registry_repository#repository_path ArtifactRegistryRepository#repository_path}
	RepositoryPath *string `field:"required" json:"repositoryPath" yaml:"repositoryPath"`
}

