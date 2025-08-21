// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectinsightsconfig


type DeveloperConnectInsightsConfigArtifactConfigs struct {
	// google_artifact_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/developer_connect_insights_config#google_artifact_analysis DeveloperConnectInsightsConfig#google_artifact_analysis}
	GoogleArtifactAnalysis *DeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactAnalysis `field:"optional" json:"googleArtifactAnalysis" yaml:"googleArtifactAnalysis"`
	// google_artifact_registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/developer_connect_insights_config#google_artifact_registry DeveloperConnectInsightsConfig#google_artifact_registry}
	GoogleArtifactRegistry *DeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry `field:"optional" json:"googleArtifactRegistry" yaml:"googleArtifactRegistry"`
	// The URI of the artifact that is deployed.
	//
	// e.g. 'us-docker.pkg.dev/my-project/my-repo/image'.
	// The URI does not include the tag / digest because it captures a lineage of
	// artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/developer_connect_insights_config#uri DeveloperConnectInsightsConfig#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

