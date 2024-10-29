// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue


type CloudTasksQueueHttpTargetOauthToken struct {
	// Service account email to be used for generating OAuth token.
	//
	// The service account must be within the same project as the queue.
	// The caller must have iam.serviceAccounts.actAs permission for the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/cloud_tasks_queue#service_account_email CloudTasksQueue#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// OAuth scope to be used for generating OAuth access token. If not specified, "https://www.googleapis.com/auth/cloud-platform" will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/cloud_tasks_queue#scope CloudTasksQueue#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

