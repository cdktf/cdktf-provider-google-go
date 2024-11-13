// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue


type CloudTasksQueueHttpTarget struct {
	// header_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/cloud_tasks_queue#header_overrides CloudTasksQueue#header_overrides}
	HeaderOverrides interface{} `field:"optional" json:"headerOverrides" yaml:"headerOverrides"`
	// The HTTP method to use for the request.
	//
	// When specified, it overrides HttpRequest for the task.
	// Note that if the value is set to GET the body of the task will be ignored at execution time. Possible values: ["HTTP_METHOD_UNSPECIFIED", "POST", "GET", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/cloud_tasks_queue#http_method CloudTasksQueue#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// oauth_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/cloud_tasks_queue#oauth_token CloudTasksQueue#oauth_token}
	OauthToken *CloudTasksQueueHttpTargetOauthToken `field:"optional" json:"oauthToken" yaml:"oauthToken"`
	// oidc_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/cloud_tasks_queue#oidc_token CloudTasksQueue#oidc_token}
	OidcToken *CloudTasksQueueHttpTargetOidcToken `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// uri_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/cloud_tasks_queue#uri_override CloudTasksQueue#uri_override}
	UriOverride *CloudTasksQueueHttpTargetUriOverride `field:"optional" json:"uriOverride" yaml:"uriOverride"`
}

