package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfig struct {
	// List of options defining data content to scan.
	//
	// If empty, text, images, and other content will be included. Possible values: ["CONTENT_TEXT", "CONTENT_IMAGE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_inspect_template#content_options DataLossPreventionInspectTemplate#content_options}
	ContentOptions *[]*string `field:"optional" json:"contentOptions" yaml:"contentOptions"`
	// custom_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_inspect_template#custom_info_types DataLossPreventionInspectTemplate#custom_info_types}
	CustomInfoTypes interface{} `field:"optional" json:"customInfoTypes" yaml:"customInfoTypes"`
	// When true, excludes type information of the findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_inspect_template#exclude_info_types DataLossPreventionInspectTemplate#exclude_info_types}
	ExcludeInfoTypes interface{} `field:"optional" json:"excludeInfoTypes" yaml:"excludeInfoTypes"`
	// When true, a contextual quote from the data that triggered a finding is included in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_inspect_template#include_quote DataLossPreventionInspectTemplate#include_quote}
	IncludeQuote interface{} `field:"optional" json:"includeQuote" yaml:"includeQuote"`
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_inspect_template#info_types DataLossPreventionInspectTemplate#info_types}
	InfoTypes interface{} `field:"optional" json:"infoTypes" yaml:"infoTypes"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_inspect_template#limits DataLossPreventionInspectTemplate#limits}
	Limits *DataLossPreventionInspectTemplateInspectConfigLimits `field:"optional" json:"limits" yaml:"limits"`
	// Only returns findings equal or above this threshold.
	//
	// See https://cloud.google.com/dlp/docs/likelihood for more info Default value: "POSSIBLE" Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_inspect_template#min_likelihood DataLossPreventionInspectTemplate#min_likelihood}
	MinLikelihood *string `field:"optional" json:"minLikelihood" yaml:"minLikelihood"`
	// rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_inspect_template#rule_set DataLossPreventionInspectTemplate#rule_set}
	RuleSet interface{} `field:"optional" json:"ruleSet" yaml:"ruleSet"`
}

