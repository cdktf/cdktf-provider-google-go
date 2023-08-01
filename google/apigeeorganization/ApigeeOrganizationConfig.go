package apigeeorganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeOrganizationConfig struct {
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
	// The project ID associated with the Apigee organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#project_id ApigeeOrganization#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#analytics_region ApigeeOrganization#analytics_region}
	AnalyticsRegion *string `field:"optional" json:"analyticsRegion" yaml:"analyticsRegion"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	//
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when 'RuntimeType' is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#authorized_network ApigeeOrganization#authorized_network}
	AuthorizedNetwork *string `field:"optional" json:"authorizedNetwork" yaml:"authorizedNetwork"`
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#billing_type ApigeeOrganization#billing_type}
	BillingType *string `field:"optional" json:"billingType" yaml:"billingType"`
	// Description of the Apigee organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#description ApigeeOrganization#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Flag that specifies whether the VPC Peering through Private Google Access should be disabled between the consumer network and Apigee.
	//
	// Required if an 'authorizedNetwork'
	// on the consumer project is not provided, in which case the flag should be set to 'true'.
	// Valid only when 'RuntimeType' is set to CLOUD. The value must be set before the creation
	// of any Apigee runtime instance and can be updated only when there are no runtime instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#disable_vpc_peering ApigeeOrganization#disable_vpc_peering}
	DisableVpcPeering interface{} `field:"optional" json:"disableVpcPeering" yaml:"disableVpcPeering"`
	// The display name of the Apigee organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#display_name ApigeeOrganization#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#id ApigeeOrganization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#properties ApigeeOrganization#properties}
	Properties *ApigeeOrganizationProperties `field:"optional" json:"properties" yaml:"properties"`
	// Optional.
	//
	// This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored. Default value: "DELETION_RETENTION_UNSPECIFIED" Possible values: ["DELETION_RETENTION_UNSPECIFIED", "MINIMUM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#retention ApigeeOrganization#retention}
	Retention *string `field:"optional" json:"retention" yaml:"retention"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	//
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when 'RuntimeType' is CLOUD. For example: 'projects/foo/locations/us/keyRings/bar/cryptoKeys/baz'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#runtime_database_encryption_key_name ApigeeOrganization#runtime_database_encryption_key_name}
	RuntimeDatabaseEncryptionKeyName *string `field:"optional" json:"runtimeDatabaseEncryptionKeyName" yaml:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased. Default value: "CLOUD" Possible values: ["CLOUD", "HYBRID"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#runtime_type ApigeeOrganization#runtime_type}
	RuntimeType *string `field:"optional" json:"runtimeType" yaml:"runtimeType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/apigee_organization#timeouts ApigeeOrganization#timeouts}
	Timeouts *ApigeeOrganizationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

