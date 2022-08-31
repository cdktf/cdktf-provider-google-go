// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataprocWorkflowTemplateParametersValidationValues struct {
	// Required. List of allowed values for the parameter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#values DataprocWorkflowTemplate#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

