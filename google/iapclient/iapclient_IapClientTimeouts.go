package iapclient


type IapClientTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_client#create IapClient#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_client#delete IapClient#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

