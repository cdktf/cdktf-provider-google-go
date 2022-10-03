package dataprocjob


type DataprocJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_job#create DataprocJob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_job#delete DataprocJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

