// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumequotarule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeQuotaRuleConfig struct {
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
	// The maximum allowed capacity in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#disk_limit_mib NetappVolumeQuotaRule#disk_limit_mib}
	DiskLimitMib *float64 `field:"required" json:"diskLimitMib" yaml:"diskLimitMib"`
	// The resource name of the quotaRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#name NetappVolumeQuotaRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Types of Quota Rule. Possible values: ["INDIVIDUAL_USER_QUOTA", "INDIVIDUAL_GROUP_QUOTA", "DEFAULT_USER_QUOTA", "DEFAULT_GROUP_QUOTA"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#type NetappVolumeQuotaRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Name of the volume to create the quotaRule in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#volume_name NetappVolumeQuotaRule#volume_name}
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
	// Description for the quota rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#description NetappVolumeQuotaRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#id NetappVolumeQuotaRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs of the quota rule. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#labels NetappVolumeQuotaRule#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Loction of the quotaRule. QuotaRules are child resources of volumes and live in the same location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#location NetappVolumeQuotaRule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#project NetappVolumeQuotaRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The quota rule applies to the specified user or group.
	//
	// Valid targets for volumes with NFS protocol enabled:
	//   - UNIX UID for individual user quota
	//   - UNIX GID for individual group quota
	// Valid targets for volumes with SMB protocol enabled:
	//   - Windows SID for individual user quota
	// Leave empty for default quotas
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#target NetappVolumeQuotaRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/netapp_volume_quota_rule#timeouts NetappVolumeQuotaRule#timeouts}
	Timeouts *NetappVolumeQuotaRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

