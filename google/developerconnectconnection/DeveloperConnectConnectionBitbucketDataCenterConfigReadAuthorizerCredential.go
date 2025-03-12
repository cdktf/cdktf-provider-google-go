// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection


type DeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential struct {
	// Required. A SecretManager resource containing the user token that authorizes the Developer Connect connection. Format: 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/developer_connect_connection#user_token_secret_version DeveloperConnectConnection#user_token_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	UserTokenSecretVersion *string `field:"required" json:"userTokenSecretVersion" yaml:"userTokenSecretVersion"`
}

