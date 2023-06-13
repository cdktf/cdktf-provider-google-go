package dataproccluster


type DataprocClusterClusterConfigSoftwareConfig struct {
	// The Cloud Dataproc image version to use for the cluster - this controls the sets of software versions installed onto the nodes when you create clusters.
	//
	// If not specified, defaults to the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#image_version DataprocCluster#image_version}
	ImageVersion *string `field:"optional" json:"imageVersion" yaml:"imageVersion"`
	// The set of optional components to activate on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#optional_components DataprocCluster#optional_components}
	OptionalComponents *[]*string `field:"optional" json:"optionalComponents" yaml:"optionalComponents"`
	// A list of override and additional properties (key/value pairs) used to modify various aspects of the common configuration files used when creating a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#override_properties DataprocCluster#override_properties}
	OverrideProperties *map[string]*string `field:"optional" json:"overrideProperties" yaml:"overrideProperties"`
}

