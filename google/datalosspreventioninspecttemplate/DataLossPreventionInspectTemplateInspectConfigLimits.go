package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigLimits struct {
	// Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_inspect_template#max_findings_per_item DataLossPreventionInspectTemplate#max_findings_per_item}
	MaxFindingsPerItem *float64 `field:"required" json:"maxFindingsPerItem" yaml:"maxFindingsPerItem"`
	// Max number of findings that will be returned per request/job. The maximum returned is 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_inspect_template#max_findings_per_request DataLossPreventionInspectTemplate#max_findings_per_request}
	MaxFindingsPerRequest *float64 `field:"required" json:"maxFindingsPerRequest" yaml:"maxFindingsPerRequest"`
	// max_findings_per_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_inspect_template#max_findings_per_info_type DataLossPreventionInspectTemplate#max_findings_per_info_type}
	MaxFindingsPerInfoType interface{} `field:"optional" json:"maxFindingsPerInfoType" yaml:"maxFindingsPerInfoType"`
}

