// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatatransferconfig


type BigqueryDataTransferConfigSensitiveParams struct {
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/bigquery_data_transfer_config#secret_access_key BigqueryDataTransferConfig#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/bigquery_data_transfer_config#secret_access_key_wo BigqueryDataTransferConfig#secret_access_key_wo}
	SecretAccessKeyWo *string `field:"optional" json:"secretAccessKeyWo" yaml:"secretAccessKeyWo"`
	// The version of the sensitive params - used to trigger updates of the write-only params.
	//
	// For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/bigquery_data_transfer_config#secret_access_key_wo_version BigqueryDataTransferConfig#secret_access_key_wo_version}
	SecretAccessKeyWoVersion *float64 `field:"optional" json:"secretAccessKeyWoVersion" yaml:"secretAccessKeyWoVersion"`
}

