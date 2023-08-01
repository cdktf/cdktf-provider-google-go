package pubsubtopiciambinding


type PubsubTopicIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/pubsub_topic_iam_binding#expression PubsubTopicIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/pubsub_topic_iam_binding#title PubsubTopicIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/pubsub_topic_iam_binding#description PubsubTopicIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

