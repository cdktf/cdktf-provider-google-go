// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerdatabase


type SpannerDatabaseEncryptionConfig struct {
	// Fully qualified name of the KMS key to use to encrypt this database.
	//
	// This key must exist
	// in the same location as the Spanner Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/spanner_database#kms_key_name SpannerDatabase#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Fully qualified name of the KMS keys to use to encrypt this database.
	//
	// The keys must exist
	// in the same locations as the Spanner Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/spanner_database#kms_key_names SpannerDatabase#kms_key_names}
	KmsKeyNames *[]*string `field:"optional" json:"kmsKeyNames" yaml:"kmsKeyNames"`
}

