package dataproccluster


type DataprocClusterVirtualClusterConfigKubernetesClusterConfig struct {
	// gke_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#gke_cluster_config DataprocCluster#gke_cluster_config}
	GkeClusterConfig *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig `field:"required" json:"gkeClusterConfig" yaml:"gkeClusterConfig"`
	// kubernetes_software_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#kubernetes_software_config DataprocCluster#kubernetes_software_config}
	KubernetesSoftwareConfig *DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig `field:"required" json:"kubernetesSoftwareConfig" yaml:"kubernetesSoftwareConfig"`
	// A namespace within the Kubernetes cluster to deploy into.
	//
	// If this namespace does not exist, it is created. If it exists, Dataproc verifies that another Dataproc VirtualCluster is not installed into it. If not specified, the name of the Dataproc Cluster is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#kubernetes_namespace DataprocCluster#kubernetes_namespace}
	KubernetesNamespace *string `field:"optional" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

