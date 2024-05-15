// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingbudget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BillingBudgetConfig struct {
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
	// amount block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#amount BillingBudget#amount}
	Amount *BillingBudgetAmount `field:"required" json:"amount" yaml:"amount"`
	// ID of the billing account to set a budget on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#billing_account BillingBudget#billing_account}
	BillingAccount *string `field:"required" json:"billingAccount" yaml:"billingAccount"`
	// all_updates_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#all_updates_rule BillingBudget#all_updates_rule}
	AllUpdatesRule *BillingBudgetAllUpdatesRule `field:"optional" json:"allUpdatesRule" yaml:"allUpdatesRule"`
	// budget_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#budget_filter BillingBudget#budget_filter}
	BudgetFilter *BillingBudgetBudgetFilter `field:"optional" json:"budgetFilter" yaml:"budgetFilter"`
	// User data for display name in UI. Must be <= 60 chars.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#display_name BillingBudget#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#id BillingBudget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ownership scope of the budget.
	//
	// The ownership scope and users'
	// IAM permissions determine who has full access to the budget's data. Possible values: ["OWNERSHIP_SCOPE_UNSPECIFIED", "ALL_USERS", "BILLING_ACCOUNT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#ownership_scope BillingBudget#ownership_scope}
	OwnershipScope *string `field:"optional" json:"ownershipScope" yaml:"ownershipScope"`
	// threshold_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#threshold_rules BillingBudget#threshold_rules}
	ThresholdRules interface{} `field:"optional" json:"thresholdRules" yaml:"thresholdRules"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/billing_budget#timeouts BillingBudget#timeouts}
	Timeouts *BillingBudgetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

