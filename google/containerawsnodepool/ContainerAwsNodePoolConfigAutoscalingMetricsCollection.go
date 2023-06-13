package containerawsnodepool


type ContainerAwsNodePoolConfigAutoscalingMetricsCollection struct {
	// The frequency at which EC2 Auto Scaling sends aggregated data to AWS CloudWatch. The only valid value is "1Minute".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_aws_node_pool#granularity ContainerAwsNodePool#granularity}
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// The metrics to enable.
	//
	// For a list of valid metrics, see https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_EnableMetricsCollection.html. If you specify granularity and don't specify any metrics, all metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_aws_node_pool#metrics ContainerAwsNodePool#metrics}
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

