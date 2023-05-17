package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfig struct {
	// custom_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#custom_info_types DataLossPreventionJobTrigger#custom_info_types}
	CustomInfoTypes interface{} `field:"optional" json:"customInfoTypes" yaml:"customInfoTypes"`
	// When true, excludes type information of the findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#exclude_info_types DataLossPreventionJobTrigger#exclude_info_types}
	ExcludeInfoTypes interface{} `field:"optional" json:"excludeInfoTypes" yaml:"excludeInfoTypes"`
	// When true, a contextual quote from the data that triggered a finding is included in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#include_quote DataLossPreventionJobTrigger#include_quote}
	IncludeQuote interface{} `field:"optional" json:"includeQuote" yaml:"includeQuote"`
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#info_types DataLossPreventionJobTrigger#info_types}
	InfoTypes interface{} `field:"optional" json:"infoTypes" yaml:"infoTypes"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#limits DataLossPreventionJobTrigger#limits}
	Limits *DataLossPreventionJobTriggerInspectJobInspectConfigLimits `field:"optional" json:"limits" yaml:"limits"`
	// Only returns findings equal or above this threshold.
	//
	// See https://cloud.google.com/dlp/docs/likelihood for more info Default value: "POSSIBLE" Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#min_likelihood DataLossPreventionJobTrigger#min_likelihood}
	MinLikelihood *string `field:"optional" json:"minLikelihood" yaml:"minLikelihood"`
	// rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#rule_set DataLossPreventionJobTrigger#rule_set}
	RuleSet interface{} `field:"optional" json:"ruleSet" yaml:"ruleSet"`
}

