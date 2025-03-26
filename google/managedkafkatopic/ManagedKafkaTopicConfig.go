// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedkafkatopic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedKafkaTopicConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#cluster ManagedKafkaTopic#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// ID of the location of the Kafka resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#location ManagedKafkaTopic#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The number of replicas of each partition. A replication factor of 3 is recommended for high availability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#replication_factor ManagedKafkaTopic#replication_factor}
	ReplicationFactor *float64 `field:"required" json:"replicationFactor" yaml:"replicationFactor"`
	// The ID to use for the topic, which will become the final component of the topic's name.
	//
	// This value is structured like: 'my-topic-name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#topic_id ManagedKafkaTopic#topic_id}
	TopicId *string `field:"required" json:"topicId" yaml:"topicId"`
	// Configuration for the topic that are overridden from the cluster defaults.
	//
	// The key of the map is a Kafka topic property name, for example: 'cleanup.policy=compact', 'compression.type=producer'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#configs ManagedKafkaTopic#configs}
	Configs *map[string]*string `field:"optional" json:"configs" yaml:"configs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#id ManagedKafkaTopic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of partitions in a topic.
	//
	// You can increase the partition count for a topic, but you cannot decrease it. Increasing partitions for a topic that uses a key might change how messages are distributed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#partition_count ManagedKafkaTopic#partition_count}
	PartitionCount *float64 `field:"optional" json:"partitionCount" yaml:"partitionCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#project ManagedKafkaTopic#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_topic#timeouts ManagedKafkaTopic#timeouts}
	Timeouts *ManagedKafkaTopicTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

