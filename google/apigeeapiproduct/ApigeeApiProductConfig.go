// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeapiproduct

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeApiProductConfig struct {
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
	// Name displayed in the UI or developer portal to developers registering for API access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#display_name ApigeeApiProduct#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Internal name of the API product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#name ApigeeApiProduct#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Apigee Organization associated with the Apigee API product, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#org_id ApigeeApiProduct#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Comma-separated list of API resources to be bundled in the API product.
	//
	// By default, the resource paths are mapped from the proxy.pathsuffix variable.
	// The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the apiResources element is defined to be /forecastrss and the base path defined for the API proxy is /weather, then only requests to /weather/forecastrss are permitted by the API product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#api_resources ApigeeApiProduct#api_resources}
	ApiResources *[]*string `field:"optional" json:"apiResources" yaml:"apiResources"`
	// Flag that specifies how API keys are approved to access the APIs defined by the API product.
	//
	// Valid values are 'auto' or 'manual'. Possible values: ["auto", "manual"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#approval_type ApigeeApiProduct#approval_type}
	ApprovalType *string `field:"optional" json:"approvalType" yaml:"approvalType"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#attributes ApigeeApiProduct#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// Description of the API product. Include key information about the API product that is not captured by other fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#description ApigeeApiProduct#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Comma-separated list of environment names to which the API product is bound.
	//
	// Requests to environments that are not listed are rejected.
	// By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#environments ApigeeApiProduct#environments}
	Environments *[]*string `field:"optional" json:"environments" yaml:"environments"`
	// graphql_operation_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#graphql_operation_group ApigeeApiProduct#graphql_operation_group}
	GraphqlOperationGroup *ApigeeApiProductGraphqlOperationGroup `field:"optional" json:"graphqlOperationGroup" yaml:"graphqlOperationGroup"`
	// grpc_operation_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#grpc_operation_group ApigeeApiProduct#grpc_operation_group}
	GrpcOperationGroup *ApigeeApiProductGrpcOperationGroup `field:"optional" json:"grpcOperationGroup" yaml:"grpcOperationGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#id ApigeeApiProduct#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// operation_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#operation_group ApigeeApiProduct#operation_group}
	OperationGroup *ApigeeApiProductOperationGroup `field:"optional" json:"operationGroup" yaml:"operationGroup"`
	// Comma-separated list of API proxy names to which this API product is bound.
	//
	// By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies.
	// Apigee rejects requests to API proxies that are not listed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#proxies ApigeeApiProduct#proxies}
	Proxies *[]*string `field:"optional" json:"proxies" yaml:"proxies"`
	// Number of request messages permitted per app by this API product for the specified quotaInterval and quotaTimeUnit.
	//
	// For example, a quota of 50, for a quotaInterval of 12 and a quotaTimeUnit of hours means 50 requests are allowed every 12 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#quota ApigeeApiProduct#quota}
	Quota *string `field:"optional" json:"quota" yaml:"quota"`
	// Scope of the quota decides how the quota counter gets applied and evaluate for quota violation.
	//
	// If the Scope is set as PROXY, then all the operations defined for the APIproduct that are associated with the same proxy will share the same quota counter set at the APIproduct level, making it a global counter at a proxy level. If the Scope is set as OPERATION, then each operations get the counter set at the API product dedicated, making it a local counter. Note that, the QuotaCounterScope applies only when an operation does not have dedicated quota set for itself. Possible values: ["QUOTA_COUNTER_SCOPE_UNSPECIFIED", "PROXY", "OPERATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#quota_counter_scope ApigeeApiProduct#quota_counter_scope}
	QuotaCounterScope *string `field:"optional" json:"quotaCounterScope" yaml:"quotaCounterScope"`
	// Time interval over which the number of request messages is calculated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#quota_interval ApigeeApiProduct#quota_interval}
	QuotaInterval *string `field:"optional" json:"quotaInterval" yaml:"quotaInterval"`
	// Time unit defined for the quotaInterval. Valid values include second, minute, hour, day, month or year.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#quota_time_unit ApigeeApiProduct#quota_time_unit}
	QuotaTimeUnit *string `field:"optional" json:"quotaTimeUnit" yaml:"quotaTimeUnit"`
	// Comma-separated list of OAuth scopes that are validated at runtime.
	//
	// Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#scopes ApigeeApiProduct#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Optional. The resource ID of the parent Space. If not set, the parent resource will be the Organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#space ApigeeApiProduct#space}
	Space *string `field:"optional" json:"space" yaml:"space"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apigee_api_product#timeouts ApigeeApiProduct#timeouts}
	Timeouts *ApigeeApiProductTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

