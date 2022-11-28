package eventarcchannel


type EventarcChannelTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/eventarc_channel#create EventarcChannel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/eventarc_channel#delete EventarcChannel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/eventarc_channel#update EventarcChannel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

