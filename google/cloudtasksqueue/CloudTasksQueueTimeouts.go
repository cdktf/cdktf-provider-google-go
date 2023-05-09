package cloudtasksqueue


type CloudTasksQueueTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/cloud_tasks_queue#create CloudTasksQueue#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/cloud_tasks_queue#delete CloudTasksQueue#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/cloud_tasks_queue#update CloudTasksQueue#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

