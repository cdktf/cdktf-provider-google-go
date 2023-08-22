package dataproccluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterConfig struct {
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
	// The name of the cluster, unique within the project and zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#name DataprocCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#cluster_config DataprocCluster#cluster_config}
	ClusterConfig *DataprocClusterClusterConfig `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a terraform apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#graceful_decommission_timeout DataprocCluster#graceful_decommission_timeout}
	GracefulDecommissionTimeout *string `field:"optional" json:"gracefulDecommissionTimeout" yaml:"gracefulDecommissionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#id DataprocCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The list of labels (key/value pairs) to be applied to instances in the cluster.
	//
	// GCP generates some itself including goog-dataproc-cluster-name which is the name of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#labels DataprocCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The ID of the project in which the cluster will exist.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#project DataprocCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region in which the cluster and associated nodes will be created in. Defaults to global.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#region DataprocCluster#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#timeouts DataprocCluster#timeouts}
	Timeouts *DataprocClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#virtual_cluster_config DataprocCluster#virtual_cluster_config}
	VirtualClusterConfig *DataprocClusterVirtualClusterConfig `field:"optional" json:"virtualClusterConfig" yaml:"virtualClusterConfig"`
}

