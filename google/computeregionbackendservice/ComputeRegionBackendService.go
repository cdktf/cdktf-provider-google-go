// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/computeregionbackendservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/compute_region_backend_service google_compute_region_backend_service}.
type ComputeRegionBackendService interface {
	cdktf.TerraformResource
	AffinityCookieTtlSec() *float64
	SetAffinityCookieTtlSec(val *float64)
	AffinityCookieTtlSecInput() *float64
	Backend() ComputeRegionBackendServiceBackendList
	BackendInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnPolicy() ComputeRegionBackendServiceCdnPolicyOutputReference
	CdnPolicyInput() *ComputeRegionBackendServiceCdnPolicy
	CircuitBreakers() ComputeRegionBackendServiceCircuitBreakersOutputReference
	CircuitBreakersInput() *ComputeRegionBackendServiceCircuitBreakers
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionDrainingTimeoutSec() *float64
	SetConnectionDrainingTimeoutSec(val *float64)
	ConnectionDrainingTimeoutSecInput() *float64
	ConsistentHash() ComputeRegionBackendServiceConsistentHashOutputReference
	ConsistentHashInput() *ComputeRegionBackendServiceConsistentHash
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnableCdn() interface{}
	SetEnableCdn(val interface{})
	EnableCdnInput() interface{}
	FailoverPolicy() ComputeRegionBackendServiceFailoverPolicyOutputReference
	FailoverPolicyInput() *ComputeRegionBackendServiceFailoverPolicy
	Fingerprint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthChecks() *[]*string
	SetHealthChecks(val *[]*string)
	HealthChecksInput() *[]*string
	Iap() ComputeRegionBackendServiceIapOutputReference
	IapInput() *ComputeRegionBackendServiceIap
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancingScheme() *string
	SetLoadBalancingScheme(val *string)
	LoadBalancingSchemeInput() *string
	LocalityLbPolicy() *string
	SetLocalityLbPolicy(val *string)
	LocalityLbPolicyInput() *string
	LogConfig() ComputeRegionBackendServiceLogConfigOutputReference
	LogConfigInput() *ComputeRegionBackendServiceLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	OutlierDetection() ComputeRegionBackendServiceOutlierDetectionOutputReference
	OutlierDetectionInput() *ComputeRegionBackendServiceOutlierDetection
	PortName() *string
	SetPortName(val *string)
	PortNameInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SelfLink() *string
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeRegionBackendServiceTimeoutsOutputReference
	TimeoutSec() *float64
	SetTimeoutSec(val *float64)
	TimeoutSecInput() *float64
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBackend(value interface{})
	PutCdnPolicy(value *ComputeRegionBackendServiceCdnPolicy)
	PutCircuitBreakers(value *ComputeRegionBackendServiceCircuitBreakers)
	PutConsistentHash(value *ComputeRegionBackendServiceConsistentHash)
	PutFailoverPolicy(value *ComputeRegionBackendServiceFailoverPolicy)
	PutIap(value *ComputeRegionBackendServiceIap)
	PutLogConfig(value *ComputeRegionBackendServiceLogConfig)
	PutOutlierDetection(value *ComputeRegionBackendServiceOutlierDetection)
	PutTimeouts(value *ComputeRegionBackendServiceTimeouts)
	ResetAffinityCookieTtlSec()
	ResetBackend()
	ResetCdnPolicy()
	ResetCircuitBreakers()
	ResetConnectionDrainingTimeoutSec()
	ResetConsistentHash()
	ResetDescription()
	ResetEnableCdn()
	ResetFailoverPolicy()
	ResetHealthChecks()
	ResetIap()
	ResetId()
	ResetLoadBalancingScheme()
	ResetLocalityLbPolicy()
	ResetLogConfig()
	ResetNetwork()
	ResetOutlierDetection()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortName()
	ResetProject()
	ResetProtocol()
	ResetRegion()
	ResetSessionAffinity()
	ResetTimeouts()
	ResetTimeoutSec()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeRegionBackendService
type jsiiProxy_ComputeRegionBackendService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRegionBackendService) AffinityCookieTtlSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"affinityCookieTtlSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) AffinityCookieTtlSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"affinityCookieTtlSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Backend() ComputeRegionBackendServiceBackendList {
	var returns ComputeRegionBackendServiceBackendList
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) BackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) CdnPolicy() ComputeRegionBackendServiceCdnPolicyOutputReference {
	var returns ComputeRegionBackendServiceCdnPolicyOutputReference
	_jsii_.Get(
		j,
		"cdnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) CdnPolicyInput() *ComputeRegionBackendServiceCdnPolicy {
	var returns *ComputeRegionBackendServiceCdnPolicy
	_jsii_.Get(
		j,
		"cdnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) CircuitBreakers() ComputeRegionBackendServiceCircuitBreakersOutputReference {
	var returns ComputeRegionBackendServiceCircuitBreakersOutputReference
	_jsii_.Get(
		j,
		"circuitBreakers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) CircuitBreakersInput() *ComputeRegionBackendServiceCircuitBreakers {
	var returns *ComputeRegionBackendServiceCircuitBreakers
	_jsii_.Get(
		j,
		"circuitBreakersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) ConnectionDrainingTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) ConnectionDrainingTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) ConsistentHash() ComputeRegionBackendServiceConsistentHashOutputReference {
	var returns ComputeRegionBackendServiceConsistentHashOutputReference
	_jsii_.Get(
		j,
		"consistentHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) ConsistentHashInput() *ComputeRegionBackendServiceConsistentHash {
	var returns *ComputeRegionBackendServiceConsistentHash
	_jsii_.Get(
		j,
		"consistentHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) EnableCdn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) EnableCdnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) FailoverPolicy() ComputeRegionBackendServiceFailoverPolicyOutputReference {
	var returns ComputeRegionBackendServiceFailoverPolicyOutputReference
	_jsii_.Get(
		j,
		"failoverPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) FailoverPolicyInput() *ComputeRegionBackendServiceFailoverPolicy {
	var returns *ComputeRegionBackendServiceFailoverPolicy
	_jsii_.Get(
		j,
		"failoverPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) HealthChecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) HealthChecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Iap() ComputeRegionBackendServiceIapOutputReference {
	var returns ComputeRegionBackendServiceIapOutputReference
	_jsii_.Get(
		j,
		"iap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) IapInput() *ComputeRegionBackendServiceIap {
	var returns *ComputeRegionBackendServiceIap
	_jsii_.Get(
		j,
		"iapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) LoadBalancingScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) LoadBalancingSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) LocalityLbPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityLbPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) LocalityLbPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityLbPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) LogConfig() ComputeRegionBackendServiceLogConfigOutputReference {
	var returns ComputeRegionBackendServiceLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) LogConfigInput() *ComputeRegionBackendServiceLogConfig {
	var returns *ComputeRegionBackendServiceLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) OutlierDetection() ComputeRegionBackendServiceOutlierDetectionOutputReference {
	var returns ComputeRegionBackendServiceOutlierDetectionOutputReference
	_jsii_.Get(
		j,
		"outlierDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) OutlierDetectionInput() *ComputeRegionBackendServiceOutlierDetection {
	var returns *ComputeRegionBackendServiceOutlierDetection
	_jsii_.Get(
		j,
		"outlierDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) Timeouts() ComputeRegionBackendServiceTimeoutsOutputReference {
	var returns ComputeRegionBackendServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) TimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) TimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/compute_region_backend_service google_compute_region_backend_service} Resource.
func NewComputeRegionBackendService(scope constructs.Construct, id *string, config *ComputeRegionBackendServiceConfig) ComputeRegionBackendService {
	_init_.Initialize()

	if err := validateNewComputeRegionBackendServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionBackendService{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/compute_region_backend_service google_compute_region_backend_service} Resource.
func NewComputeRegionBackendService_Override(c ComputeRegionBackendService, scope constructs.Construct, id *string, config *ComputeRegionBackendServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendService",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetAffinityCookieTtlSec(val *float64) {
	if err := j.validateSetAffinityCookieTtlSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affinityCookieTtlSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetConnectionDrainingTimeoutSec(val *float64) {
	if err := j.validateSetConnectionDrainingTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionDrainingTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetEnableCdn(val interface{}) {
	if err := j.validateSetEnableCdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCdn",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetHealthChecks(val *[]*string) {
	if err := j.validateSetHealthChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthChecks",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetLoadBalancingScheme(val *string) {
	if err := j.validateSetLoadBalancingSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingScheme",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetLocalityLbPolicy(val *string) {
	if err := j.validateSetLocalityLbPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localityLbPolicy",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendService)SetTimeoutSec(val *float64) {
	if err := j.validateSetTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSec",
		val,
	)
}

// Generates CDKTF code for importing a ComputeRegionBackendService resource upon running "cdktf plan <stack-name>".
func ComputeRegionBackendService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeRegionBackendService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendService",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ComputeRegionBackendService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionBackendService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionBackendService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionBackendService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionBackendService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionBackendService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeRegionBackendService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutBackend(value interface{}) {
	if err := c.validatePutBackendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBackend",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutCdnPolicy(value *ComputeRegionBackendServiceCdnPolicy) {
	if err := c.validatePutCdnPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCdnPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutCircuitBreakers(value *ComputeRegionBackendServiceCircuitBreakers) {
	if err := c.validatePutCircuitBreakersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCircuitBreakers",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutConsistentHash(value *ComputeRegionBackendServiceConsistentHash) {
	if err := c.validatePutConsistentHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConsistentHash",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutFailoverPolicy(value *ComputeRegionBackendServiceFailoverPolicy) {
	if err := c.validatePutFailoverPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFailoverPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutIap(value *ComputeRegionBackendServiceIap) {
	if err := c.validatePutIapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIap",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutLogConfig(value *ComputeRegionBackendServiceLogConfig) {
	if err := c.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutOutlierDetection(value *ComputeRegionBackendServiceOutlierDetection) {
	if err := c.validatePutOutlierDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOutlierDetection",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) PutTimeouts(value *ComputeRegionBackendServiceTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetAffinityCookieTtlSec() {
	_jsii_.InvokeVoid(
		c,
		"resetAffinityCookieTtlSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetBackend() {
	_jsii_.InvokeVoid(
		c,
		"resetBackend",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetCdnPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCdnPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetCircuitBreakers() {
	_jsii_.InvokeVoid(
		c,
		"resetCircuitBreakers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetConnectionDrainingTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectionDrainingTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetConsistentHash() {
	_jsii_.InvokeVoid(
		c,
		"resetConsistentHash",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetEnableCdn() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableCdn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetFailoverPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetFailoverPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetHealthChecks() {
	_jsii_.InvokeVoid(
		c,
		"resetHealthChecks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetIap() {
	_jsii_.InvokeVoid(
		c,
		"resetIap",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetLoadBalancingScheme() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancingScheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetLocalityLbPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalityLbPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetLogConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetOutlierDetection() {
	_jsii_.InvokeVoid(
		c,
		"resetOutlierDetection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetPortName() {
	_jsii_.InvokeVoid(
		c,
		"resetPortName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		c,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) ResetTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

