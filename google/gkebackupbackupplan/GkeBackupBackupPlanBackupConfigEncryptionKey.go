package gkebackupbackupplan


type GkeBackupBackupPlanBackupConfigEncryptionKey struct {
	// Google Cloud KMS encryption key. Format: projects/*\/locations/*\/keyRings/*\/cryptoKeys/*.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan#gcp_kms_encryption_key GkeBackupBackupPlan#gcp_kms_encryption_key}
	GcpKmsEncryptionKey *string `field:"required" json:"gcpKmsEncryptionKey" yaml:"gcpKmsEncryptionKey"`
}

