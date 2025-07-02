// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackupbackupplan


type GkeBackupBackupPlanBackupConfigEncryptionKey struct {
	// Google Cloud KMS encryption key. Format: projects/* /locations/* /keyRings/* /cryptoKeys/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/gke_backup_backup_plan#gcp_kms_encryption_key GkeBackupBackupPlan#gcp_kms_encryption_key}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	GcpKmsEncryptionKey *string `field:"required" json:"gcpKmsEncryptionKey" yaml:"gcpKmsEncryptionKey"`
}

