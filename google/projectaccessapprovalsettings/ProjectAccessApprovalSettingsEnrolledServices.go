// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectaccessapprovalsettings


type ProjectAccessApprovalSettingsEnrolledServices struct {
	// The product for which Access Approval will be enrolled.
	//
	// Allowed values are listed (case-sensitive):
	//   all
	//   appengine.googleapis.com
	//   bigquery.googleapis.com
	//   bigtable.googleapis.com
	//   cloudkms.googleapis.com
	//   compute.googleapis.com
	//   dataflow.googleapis.com
	//   iam.googleapis.com
	//   pubsub.googleapis.com
	//   storage.googleapis.com
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/project_access_approval_settings#cloud_product ProjectAccessApprovalSettings#cloud_product}
	CloudProduct *string `field:"required" json:"cloudProduct" yaml:"cloudProduct"`
	// The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/project_access_approval_settings#enrollment_level ProjectAccessApprovalSettings#enrollment_level}
	EnrollmentLevel *string `field:"optional" json:"enrollmentLevel" yaml:"enrollmentLevel"`
}

