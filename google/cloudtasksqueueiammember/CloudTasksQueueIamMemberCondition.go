package cloudtasksqueueiammember


type CloudTasksQueueIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_tasks_queue_iam_member#expression CloudTasksQueueIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_tasks_queue_iam_member#title CloudTasksQueueIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_tasks_queue_iam_member#description CloudTasksQueueIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

