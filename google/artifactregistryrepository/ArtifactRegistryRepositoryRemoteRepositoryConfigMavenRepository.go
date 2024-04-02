// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository struct {
	// Address of the remote repository. Default value: "MAVEN_CENTRAL" Possible values: ["MAVEN_CENTRAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.23.0/docs/resources/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

