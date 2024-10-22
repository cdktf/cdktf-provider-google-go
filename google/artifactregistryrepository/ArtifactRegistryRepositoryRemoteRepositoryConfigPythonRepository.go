// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository struct {
	// custom_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/artifact_registry_repository#custom_repository ArtifactRegistryRepository#custom_repository}
	CustomRepository *ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepository `field:"optional" json:"customRepository" yaml:"customRepository"`
	// Address of the remote repository. Default value: "PYPI" Possible values: ["PYPI"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

