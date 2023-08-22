package cloudrunservice


type CloudRunServiceTemplateSpecContainersResources struct {
	// Limits describes the maximum amount of compute resources allowed.
	//
	// The values of the map is string form of the 'quantity' k8s type:
	// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#limits CloudRunService#limits}
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is
	// explicitly specified, otherwise to an implementation-defined value.
	// The values of the map is string form of the 'quantity' k8s type:
	// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#requests CloudRunService#requests}
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

