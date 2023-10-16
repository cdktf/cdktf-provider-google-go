// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionAwsAccessRole struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.2.0/docs/resources/bigquery_connection#iam_role_id BigqueryConnection#iam_role_id}
	IamRoleId *string `field:"required" json:"iamRoleId" yaml:"iamRoleId"`
}

