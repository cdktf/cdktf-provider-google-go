// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrserviceconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupDrServiceConfigConfig struct {
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
	// The location in which the Service config is to be initialized.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/backup_dr_service_config#location BackupDrServiceConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource type to which the default service config will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/backup_dr_service_config#resource_type BackupDrServiceConfig#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/backup_dr_service_config#id BackupDrServiceConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/backup_dr_service_config#project BackupDrServiceConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/backup_dr_service_config#timeouts BackupDrServiceConfig#timeouts}
	Timeouts *BackupDrServiceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

