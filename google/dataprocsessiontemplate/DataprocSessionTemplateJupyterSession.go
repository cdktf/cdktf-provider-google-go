// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocsessiontemplate


type DataprocSessionTemplateJupyterSession struct {
	// Display name, shown in the Jupyter kernelspec card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dataproc_session_template#display_name DataprocSessionTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Kernel to be used with Jupyter interactive session. Possible values: ["PYTHON", "SCALA"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dataproc_session_template#kernel DataprocSessionTemplate#kernel}
	Kernel *string `field:"optional" json:"kernel" yaml:"kernel"`
}

