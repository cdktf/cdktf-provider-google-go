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
	// Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#billing_account AssuredWorkloadsWorkload#billing_account}
	BillingAccount *string `field:"required" json:"billingAccount" yaml:"billingAccount"`
	// Required.
	//
	// Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#compliance_regime AssuredWorkloadsWorkload#compliance_regime}
	ComplianceRegime *string `field:"required" json:"complianceRegime" yaml:"complianceRegime"`
	// Required.
	//
	// The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#display_name AssuredWorkloadsWorkload#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#location AssuredWorkloadsWorkload#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The organization for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#organization AssuredWorkloadsWorkload#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#id AssuredWorkloadsWorkload#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kms_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#kms_settings AssuredWorkloadsWorkload#kms_settings}
	KmsSettings *AssuredWorkloadsWorkloadKmsSettings `field:"optional" json:"kmsSettings" yaml:"kmsSettings"`
	// Optional. Labels applied to the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#labels AssuredWorkloadsWorkload#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Input only.
	//
	// The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#provisioned_resources_parent AssuredWorkloadsWorkload#provisioned_resources_parent}
	ProvisionedResourcesParent *string `field:"optional" json:"provisionedResourcesParent" yaml:"provisionedResourcesParent"`
	// resource_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#resource_settings AssuredWorkloadsWorkload#resource_settings}
	ResourceSettings interface{} `field:"optional" json:"resourceSettings" yaml:"resourceSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/assured_workloads_workload#timeouts AssuredWorkloadsWorkload#timeouts}
	Timeouts *AssuredWorkloadsWorkloadTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

