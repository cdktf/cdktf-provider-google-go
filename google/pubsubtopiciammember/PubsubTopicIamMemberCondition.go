package pubsubtopiciammember


type PubsubTopicIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_topic_iam_member#expression PubsubTopicIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_topic_iam_member#title PubsubTopicIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_topic_iam_member#description PubsubTopicIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

