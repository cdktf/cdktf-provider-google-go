package pubsubschema


type PubsubSchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/pubsub_schema#create PubsubSchema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/pubsub_schema#delete PubsubSchema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

