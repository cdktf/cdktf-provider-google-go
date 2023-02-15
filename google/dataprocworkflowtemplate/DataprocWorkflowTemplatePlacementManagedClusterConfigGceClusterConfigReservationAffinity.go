package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity struct {
	// Optional. Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#consume_reservation_type DataprocWorkflowTemplate#consume_reservation_type}
	ConsumeReservationType *string `field:"optional" json:"consumeReservationType" yaml:"consumeReservationType"`
	// Optional. Corresponds to the label key of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#key DataprocWorkflowTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Optional. Corresponds to the label values of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#values DataprocWorkflowTemplate#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

