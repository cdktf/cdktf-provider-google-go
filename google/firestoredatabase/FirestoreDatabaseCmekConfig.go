// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firestoredatabase


type FirestoreDatabaseCmekConfig struct {
	// The resource ID of a Cloud KMS key.
	//
	// If set, the database created will
	// be a Customer-managed Encryption Key (CMEK) database encrypted with
	// this key. This feature is allowlist only in initial launch.
	//
	// Only keys in the same location as this database are allowed to be used
	// for encryption. For Firestore's nam5 multi-region, this corresponds to Cloud KMS
	// multi-region us. For Firestore's eur3 multi-region, this corresponds to
	// Cloud KMS multi-region europe. See https://cloud.google.com/kms/docs/locations.
	//
	// This value should be the KMS key resource ID in the format of
	// 'projects/{project_id}/locations/{kms_location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}'.
	// How to retrieve this resource ID is listed at
	// https://cloud.google.com/kms/docs/getting-resource-ids#getting_the_id_for_a_key_and_version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/firestore_database#kms_key_name FirestoreDatabase#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

