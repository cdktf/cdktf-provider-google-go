package containerattachedcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAttachedClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Kubernetes distribution of the underlying attached cluster. Supported values: "eks", "aks".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#distribution ContainerAttachedCluster#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#fleet ContainerAttachedCluster#fleet}
	Fleet *ContainerAttachedClusterFleet `field:"required" json:"fleet" yaml:"fleet"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#location ContainerAttachedCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#name ContainerAttachedCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// oidc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#oidc_config ContainerAttachedCluster#oidc_config}
	OidcConfig *ContainerAttachedClusterOidcConfig `field:"required" json:"oidcConfig" yaml:"oidcConfig"`
	// The platform version for the cluster (e.g. '1.23.0-gke.1').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#platform_version ContainerAttachedCluster#platform_version}
	PlatformVersion *string `field:"required" json:"platformVersion" yaml:"platformVersion"`
	// Optional.
	//
	// Annotations on the cluster. This field has the same
	// restrictions as Kubernetes annotations. The total size of all keys and
	// values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	// and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#annotations ContainerAttachedCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#authorization ContainerAttachedCluster#authorization}
	Authorization *ContainerAttachedClusterAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
	// Policy to determine what flags to send on delete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#deletion_policy ContainerAttachedCluster#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#description ContainerAttachedCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#id ContainerAttachedCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#logging_config ContainerAttachedCluster#logging_config}
	LoggingConfig *ContainerAttachedClusterLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// monitoring_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#monitoring_config ContainerAttachedCluster#monitoring_config}
	MonitoringConfig *ContainerAttachedClusterMonitoringConfig `field:"optional" json:"monitoringConfig" yaml:"monitoringConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#project ContainerAttachedCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_attached_cluster#timeouts ContainerAttachedCluster#timeouts}
	Timeouts *ContainerAttachedClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

