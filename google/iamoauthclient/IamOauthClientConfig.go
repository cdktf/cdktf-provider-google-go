// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamoauthclient

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamOauthClientConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Required. The list of OAuth grant types is allowed for the OauthClient.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#allowed_grant_types IamOauthClient#allowed_grant_types}
	AllowedGrantTypes *[]*string `field:"required" json:"allowedGrantTypes" yaml:"allowedGrantTypes"`
	// Required. The list of redirect uris that is allowed to redirect back when authorization process is completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#allowed_redirect_uris IamOauthClient#allowed_redirect_uris}
	AllowedRedirectUris *[]*string `field:"required" json:"allowedRedirectUris" yaml:"allowedRedirectUris"`
	// Required. The list of scopes that the OauthClient is allowed to request during OAuth flows.
	//
	// The following scopes are supported:
	//
	// * 'https://www.googleapis.com/auth/cloud-platform': See, edit, configure,
	// and delete your Google Cloud data and see the email address for your Google
	// Account.
	// * 'openid': The OAuth client can associate you with your personal
	// information on Google Cloud.
	// * 'email': The OAuth client can read a federated identity's email address.
	// * 'groups': The OAuth client can read a federated identity's groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#allowed_scopes IamOauthClient#allowed_scopes}
	AllowedScopes *[]*string `field:"required" json:"allowedScopes" yaml:"allowedScopes"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#location IamOauthClient#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required.
	//
	// The ID to use for the OauthClient, which becomes the final component of
	// the resource name. This value should be a string of 6 to 63 lowercase
	// letters, digits, or hyphens. It must start with a letter, and cannot have a
	// trailing hyphen. The prefix 'gcp-' is reserved for use by Google, and may
	// not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#oauth_client_id IamOauthClient#oauth_client_id}
	OauthClientId *string `field:"required" json:"oauthClientId" yaml:"oauthClientId"`
	// Immutable.
	//
	// The type of OauthClient. Either public or private.
	// For private clients, the client secret can be managed using the dedicated
	// OauthClientCredential resource.
	// Possible values:
	// CLIENT_TYPE_UNSPECIFIED
	// PUBLIC_CLIENT
	// CONFIDENTIAL_CLIENT
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#client_type IamOauthClient#client_type}
	ClientType *string `field:"optional" json:"clientType" yaml:"clientType"`
	// A user-specified description of the OauthClient.
	//
	// Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#description IamOauthClient#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the OauthClient is disabled. You cannot use a disabled OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#disabled IamOauthClient#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A user-specified display name of the OauthClient.
	//
	// Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#display_name IamOauthClient#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#id IamOauthClient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#project IamOauthClient#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/iam_oauth_client#timeouts IamOauthClient#timeouts}
	Timeouts *IamOauthClientTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

