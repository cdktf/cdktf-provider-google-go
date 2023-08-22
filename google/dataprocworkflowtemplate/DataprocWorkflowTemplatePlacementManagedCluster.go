package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedCluster struct {
	// Required.
	//
	// The cluster name prefix. A unique cluster name will be formed by appending a random suffix. The name must contain only lower-case letters (a-z), numbers (0-9), and hyphens (-). Must begin with a letter. Cannot begin or end with hyphen. Must consist of between 2 and 35 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#cluster_name DataprocWorkflowTemplate#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#config DataprocWorkflowTemplate#config}
	Config *DataprocWorkflowTemplatePlacementManagedClusterConfig `field:"required" json:"config" yaml:"config"`
	// Optional.
	//
	// The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: p{Ll}p{Lo}{0,62} Label values must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: [p{Ll}p{Lo}p{N}_-]{0,63} No more than 32 labels can be associated with a given cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#labels DataprocWorkflowTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

