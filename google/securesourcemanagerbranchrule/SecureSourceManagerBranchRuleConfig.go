// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerbranchrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecureSourceManagerBranchRuleConfig struct {
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
	// The ID for the BranchRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#branch_rule_id SecureSourceManagerBranchRule#branch_rule_id}
	BranchRuleId *string `field:"required" json:"branchRuleId" yaml:"branchRuleId"`
	// The BranchRule matches branches based on the specified regular expression. Use .* to match all branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#include_pattern SecureSourceManagerBranchRule#include_pattern}
	IncludePattern *string `field:"required" json:"includePattern" yaml:"includePattern"`
	// The location for the Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#location SecureSourceManagerBranchRule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID for the Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#repository_id SecureSourceManagerBranchRule#repository_id}
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Determines if allow stale reviews or approvals before merging to the branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#allow_stale_reviews SecureSourceManagerBranchRule#allow_stale_reviews}
	AllowStaleReviews interface{} `field:"optional" json:"allowStaleReviews" yaml:"allowStaleReviews"`
	// Determines if the branch rule is disabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#disabled SecureSourceManagerBranchRule#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#id SecureSourceManagerBranchRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The minimum number of approvals required for the branch rule to be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#minimum_approvals_count SecureSourceManagerBranchRule#minimum_approvals_count}
	MinimumApprovalsCount *float64 `field:"optional" json:"minimumApprovalsCount" yaml:"minimumApprovalsCount"`
	// The minimum number of reviews required for the branch rule to be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#minimum_reviews_count SecureSourceManagerBranchRule#minimum_reviews_count}
	MinimumReviewsCount *float64 `field:"optional" json:"minimumReviewsCount" yaml:"minimumReviewsCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#project SecureSourceManagerBranchRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Determines if require comments resolved before merging to the branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#require_comments_resolved SecureSourceManagerBranchRule#require_comments_resolved}
	RequireCommentsResolved interface{} `field:"optional" json:"requireCommentsResolved" yaml:"requireCommentsResolved"`
	// Determines if require linear history before merging to the branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#require_linear_history SecureSourceManagerBranchRule#require_linear_history}
	RequireLinearHistory interface{} `field:"optional" json:"requireLinearHistory" yaml:"requireLinearHistory"`
	// Determines if the branch rule requires a pull request or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#require_pull_request SecureSourceManagerBranchRule#require_pull_request}
	RequirePullRequest interface{} `field:"optional" json:"requirePullRequest" yaml:"requirePullRequest"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_branch_rule#timeouts SecureSourceManagerBranchRule#timeouts}
	Timeouts *SecureSourceManagerBranchRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

