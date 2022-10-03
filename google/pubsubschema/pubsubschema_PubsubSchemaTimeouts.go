package pubsubschema


type PubsubSchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_schema#create PubsubSchema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_schema#delete PubsubSchema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

