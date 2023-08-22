package pubsubsubscriptioniambinding


type PubsubSubscriptionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_subscription_iam_binding#expression PubsubSubscriptionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_subscription_iam_binding#title PubsubSubscriptionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_subscription_iam_binding#description PubsubSubscriptionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

