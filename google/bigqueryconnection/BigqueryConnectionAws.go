// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionAws struct {
	// access_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.0.0/docs/resources/bigquery_connection#access_role BigqueryConnection#access_role}
	AccessRole *BigqueryConnectionAwsAccessRole `field:"required" json:"accessRole" yaml:"accessRole"`
}

