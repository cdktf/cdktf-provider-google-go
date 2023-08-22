package iapclient


type IapClientTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_client#create IapClient#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_client#delete IapClient#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

