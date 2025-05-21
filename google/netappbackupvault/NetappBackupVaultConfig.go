// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappbackupvault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappBackupVaultConfig struct {
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
	// Location (region) of the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#location NetappBackupVault#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the backup vault. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#name NetappBackupVault#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Region in which backup is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#backup_region NetappBackupVault#backup_region}
	BackupRegion *string `field:"optional" json:"backupRegion" yaml:"backupRegion"`
	// Type of the backup vault to be created. Default is IN_REGION. Possible values: ["BACKUP_VAULT_TYPE_UNSPECIFIED", "IN_REGION", "CROSS_REGION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#backup_vault_type NetappBackupVault#backup_vault_type}
	BackupVaultType *string `field:"optional" json:"backupVaultType" yaml:"backupVaultType"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#description NetappBackupVault#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#id NetappBackupVault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#labels NetappBackupVault#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#project NetappBackupVault#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/netapp_backup_vault#timeouts NetappBackupVault#timeouts}
	Timeouts *NetappBackupVaultTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

