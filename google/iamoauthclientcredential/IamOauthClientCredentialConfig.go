// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamoauthclientcredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamOauthClientCredentialConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/iam_oauth_client_credential#location IamOauthClientCredential#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/iam_oauth_client_credential#oauthclient IamOauthClientCredential#oauthclient}
	Oauthclient *string `field:"required" json:"oauthclient" yaml:"oauthclient"`
	// Required.
	//
	// The ID to use for the OauthClientCredential, which becomes the
	// final component of the resource name. This value should be 4-32 characters,
	// and may contain the characters [a-z0-9-]. The prefix 'gcp-' is
	// reserved for use by Google, and may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/iam_oauth_client_credential#oauth_client_credential_id IamOauthClientCredential#oauth_client_credential_id}
	OauthClientCredentialId *string `field:"required" json:"oauthClientCredentialId" yaml:"oauthClientCredentialId"`
	// Whether the OauthClientCredential is disabled. You cannot use a disabled OauthClientCredential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/iam_oauth_client_credential#disabled IamOauthClientCredential#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A user-specified display name of the OauthClientCredential.
	//
	// Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/iam_oauth_client_credential#display_name IamOauthClientCredential#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/iam_oauth_client_credential#id IamOauthClientCredential#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/iam_oauth_client_credential#project IamOauthClientCredential#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/iam_oauth_client_credential#timeouts IamOauthClientCredential#timeouts}
	Timeouts *IamOauthClientCredentialTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

