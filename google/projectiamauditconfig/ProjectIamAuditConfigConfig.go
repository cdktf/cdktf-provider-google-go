package projectiamauditconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIamAuditConfigConfig struct {
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
	// audit_log_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_audit_config#audit_log_config ProjectIamAuditConfig#audit_log_config}
	AuditLogConfig interface{} `field:"required" json:"auditLogConfig" yaml:"auditLogConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_audit_config#project ProjectIamAuditConfig#project}.
	Project *string `field:"required" json:"project" yaml:"project"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_audit_config#service ProjectIamAuditConfig#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_iam_audit_config#id ProjectIamAuditConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

