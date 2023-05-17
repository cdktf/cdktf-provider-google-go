package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity struct {
	// Required.
	//
	// The URI of a sole-tenant [node group resource](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups) that the cluster will be created on. A full URL, partial URI, or node group name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `node-group-1`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#node_group DataprocWorkflowTemplate#node_group}
	NodeGroup *string `field:"required" json:"nodeGroup" yaml:"nodeGroup"`
}

