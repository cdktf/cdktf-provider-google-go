package alloydbcluster


type AlloydbClusterContinuousBackupConfig struct {
	// Whether continuous backup recovery is enabled. If not set, defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/alloydb_cluster#enabled AlloydbCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/alloydb_cluster#encryption_config AlloydbCluster#encryption_config}
	EncryptionConfig *AlloydbClusterContinuousBackupConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The numbers of days that are eligible to restore from using PITR.
	//
	// To support the entire recovery window, backups and logs are retained for one day more than the recovery window.
	//
	// If not set, defaults to 14 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/alloydb_cluster#recovery_window_days AlloydbCluster#recovery_window_days}
	RecoveryWindowDays *float64 `field:"optional" json:"recoveryWindowDays" yaml:"recoveryWindowDays"`
}

