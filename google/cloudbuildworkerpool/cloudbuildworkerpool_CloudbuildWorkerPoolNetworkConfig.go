package cloudbuildworkerpool


type CloudbuildWorkerPoolNetworkConfig struct {
	// Required.
	//
	// Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to `WorkerPool.project_id` on the service producer network. Must be in the format `projects/{project}/global/networks/{network}`, where `{project}` is a project number, such as `12345`, and `{network}` is the name of a VPC network in the project. See [Understanding network configuration options](https://cloud.google.com/cloud-build/docs/custom-workers/set-up-custom-worker-pool-environment#understanding_the_network_configuration_options)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_worker_pool#peered_network CloudbuildWorkerPool#peered_network}
	PeeredNetwork *string `field:"required" json:"peeredNetwork" yaml:"peeredNetwork"`
}

