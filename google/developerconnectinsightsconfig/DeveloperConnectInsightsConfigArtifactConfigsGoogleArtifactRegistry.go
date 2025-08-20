// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectinsightsconfig


type DeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry struct {
	// The name of the artifact registry package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_insights_config#artifact_registry_package DeveloperConnectInsightsConfig#artifact_registry_package}
	ArtifactRegistryPackage *string `field:"required" json:"artifactRegistryPackage" yaml:"artifactRegistryPackage"`
	// The host project of Artifact Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_insights_config#project_id DeveloperConnectInsightsConfig#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

