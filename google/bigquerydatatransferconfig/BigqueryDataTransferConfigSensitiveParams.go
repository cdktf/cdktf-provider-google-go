// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatatransferconfig


type BigqueryDataTransferConfigSensitiveParams struct {
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/bigquery_data_transfer_config#secret_access_key BigqueryDataTransferConfig#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

