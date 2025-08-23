// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package assuredworkloadsworkload

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AssuredWorkloadsWorkloadConfig struct {
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
	// Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, HITRUST, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS, ISR_REGIONS, ISR_REGIONS_AND_SUPPORT, CA_PROTECTED_B, IL5, IL2, JP_REGIONS_AND_SUPPORT, KSA_REGIONS_AND_SUPPORT_WITH_SOVEREIGNTY_CONTROLS, REGIONAL_CONTROLS, HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS, HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS_US_SUPPORT, IRS_1075
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#compliance_regime AssuredWorkloadsWorkload#compliance_regime}
	ComplianceRegime *string `field:"required" json:"complianceRegime" yaml:"complianceRegime"`
	// Required.
	//
	// The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#display_name AssuredWorkloadsWorkload#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#location AssuredWorkloadsWorkload#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The organization for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#organization AssuredWorkloadsWorkload#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Optional.
	//
	// Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#billing_account AssuredWorkloadsWorkload#billing_account}
	BillingAccount *string `field:"optional" json:"billingAccount" yaml:"billingAccount"`
	// Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#enable_sovereign_controls AssuredWorkloadsWorkload#enable_sovereign_controls}
	EnableSovereignControls interface{} `field:"optional" json:"enableSovereignControls" yaml:"enableSovereignControls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#id AssuredWorkloadsWorkload#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kms_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#kms_settings AssuredWorkloadsWorkload#kms_settings}
	KmsSettings *AssuredWorkloadsWorkloadKmsSettings `field:"optional" json:"kmsSettings" yaml:"kmsSettings"`
	// Optional. Labels applied to the workload.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effective_labels` for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#labels AssuredWorkloadsWorkload#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Optional. Partner regime associated with this workload. Possible values: PARTNER_UNSPECIFIED, LOCAL_CONTROLS_BY_S3NS, SOVEREIGN_CONTROLS_BY_T_SYSTEMS, SOVEREIGN_CONTROLS_BY_SIA_MINSAIT, SOVEREIGN_CONTROLS_BY_PSN, SOVEREIGN_CONTROLS_BY_CNTXT, SOVEREIGN_CONTROLS_BY_CNTXT_NO_EKM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#partner AssuredWorkloadsWorkload#partner}
	Partner *string `field:"optional" json:"partner" yaml:"partner"`
	// partner_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#partner_permissions AssuredWorkloadsWorkload#partner_permissions}
	PartnerPermissions *AssuredWorkloadsWorkloadPartnerPermissions `field:"optional" json:"partnerPermissions" yaml:"partnerPermissions"`
	// Optional.
	//
	// Input only. Billing account necessary for purchasing services from Sovereign Partners. This field is required for creating SIA/PSN/CNTXT partner workloads. The caller should have 'billing.resourceAssociations.create' IAM permission on this billing-account. The format of this string is billingAccounts/AAAAAA-BBBBBB-CCCCCC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#partner_services_billing_account AssuredWorkloadsWorkload#partner_services_billing_account}
	PartnerServicesBillingAccount *string `field:"optional" json:"partnerServicesBillingAccount" yaml:"partnerServicesBillingAccount"`
	// Input only.
	//
	// The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#provisioned_resources_parent AssuredWorkloadsWorkload#provisioned_resources_parent}
	ProvisionedResourcesParent *string `field:"optional" json:"provisionedResourcesParent" yaml:"provisionedResourcesParent"`
	// resource_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#resource_settings AssuredWorkloadsWorkload#resource_settings}
	ResourceSettings interface{} `field:"optional" json:"resourceSettings" yaml:"resourceSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#timeouts AssuredWorkloadsWorkload#timeouts}
	Timeouts *AssuredWorkloadsWorkloadTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional.
	//
	// Indicates whether the e-mail notification for a violation is enabled for a workload. This value will be by default True, and if not present will be considered as true. This should only be updated via updateWorkload call. Any Changes to this field during the createWorkload call will not be honored. This will always be true while creating the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#violation_notifications_enabled AssuredWorkloadsWorkload#violation_notifications_enabled}
	ViolationNotificationsEnabled interface{} `field:"optional" json:"violationNotificationsEnabled" yaml:"violationNotificationsEnabled"`
	// workload_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/assured_workloads_workload#workload_options AssuredWorkloadsWorkload#workload_options}
	WorkloadOptions *AssuredWorkloadsWorkloadWorkloadOptions `field:"optional" json:"workloadOptions" yaml:"workloadOptions"`
}

