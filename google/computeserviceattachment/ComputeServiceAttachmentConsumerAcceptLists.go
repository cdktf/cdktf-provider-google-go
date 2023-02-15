package computeserviceattachment


type ComputeServiceAttachmentConsumerAcceptLists struct {
	// The number of consumer forwarding rules the consumer project can create.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_service_attachment#connection_limit ComputeServiceAttachment#connection_limit}
	ConnectionLimit *float64 `field:"required" json:"connectionLimit" yaml:"connectionLimit"`
	// A project that is allowed to connect to this service attachment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_service_attachment#project_id_or_num ComputeServiceAttachment#project_id_or_num}
	ProjectIdOrNum *string `field:"required" json:"projectIdOrNum" yaml:"projectIdOrNum"`
}

