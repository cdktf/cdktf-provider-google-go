// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatatransferconfig


type BigqueryDataTransferConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/bigquery_data_transfer_config#create BigqueryDataTransferConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/bigquery_data_transfer_config#delete BigqueryDataTransferConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/bigquery_data_transfer_config#update BigqueryDataTransferConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

