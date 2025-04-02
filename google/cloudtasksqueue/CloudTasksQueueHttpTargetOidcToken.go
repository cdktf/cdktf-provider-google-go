// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue


type CloudTasksQueueHttpTargetOidcToken struct {
	// Service account email to be used for generating OIDC token.
	//
	// The service account must be within the same project as the queue.
	// The caller must have iam.serviceAccounts.actAs permission for the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/cloud_tasks_queue#service_account_email CloudTasksQueue#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/cloud_tasks_queue#audience CloudTasksQueue#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
}

