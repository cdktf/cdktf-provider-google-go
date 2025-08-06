// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#name CloudRunV2WorkerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_sql_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#cloud_sql_instance CloudRunV2WorkerPool#cloud_sql_instance}
	CloudSqlInstance *CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance `field:"optional" json:"cloudSqlInstance" yaml:"cloudSqlInstance"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#empty_dir CloudRunV2WorkerPool#empty_dir}
	EmptyDir *CloudRunV2WorkerPoolTemplateVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#gcs CloudRunV2WorkerPool#gcs}
	Gcs *CloudRunV2WorkerPoolTemplateVolumesGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#nfs CloudRunV2WorkerPool#nfs}
	Nfs *CloudRunV2WorkerPoolTemplateVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/cloud_run_v2_worker_pool#secret CloudRunV2WorkerPool#secret}
	Secret *CloudRunV2WorkerPoolTemplateVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

