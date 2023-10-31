// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxagent


type DialogflowCxAgentGitIntegrationSettings struct {
	// github_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/dialogflow_cx_agent#github_settings DialogflowCxAgent#github_settings}
	GithubSettings *DialogflowCxAgentGitIntegrationSettingsGithubSettings `field:"optional" json:"githubSettings" yaml:"githubSettings"`
}

