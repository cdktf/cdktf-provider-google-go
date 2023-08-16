package pubsublitesubscription


type PubsubLiteSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/pubsub_lite_subscription#create PubsubLiteSubscription#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/pubsub_lite_subscription#delete PubsubLiteSubscription#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/pubsub_lite_subscription#update PubsubLiteSubscription#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

