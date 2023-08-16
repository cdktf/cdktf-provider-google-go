package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacement struct {
	// cluster_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#cluster_selector DataprocWorkflowTemplate#cluster_selector}
	ClusterSelector *DataprocWorkflowTemplatePlacementClusterSelector `field:"optional" json:"clusterSelector" yaml:"clusterSelector"`
	// managed_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#managed_cluster DataprocWorkflowTemplate#managed_cluster}
	ManagedCluster *DataprocWorkflowTemplatePlacementManagedCluster `field:"optional" json:"managedCluster" yaml:"managedCluster"`
}

