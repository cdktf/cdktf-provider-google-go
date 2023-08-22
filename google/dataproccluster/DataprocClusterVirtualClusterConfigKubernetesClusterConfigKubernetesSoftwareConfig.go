package dataproccluster


type DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig struct {
	// The components that should be installed in this Dataproc cluster.
	//
	// The key must be a string from the KubernetesComponent enumeration. The value is the version of the software to be installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#component_version DataprocCluster#component_version}
	ComponentVersion *map[string]*string `field:"required" json:"componentVersion" yaml:"componentVersion"`
	// The properties to set on daemon config files. Property keys are specified in prefix:property format, for example spark:spark.kubernetes.container.image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_cluster#properties DataprocCluster#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

