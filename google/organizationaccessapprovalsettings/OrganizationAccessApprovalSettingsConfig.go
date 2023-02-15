package organizationaccessapprovalsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationAccessApprovalSettingsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// enrolled_services block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_access_approval_settings#enrolled_services OrganizationAccessApprovalSettings#enrolled_services}
	EnrolledServices interface{} `field:"required" json:"enrolledServices" yaml:"enrolledServices"`
	// ID of the organization of the access approval settings.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_access_approval_settings#organization_id OrganizationAccessApprovalSettings#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// The asymmetric crypto key version to use for signing approval requests.
	//
	// Empty active_key_version indicates that a Google-managed key should be used for signing.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_access_approval_settings#active_key_version OrganizationAccessApprovalSettings#active_key_version}
	ActiveKeyVersion *string `field:"optional" json:"activeKeyVersion" yaml:"activeKeyVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_access_approval_settings#id OrganizationAccessApprovalSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of email addresses to which notifications relating to approval requests should be sent.
	//
	// Notifications relating to a resource will be sent to all emails in the settings of ancestor
	// resources of that resource. A maximum of 50 email addresses are allowed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_access_approval_settings#notification_emails OrganizationAccessApprovalSettings#notification_emails}
	NotificationEmails *[]*string `field:"optional" json:"notificationEmails" yaml:"notificationEmails"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_access_approval_settings#timeouts OrganizationAccessApprovalSettings#timeouts}
	Timeouts *OrganizationAccessApprovalSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

