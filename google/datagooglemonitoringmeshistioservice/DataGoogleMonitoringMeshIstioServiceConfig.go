package datagooglemonitoringmeshistioservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleMonitoringMeshIstioServiceConfig struct {
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
	// Identifier for the mesh in which this Istio service is defined.
	//
	// Corresponds to the meshUid metric label in Istio metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/monitoring_mesh_istio_service#mesh_uid DataGoogleMonitoringMeshIstioService#mesh_uid}
	MeshUid *string `field:"required" json:"meshUid" yaml:"meshUid"`
	// The name of the Istio service underlying this service.
	//
	// Corresponds to the destination_service_name metric label in Istio metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/monitoring_mesh_istio_service#service_name DataGoogleMonitoringMeshIstioService#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The namespace of the Istio service underlying this service.
	//
	// Corresponds to the destination_service_namespace metric label in Istio metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/monitoring_mesh_istio_service#service_namespace DataGoogleMonitoringMeshIstioService#service_namespace}
	ServiceNamespace *string `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/monitoring_mesh_istio_service#id DataGoogleMonitoringMeshIstioService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/data-sources/monitoring_mesh_istio_service#project DataGoogleMonitoringMeshIstioService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

