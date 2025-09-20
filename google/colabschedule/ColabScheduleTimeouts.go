// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabschedule


type ColabScheduleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#create ColabSchedule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#delete ColabSchedule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#update ColabSchedule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

