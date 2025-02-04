// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionCloudSqlCredential struct {
	// Password for database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/bigquery_connection#password BigqueryConnection#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username for database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/bigquery_connection#username BigqueryConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

