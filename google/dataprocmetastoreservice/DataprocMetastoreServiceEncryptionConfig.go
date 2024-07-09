// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceEncryptionConfig struct {
	// The fully qualified customer provided Cloud KMS key name to use for customer data encryption. Use the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/dataproc_metastore_service#kms_key DataprocMetastoreService#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

