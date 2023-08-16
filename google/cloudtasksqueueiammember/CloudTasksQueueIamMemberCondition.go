package cloudtasksqueueiammember


type CloudTasksQueueIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_tasks_queue_iam_member#expression CloudTasksQueueIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_tasks_queue_iam_member#title CloudTasksQueueIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_tasks_queue_iam_member#description CloudTasksQueueIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

