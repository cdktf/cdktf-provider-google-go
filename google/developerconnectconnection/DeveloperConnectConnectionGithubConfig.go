// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection


type DeveloperConnectConnectionGithubConfig struct {
	// Required. Immutable. The GitHub Application that was installed to the GitHub user or organization. Possible values: GIT_HUB_APP_UNSPECIFIED DEVELOPER_CONNECT FIREBASE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/developer_connect_connection#github_app DeveloperConnectConnection#github_app}
	GithubApp *string `field:"required" json:"githubApp" yaml:"githubApp"`
	// Optional. GitHub App installation id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/developer_connect_connection#app_installation_id DeveloperConnectConnection#app_installation_id}
	AppInstallationId *string `field:"optional" json:"appInstallationId" yaml:"appInstallationId"`
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/developer_connect_connection#authorizer_credential DeveloperConnectConnection#authorizer_credential}
	AuthorizerCredential *DeveloperConnectConnectionGithubConfigAuthorizerCredential `field:"optional" json:"authorizerCredential" yaml:"authorizerCredential"`
}

