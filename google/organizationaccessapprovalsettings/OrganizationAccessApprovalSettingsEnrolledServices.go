package organizationaccessapprovalsettings


type OrganizationAccessApprovalSettingsEnrolledServices struct {
	// The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive): all appengine.googleapis.com bigquery.googleapis.com bigtable.googleapis.com cloudkms.googleapis.com compute.googleapis.com dataflow.googleapis.com iam.googleapis.com pubsub.googleapis.com storage.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/organization_access_approval_settings#cloud_product OrganizationAccessApprovalSettings#cloud_product}
	CloudProduct *string `field:"required" json:"cloudProduct" yaml:"cloudProduct"`
	// The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/organization_access_approval_settings#enrollment_level OrganizationAccessApprovalSettings#enrollment_level}
	EnrollmentLevel *string `field:"optional" json:"enrollmentLevel" yaml:"enrollmentLevel"`
}

