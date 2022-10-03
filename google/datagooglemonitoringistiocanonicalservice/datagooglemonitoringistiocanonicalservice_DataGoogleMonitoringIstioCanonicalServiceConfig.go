package datagooglemonitoringistiocanonicalservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleMonitoringIstioCanonicalServiceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of the canonical service underlying this service..                       Corresponds to the destination_service_name metric label in Istio metrics.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/monitoring_istio_canonical_service#canonical_service DataGoogleMonitoringIstioCanonicalService#canonical_service}
	CanonicalService *string `field:"required" json:"canonicalService" yaml:"canonicalService"`
	// The namespace of the canonical service underlying this service.
	//
	// Corresponds to the destination_service_namespace metric label in Istio metrics.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/monitoring_istio_canonical_service#canonical_service_namespace DataGoogleMonitoringIstioCanonicalService#canonical_service_namespace}
	CanonicalServiceNamespace *string `field:"required" json:"canonicalServiceNamespace" yaml:"canonicalServiceNamespace"`
	// Identifier for the Istio mesh in which this canonical service is defined.
	//
	// Corresponds to the meshUid metric label in Istio metrics.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/monitoring_istio_canonical_service#mesh_uid DataGoogleMonitoringIstioCanonicalService#mesh_uid}
	MeshUid *string `field:"required" json:"meshUid" yaml:"meshUid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/monitoring_istio_canonical_service#id DataGoogleMonitoringIstioCanonicalService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/monitoring_istio_canonical_service#project DataGoogleMonitoringIstioCanonicalService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

