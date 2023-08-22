package pubsubsubscription


type PubsubSubscriptionBigqueryConfig struct {
	// The name of the table to which to write data, of the form {projectId}:{datasetId}.{tableId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_subscription#table PubsubSubscription#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// When true and useTopicSchema is true, any fields that are a part of the topic schema that are not part of the BigQuery table schema are dropped when writing to BigQuery.
	//
	// Otherwise, the schemas must be kept in sync and any messages with extra fields are not written and remain in the subscription's backlog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_subscription#drop_unknown_fields PubsubSubscription#drop_unknown_fields}
	DropUnknownFields interface{} `field:"optional" json:"dropUnknownFields" yaml:"dropUnknownFields"`
	// When true, use the topic's schema as the columns to write to in BigQuery, if it exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_subscription#use_topic_schema PubsubSubscription#use_topic_schema}
	UseTopicSchema interface{} `field:"optional" json:"useTopicSchema" yaml:"useTopicSchema"`
	// When true, write the subscription name, messageId, publishTime, attributes, and orderingKey to additional columns in the table.
	//
	// The subscription name, messageId, and publishTime fields are put in their own columns while all other message properties (other than data) are written to a JSON object in the attributes column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_subscription#write_metadata PubsubSubscription#write_metadata}
	WriteMetadata interface{} `field:"optional" json:"writeMetadata" yaml:"writeMetadata"`
}

