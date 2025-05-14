// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupvault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupDrBackupVaultConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Required.
	//
	// The default and minimum enforced retention for each backup within the backup vault. The enforced retention for each backup can be extended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#backup_minimum_enforced_retention_duration BackupDrBackupVault#backup_minimum_enforced_retention_duration}
	BackupMinimumEnforcedRetentionDuration *string `field:"required" json:"backupMinimumEnforcedRetentionDuration" yaml:"backupMinimumEnforcedRetentionDuration"`
	// Required. ID of the requesting object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#backup_vault_id BackupDrBackupVault#backup_vault_id}
	BackupVaultId *string `field:"required" json:"backupVaultId" yaml:"backupVaultId"`
	// The GCP location for the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#location BackupDrBackupVault#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Access restriction for the backup vault.
	//
	// Default value is 'WITHIN_ORGANIZATION' if not provided during creation. Default value: "WITHIN_ORGANIZATION" Possible values: ["ACCESS_RESTRICTION_UNSPECIFIED", "WITHIN_PROJECT", "WITHIN_ORGANIZATION", "UNRESTRICTED", "WITHIN_ORG_BUT_UNRESTRICTED_FOR_BA"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#access_restriction BackupDrBackupVault#access_restriction}
	AccessRestriction *string `field:"optional" json:"accessRestriction" yaml:"accessRestriction"`
	// Allow idempotent deletion of backup vault. The request will still succeed in case the backup vault does not exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#allow_missing BackupDrBackupVault#allow_missing}
	AllowMissing interface{} `field:"optional" json:"allowMissing" yaml:"allowMissing"`
	// Optional. User annotations. See https://google.aip.dev/128#annotations Stores small amounts of arbitrary data.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#annotations BackupDrBackupVault#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Optional. The description of the BackupVault instance (2048 characters or less).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#description BackupDrBackupVault#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional. Time after which the BackupVault resource is locked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#effective_time BackupDrBackupVault#effective_time}
	EffectiveTime *string `field:"optional" json:"effectiveTime" yaml:"effectiveTime"`
	// If set, the following restrictions against deletion of the backup vault instance can be overridden:    * deletion of a backup vault instance containing no backups, but still containing empty datasources.
	//
	// * deletion of a backup vault instance that is being referenced by an active backup plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#force_delete BackupDrBackupVault#force_delete}
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// If set, allow update to extend the minimum enforced retention for backup vault.
	//
	// This overrides
	//  the restriction against conflicting retention periods. This conflict may occur when the
	//  expiration schedule defined by the associated backup plan is shorter than the minimum
	//  retention set by the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#force_update BackupDrBackupVault#force_update}
	ForceUpdate interface{} `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#id BackupDrBackupVault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If set, the following restrictions against deletion of the backup vault instance can be overridden:    * deletion of a backup vault instance that is being referenced by an active backup plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#ignore_backup_plan_references BackupDrBackupVault#ignore_backup_plan_references}
	IgnoreBackupPlanReferences interface{} `field:"optional" json:"ignoreBackupPlanReferences" yaml:"ignoreBackupPlanReferences"`
	// If set, the following restrictions against deletion of the backup vault instance can be overridden:    * deletion of a backup vault instance containing no backups, but still containing empty datasources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#ignore_inactive_datasources BackupDrBackupVault#ignore_inactive_datasources}
	IgnoreInactiveDatasources interface{} `field:"optional" json:"ignoreInactiveDatasources" yaml:"ignoreInactiveDatasources"`
	// Optional. Resource labels to represent user provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#labels BackupDrBackupVault#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#project BackupDrBackupVault#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/backup_dr_backup_vault#timeouts BackupDrBackupVault#timeouts}
	Timeouts *BackupDrBackupVaultTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

