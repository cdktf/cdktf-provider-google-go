// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryjob


type BigqueryJobQueryDestinationEncryptionConfiguration struct {
	// Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.
	//
	// The BigQuery Service Account associated with your project requires access to this encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/bigquery_job#kms_key_name BigqueryJob#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

