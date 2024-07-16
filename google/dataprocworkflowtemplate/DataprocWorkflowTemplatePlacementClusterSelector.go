// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementClusterSelector struct {
	// Required. The cluster labels. Cluster must have all labels to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dataproc_workflow_template#cluster_labels DataprocWorkflowTemplate#cluster_labels}
	ClusterLabels *map[string]*string `field:"required" json:"clusterLabels" yaml:"clusterLabels"`
	// Optional.
	//
	// The zone where workflow process executes. This parameter does not affect the selection of the cluster. If unspecified, the zone of the first cluster matching the selector is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dataproc_workflow_template#zone DataprocWorkflowTemplate#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

