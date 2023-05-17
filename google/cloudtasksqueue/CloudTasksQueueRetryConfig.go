package cloudtasksqueue


type CloudTasksQueueRetryConfig struct {
	// Number of attempts per task.
	//
	// Cloud Tasks will attempt the task maxAttempts times (that is, if
	// the first attempt fails, then there will be maxAttempts - 1
	// retries). Must be >= -1.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the default.
	//
	// -1 indicates unlimited attempts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_tasks_queue#max_attempts CloudTasksQueue#max_attempts}
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// A task will be scheduled for retry between minBackoff and maxBackoff duration after it fails, if the queue's RetryConfig specifies that the task should be retried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_tasks_queue#max_backoff CloudTasksQueue#max_backoff}
	MaxBackoff *string `field:"optional" json:"maxBackoff" yaml:"maxBackoff"`
	// The time between retries will double maxDoublings times.
	//
	// A task's retry interval starts at minBackoff, then doubles maxDoublings times,
	// then increases linearly, and finally retries retries at intervals of maxBackoff
	// up to maxAttempts times.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_tasks_queue#max_doublings CloudTasksQueue#max_doublings}
	MaxDoublings *float64 `field:"optional" json:"maxDoublings" yaml:"maxDoublings"`
	// If positive, maxRetryDuration specifies the time limit for retrying a failed task, measured from when the task was first attempted.
	//
	// Once maxRetryDuration time has passed and the task has
	// been attempted maxAttempts times, no further attempts will be
	// made and the task will be deleted.
	//
	// If zero, then the task age is unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_tasks_queue#max_retry_duration CloudTasksQueue#max_retry_duration}
	MaxRetryDuration *string `field:"optional" json:"maxRetryDuration" yaml:"maxRetryDuration"`
	// A task will be scheduled for retry between minBackoff and maxBackoff duration after it fails, if the queue's RetryConfig specifies that the task should be retried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_tasks_queue#min_backoff CloudTasksQueue#min_backoff}
	MinBackoff *string `field:"optional" json:"minBackoff" yaml:"minBackoff"`
}

