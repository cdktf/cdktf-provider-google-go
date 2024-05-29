// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappbackuppolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappBackupPolicyConfig struct {
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
	// Number of daily backups to keep. Note that the minimum daily backup limit is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#daily_backup_limit NetappBackupPolicy#daily_backup_limit}
	DailyBackupLimit *float64 `field:"required" json:"dailyBackupLimit" yaml:"dailyBackupLimit"`
	// Name of the region for the policy to apply to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#location NetappBackupPolicy#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Number of monthly backups to keep.
	//
	// Note that the sum of daily, weekly and monthly backups should be greater than 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#monthly_backup_limit NetappBackupPolicy#monthly_backup_limit}
	MonthlyBackupLimit *float64 `field:"required" json:"monthlyBackupLimit" yaml:"monthlyBackupLimit"`
	// The name of the backup policy. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#name NetappBackupPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Number of weekly backups to keep.
	//
	// Note that the sum of daily, weekly and monthly backups should be greater than 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#weekly_backup_limit NetappBackupPolicy#weekly_backup_limit}
	WeeklyBackupLimit *float64 `field:"required" json:"weeklyBackupLimit" yaml:"weeklyBackupLimit"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#description NetappBackupPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If enabled, make backups automatically according to the schedules.
	//
	// This will be applied to all volumes that have this policy attached and enforced on volume level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#enabled NetappBackupPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#id NetappBackupPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#labels NetappBackupPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#project NetappBackupPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/netapp_backup_policy#timeouts NetappBackupPolicy#timeouts}
	Timeouts *NetappBackupPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

