// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ProjectAccessApprovalSettingsEnrolledServices struct {
	// The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive): all appengine.googleapis.com bigquery.googleapis.com bigtable.googleapis.com cloudkms.googleapis.com compute.googleapis.com dataflow.googleapis.com iam.googleapis.com pubsub.googleapis.com storage.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_access_approval_settings#cloud_product ProjectAccessApprovalSettings#cloud_product}
	CloudProduct *string `field:"required" json:"cloudProduct" yaml:"cloudProduct"`
	// The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_access_approval_settings#enrollment_level ProjectAccessApprovalSettings#enrollment_level}
	EnrollmentLevel *string `field:"optional" json:"enrollmentLevel" yaml:"enrollmentLevel"`
}

