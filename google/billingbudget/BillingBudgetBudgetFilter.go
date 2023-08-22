package billingbudget


type BillingBudgetBudgetFilter struct {
	// A CalendarPeriod represents the abstract concept of a recurring time period that has a canonical start.
	//
	// Grammatically, "the start of the current CalendarPeriod".
	// All calendar times begin at 12 AM US and Canadian Pacific Time (UTC-8).
	//
	// Exactly one of 'calendar_period', 'custom_period' must be provided. Possible values: ["MONTH", "QUARTER", "YEAR", "CALENDAR_PERIOD_UNSPECIFIED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#calendar_period BillingBudget#calendar_period}
	CalendarPeriod *string `field:"optional" json:"calendarPeriod" yaml:"calendarPeriod"`
	// Optional.
	//
	// If creditTypesTreatment is INCLUDE_SPECIFIED_CREDITS,
	// this is a list of credit types to be subtracted from gross cost to determine the spend for threshold calculations. See a list of acceptable credit type values.
	// If creditTypesTreatment is not INCLUDE_SPECIFIED_CREDITS, this field must be empty.
	//
	// *Note:** If the field has a value in the config and needs to be removed, the field has to be an emtpy array in the config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#credit_types BillingBudget#credit_types}
	CreditTypes *[]*string `field:"optional" json:"creditTypes" yaml:"creditTypes"`
	// Specifies how credits should be treated when determining spend for threshold calculations. Default value: "INCLUDE_ALL_CREDITS" Possible values: ["INCLUDE_ALL_CREDITS", "EXCLUDE_ALL_CREDITS", "INCLUDE_SPECIFIED_CREDITS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#credit_types_treatment BillingBudget#credit_types_treatment}
	CreditTypesTreatment *string `field:"optional" json:"creditTypesTreatment" yaml:"creditTypesTreatment"`
	// custom_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#custom_period BillingBudget#custom_period}
	CustomPeriod *BillingBudgetBudgetFilterCustomPeriod `field:"optional" json:"customPeriod" yaml:"customPeriod"`
	// A single label and value pair specifying that usage from only this set of labeled resources should be included in the budget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#labels BillingBudget#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// A set of projects of the form projects/{project_number}, specifying that usage from only this set of projects should be included in the budget.
	//
	// If omitted, the report will include
	// all usage for the billing account, regardless of which project
	// the usage occurred on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#projects BillingBudget#projects}
	Projects *[]*string `field:"optional" json:"projects" yaml:"projects"`
	// A set of folder and organization names of the form folders/{folderId} or organizations/{organizationId}, specifying that usage from only this set of folders and organizations should be included in the budget.
	//
	// If omitted, the budget includes all usage that the billing account pays for. If the folder or organization
	// contains projects that are paid for by a different Cloud Billing account, the budget doesn't apply to those projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#resource_ancestors BillingBudget#resource_ancestors}
	ResourceAncestors *[]*string `field:"optional" json:"resourceAncestors" yaml:"resourceAncestors"`
	// A set of services of the form services/{service_id}, specifying that usage from only this set of services should be included in the budget.
	//
	// If omitted, the report will include
	// usage for all the services. The service names are available
	// through the Catalog API:
	// https://cloud.google.com/billing/v1/how-tos/catalog-api.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#services BillingBudget#services}
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// A set of subaccounts of the form billingAccounts/{account_id}, specifying that usage from only this set of subaccounts should be included in the budget.
	//
	// If a subaccount is set to the name of
	// the parent account, usage from the parent account will be included.
	// If the field is omitted, the report will include usage from the parent
	// account and all subaccounts, if they exist.
	//
	// *Note:** If the field has a value in the config and needs to be removed, the field has to be an emtpy array in the config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/billing_budget#subaccounts BillingBudget#subaccounts}
	Subaccounts *[]*string `field:"optional" json:"subaccounts" yaml:"subaccounts"`
}

