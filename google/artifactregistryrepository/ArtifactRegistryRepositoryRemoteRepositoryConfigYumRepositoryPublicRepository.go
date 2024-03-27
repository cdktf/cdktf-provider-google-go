// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepository struct {
	// A common public repository base for Yum. Possible values: ["CENTOS", "CENTOS_DEBUG", "CENTOS_VAULT", "CENTOS_STREAM", "ROCKY", "EPEL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/artifact_registry_repository#repository_base ArtifactRegistryRepository#repository_base}
	RepositoryBase *string `field:"required" json:"repositoryBase" yaml:"repositoryBase"`
	// Specific repository from the base, e.g. '"centos/8-stream/BaseOS/x86_64/os"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/artifact_registry_repository#repository_path ArtifactRegistryRepository#repository_path}
	RepositoryPath *string `field:"required" json:"repositoryPath" yaml:"repositoryPath"`
}

