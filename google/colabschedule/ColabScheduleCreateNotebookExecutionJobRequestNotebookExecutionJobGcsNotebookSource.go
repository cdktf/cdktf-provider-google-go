// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabschedule


type ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource struct {
	// The Cloud Storage uri pointing to the ipynb file. Format: gs://bucket/notebook_file.ipynb.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#uri ColabSchedule#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The version of the Cloud Storage object to read.
	//
	// If unset, the current version of the object is read. See https://cloud.google.com/storage/docs/metadata#generation-number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#generation ColabSchedule#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

