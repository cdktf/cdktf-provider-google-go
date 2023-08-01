package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig struct {
	// Optional. The time when cluster will be auto-deleted (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_workflow_template#auto_delete_time DataprocWorkflowTemplate#auto_delete_time}
	AutoDeleteTime *string `field:"optional" json:"autoDeleteTime" yaml:"autoDeleteTime"`
	// Optional.
	//
	// The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_workflow_template#auto_delete_ttl DataprocWorkflowTemplate#auto_delete_ttl}
	AutoDeleteTtl *string `field:"optional" json:"autoDeleteTtl" yaml:"autoDeleteTtl"`
	// Optional.
	//
	// The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_workflow_template#idle_delete_ttl DataprocWorkflowTemplate#idle_delete_ttl}
	IdleDeleteTtl *string `field:"optional" json:"idleDeleteTtl" yaml:"idleDeleteTtl"`
}

