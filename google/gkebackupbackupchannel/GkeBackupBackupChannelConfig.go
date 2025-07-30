// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackupbackupchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeBackupBackupChannelConfig struct {
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
	// The project where Backups are allowed to be stored. The format is 'projects/{project}'. {project} can be project number or project id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_backup_backup_channel#destination_project GkeBackupBackupChannel#destination_project}
	DestinationProject *string `field:"required" json:"destinationProject" yaml:"destinationProject"`
	// The region of the Backup Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_backup_backup_channel#location GkeBackupBackupChannel#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The full name of the BackupChannel Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_backup_backup_channel#name GkeBackupBackupChannel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// User specified descriptive string for this BackupChannel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_backup_backup_channel#description GkeBackupBackupChannel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_backup_backup_channel#id GkeBackupBackupChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Description: A set of custom labels supplied by the user.
	//
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_backup_backup_channel#labels GkeBackupBackupChannel#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_backup_backup_channel#project GkeBackupBackupChannel#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gke_backup_backup_channel#timeouts GkeBackupBackupChannel#timeouts}
	Timeouts *GkeBackupBackupChannelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

