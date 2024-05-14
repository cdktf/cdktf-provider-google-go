// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository struct {
	// public_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *ArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepository `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

