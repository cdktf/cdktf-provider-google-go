// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouternat

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/computerouternat/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/compute_router_nat google_compute_router_nat}.
type ComputeRouterNat interface {
	cdktf.TerraformResource
	AutoNetworkTier() *string
	SetAutoNetworkTier(val *string)
	AutoNetworkTierInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DrainNatIps() *[]*string
	SetDrainNatIps(val *[]*string)
	DrainNatIpsInput() *[]*string
	EnableDynamicPortAllocation() interface{}
	SetEnableDynamicPortAllocation(val interface{})
	EnableDynamicPortAllocationInput() interface{}
	EnableEndpointIndependentMapping() interface{}
	SetEnableEndpointIndependentMapping(val interface{})
	EnableEndpointIndependentMappingInput() interface{}
	EndpointTypes() *[]*string
	SetEndpointTypes(val *[]*string)
	EndpointTypesInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IcmpIdleTimeoutSec() *float64
	SetIcmpIdleTimeoutSec(val *float64)
	IcmpIdleTimeoutSecInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitialNatIps() *[]*string
	SetInitialNatIps(val *[]*string)
	InitialNatIpsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogConfig() ComputeRouterNatLogConfigOutputReference
	LogConfigInput() *ComputeRouterNatLogConfig
	MaxPortsPerVm() *float64
	SetMaxPortsPerVm(val *float64)
	MaxPortsPerVmInput() *float64
	MinPortsPerVm() *float64
	SetMinPortsPerVm(val *float64)
	MinPortsPerVmInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nat64Subnetwork() ComputeRouterNatNat64SubnetworkList
	Nat64SubnetworkInput() interface{}
	NatIpAllocateOption() *string
	SetNatIpAllocateOption(val *string)
	NatIpAllocateOptionInput() *string
	NatIps() *[]*string
	SetNatIps(val *[]*string)
	NatIpsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	Router() *string
	SetRouter(val *string)
	RouterInput() *string
	Rules() ComputeRouterNatRulesList
	RulesInput() interface{}
	SourceSubnetworkIpRangesToNat() *string
	SetSourceSubnetworkIpRangesToNat(val *string)
	SourceSubnetworkIpRangesToNat64() *string
	SetSourceSubnetworkIpRangesToNat64(val *string)
	SourceSubnetworkIpRangesToNat64Input() *string
	SourceSubnetworkIpRangesToNatInput() *string
	Subnetwork() ComputeRouterNatSubnetworkList
	SubnetworkInput() interface{}
	TcpEstablishedIdleTimeoutSec() *float64
	SetTcpEstablishedIdleTimeoutSec(val *float64)
	TcpEstablishedIdleTimeoutSecInput() *float64
	TcpTimeWaitTimeoutSec() *float64
	SetTcpTimeWaitTimeoutSec(val *float64)
	TcpTimeWaitTimeoutSecInput() *float64
	TcpTransitoryIdleTimeoutSec() *float64
	SetTcpTransitoryIdleTimeoutSec(val *float64)
	TcpTransitoryIdleTimeoutSecInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeRouterNatTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UdpIdleTimeoutSec() *float64
	SetUdpIdleTimeoutSec(val *float64)
	UdpIdleTimeoutSecInput() *float64
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
	PutLogConfig(value *ComputeRouterNatLogConfig)
	PutNat64Subnetwork(value interface{})
	PutRules(value interface{})
	PutSubnetwork(value interface{})
	PutTimeouts(value *ComputeRouterNatTimeouts)
	ResetAutoNetworkTier()
	ResetDrainNatIps()
	ResetEnableDynamicPortAllocation()
	ResetEnableEndpointIndependentMapping()
	ResetEndpointTypes()
	ResetIcmpIdleTimeoutSec()
	ResetId()
	ResetInitialNatIps()
	ResetLogConfig()
	ResetMaxPortsPerVm()
	ResetMinPortsPerVm()
	ResetNat64Subnetwork()
	ResetNatIpAllocateOption()
	ResetNatIps()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetRules()
	ResetSourceSubnetworkIpRangesToNat64()
	ResetSubnetwork()
	ResetTcpEstablishedIdleTimeoutSec()
	ResetTcpTimeWaitTimeoutSec()
	ResetTcpTransitoryIdleTimeoutSec()
	ResetTimeouts()
	ResetType()
	ResetUdpIdleTimeoutSec()
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

// The jsii proxy struct for ComputeRouterNat
type jsiiProxy_ComputeRouterNat struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRouterNat) AutoNetworkTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoNetworkTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) AutoNetworkTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoNetworkTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) DrainNatIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"drainNatIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) DrainNatIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"drainNatIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) EnableDynamicPortAllocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDynamicPortAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) EnableDynamicPortAllocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDynamicPortAllocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) EnableEndpointIndependentMapping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEndpointIndependentMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) EnableEndpointIndependentMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEndpointIndependentMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) EndpointTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endpointTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) EndpointTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endpointTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) IcmpIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) IcmpIdleTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpIdleTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) InitialNatIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"initialNatIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) InitialNatIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"initialNatIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) LogConfig() ComputeRouterNatLogConfigOutputReference {
	var returns ComputeRouterNatLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) LogConfigInput() *ComputeRouterNatLogConfig {
	var returns *ComputeRouterNatLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) MaxPortsPerVm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPortsPerVm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) MaxPortsPerVmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPortsPerVmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) MinPortsPerVm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPortsPerVm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) MinPortsPerVmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPortsPerVmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Nat64Subnetwork() ComputeRouterNatNat64SubnetworkList {
	var returns ComputeRouterNatNat64SubnetworkList
	_jsii_.Get(
		j,
		"nat64Subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Nat64SubnetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nat64SubnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) NatIpAllocateOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAllocateOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) NatIpAllocateOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAllocateOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) NatIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) NatIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Router() *string {
	var returns *string
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) RouterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Rules() ComputeRouterNatRulesList {
	var returns ComputeRouterNatRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) SourceSubnetworkIpRangesToNat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSubnetworkIpRangesToNat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) SourceSubnetworkIpRangesToNat64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSubnetworkIpRangesToNat64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) SourceSubnetworkIpRangesToNat64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSubnetworkIpRangesToNat64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) SourceSubnetworkIpRangesToNatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSubnetworkIpRangesToNatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Subnetwork() ComputeRouterNatSubnetworkList {
	var returns ComputeRouterNatSubnetworkList
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) SubnetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TcpEstablishedIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TcpEstablishedIdleTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedIdleTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TcpTimeWaitTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTimeWaitTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TcpTimeWaitTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTimeWaitTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TcpTransitoryIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTransitoryIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TcpTransitoryIdleTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTransitoryIdleTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Timeouts() ComputeRouterNatTimeoutsOutputReference {
	var returns ComputeRouterNatTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) UdpIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterNat) UdpIdleTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpIdleTimeoutSecInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/compute_router_nat google_compute_router_nat} Resource.
func NewComputeRouterNat(scope constructs.Construct, id *string, config *ComputeRouterNatConfig) ComputeRouterNat {
	_init_.Initialize()

	if err := validateNewComputeRouterNatParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRouterNat{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRouterNat.ComputeRouterNat",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/compute_router_nat google_compute_router_nat} Resource.
func NewComputeRouterNat_Override(c ComputeRouterNat, scope constructs.Construct, id *string, config *ComputeRouterNatConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRouterNat.ComputeRouterNat",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetAutoNetworkTier(val *string) {
	if err := j.validateSetAutoNetworkTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoNetworkTier",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetDrainNatIps(val *[]*string) {
	if err := j.validateSetDrainNatIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainNatIps",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetEnableDynamicPortAllocation(val interface{}) {
	if err := j.validateSetEnableDynamicPortAllocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDynamicPortAllocation",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetEnableEndpointIndependentMapping(val interface{}) {
	if err := j.validateSetEnableEndpointIndependentMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEndpointIndependentMapping",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetEndpointTypes(val *[]*string) {
	if err := j.validateSetEndpointTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointTypes",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetIcmpIdleTimeoutSec(val *float64) {
	if err := j.validateSetIcmpIdleTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpIdleTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetInitialNatIps(val *[]*string) {
	if err := j.validateSetInitialNatIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialNatIps",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetMaxPortsPerVm(val *float64) {
	if err := j.validateSetMaxPortsPerVmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPortsPerVm",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetMinPortsPerVm(val *float64) {
	if err := j.validateSetMinPortsPerVmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPortsPerVm",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetNatIpAllocateOption(val *string) {
	if err := j.validateSetNatIpAllocateOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natIpAllocateOption",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetNatIps(val *[]*string) {
	if err := j.validateSetNatIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natIps",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetRouter(val *string) {
	if err := j.validateSetRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"router",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetSourceSubnetworkIpRangesToNat(val *string) {
	if err := j.validateSetSourceSubnetworkIpRangesToNatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSubnetworkIpRangesToNat",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetSourceSubnetworkIpRangesToNat64(val *string) {
	if err := j.validateSetSourceSubnetworkIpRangesToNat64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSubnetworkIpRangesToNat64",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetTcpEstablishedIdleTimeoutSec(val *float64) {
	if err := j.validateSetTcpEstablishedIdleTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpEstablishedIdleTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetTcpTimeWaitTimeoutSec(val *float64) {
	if err := j.validateSetTcpTimeWaitTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpTimeWaitTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetTcpTransitoryIdleTimeoutSec(val *float64) {
	if err := j.validateSetTcpTransitoryIdleTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpTransitoryIdleTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterNat)SetUdpIdleTimeoutSec(val *float64) {
	if err := j.validateSetUdpIdleTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpIdleTimeoutSec",
		val,
	)
}

// Generates CDKTF code for importing a ComputeRouterNat resource upon running "cdktf plan <stack-name>".
func ComputeRouterNat_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeRouterNat_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRouterNat.ComputeRouterNat",
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
func ComputeRouterNat_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRouterNat_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRouterNat.ComputeRouterNat",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRouterNat_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRouterNat_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRouterNat.ComputeRouterNat",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRouterNat_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRouterNat_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRouterNat.ComputeRouterNat",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeRouterNat_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRouterNat.ComputeRouterNat",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRouterNat) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeRouterNat) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeRouterNat) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRouterNat) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRouterNat) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRouterNat) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRouterNat) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRouterNat) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRouterNat) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRouterNat) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRouterNat) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRouterNat) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterNat) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeRouterNat) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRouterNat) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeRouterNat) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeRouterNat) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeRouterNat) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRouterNat) PutLogConfig(value *ComputeRouterNatLogConfig) {
	if err := c.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRouterNat) PutNat64Subnetwork(value interface{}) {
	if err := c.validatePutNat64SubnetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNat64Subnetwork",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRouterNat) PutRules(value interface{}) {
	if err := c.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRules",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRouterNat) PutSubnetwork(value interface{}) {
	if err := c.validatePutSubnetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSubnetwork",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRouterNat) PutTimeouts(value *ComputeRouterNatTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetAutoNetworkTier() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoNetworkTier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetDrainNatIps() {
	_jsii_.InvokeVoid(
		c,
		"resetDrainNatIps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetEnableDynamicPortAllocation() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableDynamicPortAllocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetEnableEndpointIndependentMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableEndpointIndependentMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetEndpointTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetEndpointTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetIcmpIdleTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetIcmpIdleTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetInitialNatIps() {
	_jsii_.InvokeVoid(
		c,
		"resetInitialNatIps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetLogConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetMaxPortsPerVm() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxPortsPerVm",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetMinPortsPerVm() {
	_jsii_.InvokeVoid(
		c,
		"resetMinPortsPerVm",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetNat64Subnetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNat64Subnetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetNatIpAllocateOption() {
	_jsii_.InvokeVoid(
		c,
		"resetNatIpAllocateOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetNatIps() {
	_jsii_.InvokeVoid(
		c,
		"resetNatIps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetRules() {
	_jsii_.InvokeVoid(
		c,
		"resetRules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetSourceSubnetworkIpRangesToNat64() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceSubnetworkIpRangesToNat64",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetTcpEstablishedIdleTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetTcpEstablishedIdleTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetTcpTimeWaitTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetTcpTimeWaitTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetTcpTransitoryIdleTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetTcpTransitoryIdleTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) ResetUdpIdleTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetUdpIdleTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterNat) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterNat) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterNat) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterNat) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterNat) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterNat) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

