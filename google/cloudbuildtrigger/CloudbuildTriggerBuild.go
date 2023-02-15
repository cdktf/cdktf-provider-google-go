package cloudbuildtrigger


type CloudbuildTriggerBuild struct {
	// step block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#step CloudbuildTrigger#step}
	Step interface{} `field:"required" json:"step" yaml:"step"`
	// artifacts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#artifacts CloudbuildTrigger#artifacts}
	Artifacts *CloudbuildTriggerBuildArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	// available_secrets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#available_secrets CloudbuildTrigger#available_secrets}
	AvailableSecrets *CloudbuildTriggerBuildAvailableSecrets `field:"optional" json:"availableSecrets" yaml:"availableSecrets"`
	// A list of images to be pushed upon the successful completion of all build steps.
	//
	// The images are pushed using the builder service account's credentials.
	// The digests of the pushed images will be stored in the Build resource's results field.
	// If any of the images fail to be pushed, the build status is marked FAILURE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#images CloudbuildTrigger#images}
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	// Google Cloud Storage bucket where logs should be written.  Logs file names will be of the format ${logsBucket}/log-${build_id}.txt.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#logs_bucket CloudbuildTrigger#logs_bucket}
	LogsBucket *string `field:"optional" json:"logsBucket" yaml:"logsBucket"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#options CloudbuildTrigger#options}
	Options *CloudbuildTriggerBuildOptions `field:"optional" json:"options" yaml:"options"`
	// TTL in queue for this build.
	//
	// If provided and the build is enqueued longer than this value,
	// the build will expire and the build status will be EXPIRED.
	// The TTL starts ticking from createTime.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#queue_ttl CloudbuildTrigger#queue_ttl}
	QueueTtl *string `field:"optional" json:"queueTtl" yaml:"queueTtl"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#secret CloudbuildTrigger#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#source CloudbuildTrigger#source}
	Source *CloudbuildTriggerBuildSource `field:"optional" json:"source" yaml:"source"`
	// Substitutions data for Build resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#substitutions CloudbuildTrigger#substitutions}
	Substitutions *map[string]*string `field:"optional" json:"substitutions" yaml:"substitutions"`
	// Tags for annotation of a Build. These are not docker tags.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#tags CloudbuildTrigger#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Amount of time that this build should be allowed to run, to second granularity.
	//
	// If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.
	// This timeout must be equal to or greater than the sum of the timeouts for build steps within the build.
	// The expected format is the number of seconds followed by s.
	// Default time is ten minutes (600s).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#timeout CloudbuildTrigger#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

