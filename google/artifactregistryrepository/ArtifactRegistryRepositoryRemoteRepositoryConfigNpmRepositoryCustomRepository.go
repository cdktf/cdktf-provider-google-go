// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryCustomRepository struct {
	// Specific uri to the registry, e.g. '"https://registry.npmjs.org"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/artifact_registry_repository#uri ArtifactRegistryRepository#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

