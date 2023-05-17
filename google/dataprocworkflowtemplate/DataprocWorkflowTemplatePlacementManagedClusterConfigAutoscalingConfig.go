package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig struct {
	// Optional.
	//
	// The autoscaling policy used by the cluster. Only resource names including projectid and location (region) are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` * `projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` Note that the policy must be in the same project and Dataproc region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#policy DataprocWorkflowTemplate#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

