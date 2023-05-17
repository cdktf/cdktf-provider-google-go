package pubsublitetopic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PubsubLiteTopicConfig struct {
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
	// Name of the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#name PubsubLiteTopic#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#id PubsubLiteTopic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// partition_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#partition_config PubsubLiteTopic#partition_config}
	PartitionConfig *PubsubLiteTopicPartitionConfig `field:"optional" json:"partitionConfig" yaml:"partitionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#project PubsubLiteTopic#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the pubsub lite topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#region PubsubLiteTopic#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// reservation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#reservation_config PubsubLiteTopic#reservation_config}
	ReservationConfig *PubsubLiteTopicReservationConfig `field:"optional" json:"reservationConfig" yaml:"reservationConfig"`
	// retention_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#retention_config PubsubLiteTopic#retention_config}
	RetentionConfig *PubsubLiteTopicRetentionConfig `field:"optional" json:"retentionConfig" yaml:"retentionConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#timeouts PubsubLiteTopic#timeouts}
	Timeouts *PubsubLiteTopicTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The zone of the pubsub lite topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/pubsub_lite_topic#zone PubsubLiteTopic#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

