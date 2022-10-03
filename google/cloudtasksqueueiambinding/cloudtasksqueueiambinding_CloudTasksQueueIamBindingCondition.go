package cloudtasksqueueiambinding


type CloudTasksQueueIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_tasks_queue_iam_binding#expression CloudTasksQueueIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_tasks_queue_iam_binding#title CloudTasksQueueIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_tasks_queue_iam_binding#description CloudTasksQueueIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

