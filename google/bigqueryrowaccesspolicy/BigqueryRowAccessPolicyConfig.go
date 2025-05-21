// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryrowaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryRowAccessPolicyConfig struct {
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
	// The ID of the dataset containing this row access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/bigquery_row_access_policy#dataset_id BigqueryRowAccessPolicy#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// A SQL boolean expression that represents the rows defined by this row access policy, similar to the boolean expression in a WHERE clause of a SELECT query on a table.
	//
	// References to other tables, routines, and temporary functions are not
	// supported.
	//
	// Examples: region="EU"
	// date_field = CAST('2019-9-27' as DATE)
	// nullable_field is not NULL
	// numeric_field BETWEEN 1.0 AND 5.0
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/bigquery_row_access_policy#filter_predicate BigqueryRowAccessPolicy#filter_predicate}
	FilterPredicate *string `field:"required" json:"filterPredicate" yaml:"filterPredicate"`
	// The ID of the row access policy.
	//
	// The ID must contain only
	// letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/bigquery_row_access_policy#policy_id BigqueryRowAccessPolicy#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// The ID of the table containing this row access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/bigquery_row_access_policy#table_id BigqueryRowAccessPolicy#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// Input only.
	//
	// The optional list of iam_member users or groups that specifies the initial
	// members that the row-level access policy should be created with.
	//
	// grantees types:
	// - "user:alice@example.com": An email address that represents a specific
	// Google account.
	// - "serviceAccount:my-other-app@appspot.gserviceaccount.com": An email
	// address that represents a service account.
	// - "group:admins@example.com": An email address that represents a Google
	// group.
	// - "domain:example.com":The Google Workspace domain (primary) that
	// represents all the users of that domain.
	// - "allAuthenticatedUsers": A special identifier that represents all service
	// accounts and all users on the internet who have authenticated with a Google
	// Account. This identifier includes accounts that aren't connected to a
	// Google Workspace or Cloud Identity domain, such as personal Gmail accounts.
	// Users who aren't authenticated, such as anonymous visitors, aren't
	// included.
	// - "allUsers":A special identifier that represents anyone who is on
	// the internet, including authenticated and unauthenticated users. Because
	// BigQuery requires authentication before a user can access the service,
	// allUsers includes only authenticated users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/bigquery_row_access_policy#grantees BigqueryRowAccessPolicy#grantees}
	Grantees *[]*string `field:"optional" json:"grantees" yaml:"grantees"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/bigquery_row_access_policy#id BigqueryRowAccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/bigquery_row_access_policy#project BigqueryRowAccessPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/bigquery_row_access_policy#timeouts BigqueryRowAccessPolicy#timeouts}
	Timeouts *BigqueryRowAccessPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

