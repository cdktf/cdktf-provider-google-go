package dataprocworkflowtemplate


type DataprocWorkflowTemplateParametersValidationValues struct {
	// Required. List of allowed values for the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#values DataprocWorkflowTemplate#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

