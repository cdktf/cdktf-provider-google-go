// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudquotasquotapreference

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudQuotasQuotaPreferenceConfig struct {
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
	// quota_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#quota_config CloudQuotasQuotaPreference#quota_config}
	QuotaConfig *CloudQuotasQuotaPreferenceQuotaConfig `field:"required" json:"quotaConfig" yaml:"quotaConfig"`
	// An email address that can be used for quota related communication between the Google Cloud and the user in case the Google Cloud needs further information to make a decision on whether the user preferred quota can be granted.
	//
	// The Google account for the email address must have quota update permission for the project, folder or organization this quota preference is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#contact_email CloudQuotasQuotaPreference#contact_email}
	ContactEmail *string `field:"optional" json:"contactEmail" yaml:"contactEmail"`
	// The dimensions that this quota preference applies to.
	//
	// The key of the map entry is the name of a dimension, such as "region", "zone", "network_id", and the value of the map entry is the dimension value. If a dimension is missing from the map of dimensions, the quota preference applies to all the dimension values except for those that have other quota preferences configured for the specific value.
	//
	// NOTE: QuotaPreferences can only be applied across all values of "user" and "resource" dimension. Do not set values for "user" or "resource" in the dimension map.
	//
	// Example: '{"provider": "Foo Inc"}' where "provider" is a service specific dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#dimensions CloudQuotasQuotaPreference#dimensions}
	Dimensions *map[string]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#id CloudQuotasQuotaPreference#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The list of quota safety checks to be ignored. Default value: "QUOTA_SAFETY_CHECK_UNSPECIFIED" Possible values: ["QUOTA_SAFETY_CHECK_UNSPECIFIED", "QUOTA_DECREASE_BELOW_USAGE", "QUOTA_DECREASE_PERCENTAGE_TOO_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#ignore_safety_checks CloudQuotasQuotaPreference#ignore_safety_checks}
	IgnoreSafetyChecks *string `field:"optional" json:"ignoreSafetyChecks" yaml:"ignoreSafetyChecks"`
	// The reason / justification for this quota preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#justification CloudQuotasQuotaPreference#justification}
	Justification *string `field:"optional" json:"justification" yaml:"justification"`
	// The resource name of the quota preference. Required except in the CREATE requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#name CloudQuotasQuotaPreference#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parent of the quota preference.
	//
	// Allowed parents are "projects/[project-id / number]" or "folders/[folder-id / number]" or "organizations/[org-id / number]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#parent CloudQuotasQuotaPreference#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// The id of the quota to which the quota preference is applied.
	//
	// A quota id is unique in the service.
	// Example: 'CPUS-per-project-region'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#quota_id CloudQuotasQuotaPreference#quota_id}
	QuotaId *string `field:"optional" json:"quotaId" yaml:"quotaId"`
	// The name of the service to which the quota preference is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#service CloudQuotasQuotaPreference#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/cloud_quotas_quota_preference#timeouts CloudQuotasQuotaPreference#timeouts}
	Timeouts *CloudQuotasQuotaPreferenceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

