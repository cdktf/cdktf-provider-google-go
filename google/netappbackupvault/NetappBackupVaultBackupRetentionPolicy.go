// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappbackupvault


type NetappBackupVaultBackupRetentionPolicy struct {
	// Minimum retention duration in days for backups in the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/netapp_backup_vault#backup_minimum_enforced_retention_days NetappBackupVault#backup_minimum_enforced_retention_days}
	BackupMinimumEnforcedRetentionDays *float64 `field:"required" json:"backupMinimumEnforcedRetentionDays" yaml:"backupMinimumEnforcedRetentionDays"`
	// Indicates if the daily backups are immutable. At least one of daily_backup_immutable, weekly_backup_immutable, monthly_backup_immutable and manual_backup_immutable must be true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/netapp_backup_vault#daily_backup_immutable NetappBackupVault#daily_backup_immutable}
	DailyBackupImmutable interface{} `field:"optional" json:"dailyBackupImmutable" yaml:"dailyBackupImmutable"`
	// Indicates if the manual backups are immutable. At least one of daily_backup_immutable, weekly_backup_immutable, monthly_backup_immutable and manual_backup_immutable must be true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/netapp_backup_vault#manual_backup_immutable NetappBackupVault#manual_backup_immutable}
	ManualBackupImmutable interface{} `field:"optional" json:"manualBackupImmutable" yaml:"manualBackupImmutable"`
	// Indicates if the monthly backups are immutable. At least one of daily_backup_immutable, weekly_backup_immutable, monthly_backup_immutable and manual_backup_immutable must be true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/netapp_backup_vault#monthly_backup_immutable NetappBackupVault#monthly_backup_immutable}
	MonthlyBackupImmutable interface{} `field:"optional" json:"monthlyBackupImmutable" yaml:"monthlyBackupImmutable"`
	// Indicates if the weekly backups are immutable. At least one of daily_backup_immutable, weekly_backup_immutable, monthly_backup_immutable and manual_backup_immutable must be true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/netapp_backup_vault#weekly_backup_immutable NetappBackupVault#weekly_backup_immutable}
	WeeklyBackupImmutable interface{} `field:"optional" json:"weeklyBackupImmutable" yaml:"weeklyBackupImmutable"`
}

