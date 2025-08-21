// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedkafkacluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedKafkaClusterConfig struct {
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
	// capacity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#capacity_config ManagedKafkaCluster#capacity_config}
	CapacityConfig *ManagedKafkaClusterCapacityConfig `field:"required" json:"capacityConfig" yaml:"capacityConfig"`
	// The ID to use for the cluster, which will become the final component of the cluster's name.
	//
	// The ID must be 1-63 characters long, and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' to comply with RFC 1035. This value is structured like: 'my-cluster-id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#cluster_id ManagedKafkaCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// gcp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#gcp_config ManagedKafkaCluster#gcp_config}
	GcpConfig *ManagedKafkaClusterGcpConfig `field:"required" json:"gcpConfig" yaml:"gcpConfig"`
	// ID of the location of the Kafka resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#location ManagedKafkaCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#id ManagedKafkaCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of label KEY=VALUE pairs to add.
	//
	// Keys must start with a lowercase character and contain only hyphens (-), underscores ( ), lowercase characters, and numbers. Values must contain only hyphens (-), underscores ( ), lowercase characters, and numbers.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#labels ManagedKafkaCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#project ManagedKafkaCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rebalance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#rebalance_config ManagedKafkaCluster#rebalance_config}
	RebalanceConfig *ManagedKafkaClusterRebalanceConfig `field:"optional" json:"rebalanceConfig" yaml:"rebalanceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#timeouts ManagedKafkaCluster#timeouts}
	Timeouts *ManagedKafkaClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/managed_kafka_cluster#tls_config ManagedKafkaCluster#tls_config}
	TlsConfig *ManagedKafkaClusterTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

