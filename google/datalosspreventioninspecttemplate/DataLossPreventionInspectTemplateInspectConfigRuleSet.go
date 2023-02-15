package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigRuleSet struct {
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#info_types DataLossPreventionInspectTemplate#info_types}
	InfoTypes interface{} `field:"required" json:"infoTypes" yaml:"infoTypes"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#rules DataLossPreventionInspectTemplate#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

