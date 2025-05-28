// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2job


type CloudRunV2JobTemplateTemplateVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/cloud_run_v2_job#name CloudRunV2Job#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_sql_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/cloud_run_v2_job#cloud_sql_instance CloudRunV2Job#cloud_sql_instance}
	CloudSqlInstance *CloudRunV2JobTemplateTemplateVolumesCloudSqlInstance `field:"optional" json:"cloudSqlInstance" yaml:"cloudSqlInstance"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/cloud_run_v2_job#empty_dir CloudRunV2Job#empty_dir}
	EmptyDir *CloudRunV2JobTemplateTemplateVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/cloud_run_v2_job#gcs CloudRunV2Job#gcs}
	Gcs *CloudRunV2JobTemplateTemplateVolumesGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/cloud_run_v2_job#nfs CloudRunV2Job#nfs}
	Nfs *CloudRunV2JobTemplateTemplateVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/cloud_run_v2_job#secret CloudRunV2Job#secret}
	Secret *CloudRunV2JobTemplateTemplateVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

