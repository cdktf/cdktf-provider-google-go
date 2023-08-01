package storagetransferagentpool


type StorageTransferAgentPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_agent_pool#create StorageTransferAgentPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_agent_pool#delete StorageTransferAgentPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_agent_pool#update StorageTransferAgentPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

