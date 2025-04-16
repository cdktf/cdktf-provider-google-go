// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagementvpcflowlogsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/networkmanagementvpcflowlogsconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/network_management_vpc_flow_logs_config google_network_management_vpc_flow_logs_config}.
type NetworkManagementVpcFlowLogsConfig interface {
	cdktf.TerraformResource
	AggregationInterval() *string
	SetAggregationInterval(val *string)
	AggregationIntervalInput() *string
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	FilterExpr() *string
	SetFilterExpr(val *string)
	FilterExprInput() *string
	FlowSampling() *float64
	SetFlowSampling(val *float64)
	FlowSamplingInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InterconnectAttachment() *string
	SetInterconnectAttachment(val *string)
	InterconnectAttachmentInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Metadata() *string
	SetMetadata(val *string)
	MetadataFields() *[]*string
	SetMetadataFields(val *[]*string)
	MetadataFieldsInput() *[]*string
	MetadataInput() *string
	Name() *string
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
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkManagementVpcFlowLogsConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	VpcFlowLogsConfigId() *string
	SetVpcFlowLogsConfigId(val *string)
	VpcFlowLogsConfigIdInput() *string
	VpnTunnel() *string
	SetVpnTunnel(val *string)
	VpnTunnelInput() *string
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
	PutTimeouts(value *NetworkManagementVpcFlowLogsConfigTimeouts)
	ResetAggregationInterval()
	ResetDescription()
	ResetFilterExpr()
	ResetFlowSampling()
	ResetId()
	ResetInterconnectAttachment()
	ResetLabels()
	ResetMetadata()
	ResetMetadataFields()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetState()
	ResetTimeouts()
	ResetVpnTunnel()
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

// The jsii proxy struct for NetworkManagementVpcFlowLogsConfig
type jsiiProxy_NetworkManagementVpcFlowLogsConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) AggregationInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) AggregationIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) FilterExpr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExpr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) FilterExprInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExprInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) FlowSampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowSampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) FlowSamplingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowSamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) InterconnectAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) InterconnectAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) MetadataFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) MetadataFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) Timeouts() NetworkManagementVpcFlowLogsConfigTimeoutsOutputReference {
	var returns NetworkManagementVpcFlowLogsConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) VpcFlowLogsConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogsConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) VpcFlowLogsConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogsConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) VpnTunnel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig) VpnTunnelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnTunnelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/network_management_vpc_flow_logs_config google_network_management_vpc_flow_logs_config} Resource.
func NewNetworkManagementVpcFlowLogsConfig(scope constructs.Construct, id *string, config *NetworkManagementVpcFlowLogsConfigConfig) NetworkManagementVpcFlowLogsConfig {
	_init_.Initialize()

	if err := validateNewNetworkManagementVpcFlowLogsConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkManagementVpcFlowLogsConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.networkManagementVpcFlowLogsConfig.NetworkManagementVpcFlowLogsConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/network_management_vpc_flow_logs_config google_network_management_vpc_flow_logs_config} Resource.
func NewNetworkManagementVpcFlowLogsConfig_Override(n NetworkManagementVpcFlowLogsConfig, scope constructs.Construct, id *string, config *NetworkManagementVpcFlowLogsConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkManagementVpcFlowLogsConfig.NetworkManagementVpcFlowLogsConfig",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetAggregationInterval(val *string) {
	if err := j.validateSetAggregationIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregationInterval",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetFilterExpr(val *string) {
	if err := j.validateSetFilterExprParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterExpr",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetFlowSampling(val *float64) {
	if err := j.validateSetFlowSamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowSampling",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetInterconnectAttachment(val *string) {
	if err := j.validateSetInterconnectAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interconnectAttachment",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetMetadata(val *string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetMetadataFields(val *[]*string) {
	if err := j.validateSetMetadataFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataFields",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetVpcFlowLogsConfigId(val *string) {
	if err := j.validateSetVpcFlowLogsConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcFlowLogsConfigId",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementVpcFlowLogsConfig)SetVpnTunnel(val *string) {
	if err := j.validateSetVpnTunnelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnTunnel",
		val,
	)
}

// Generates CDKTF code for importing a NetworkManagementVpcFlowLogsConfig resource upon running "cdktf plan <stack-name>".
func NetworkManagementVpcFlowLogsConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkManagementVpcFlowLogsConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkManagementVpcFlowLogsConfig.NetworkManagementVpcFlowLogsConfig",
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
func NetworkManagementVpcFlowLogsConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkManagementVpcFlowLogsConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkManagementVpcFlowLogsConfig.NetworkManagementVpcFlowLogsConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkManagementVpcFlowLogsConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkManagementVpcFlowLogsConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkManagementVpcFlowLogsConfig.NetworkManagementVpcFlowLogsConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkManagementVpcFlowLogsConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkManagementVpcFlowLogsConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkManagementVpcFlowLogsConfig.NetworkManagementVpcFlowLogsConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkManagementVpcFlowLogsConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.networkManagementVpcFlowLogsConfig.NetworkManagementVpcFlowLogsConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) PutTimeouts(value *NetworkManagementVpcFlowLogsConfigTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetAggregationInterval() {
	_jsii_.InvokeVoid(
		n,
		"resetAggregationInterval",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetFilterExpr() {
	_jsii_.InvokeVoid(
		n,
		"resetFilterExpr",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetFlowSampling() {
	_jsii_.InvokeVoid(
		n,
		"resetFlowSampling",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetInterconnectAttachment() {
	_jsii_.InvokeVoid(
		n,
		"resetInterconnectAttachment",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetMetadata() {
	_jsii_.InvokeVoid(
		n,
		"resetMetadata",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetMetadataFields() {
	_jsii_.InvokeVoid(
		n,
		"resetMetadataFields",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetState() {
	_jsii_.InvokeVoid(
		n,
		"resetState",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ResetVpnTunnel() {
	_jsii_.InvokeVoid(
		n,
		"resetVpnTunnel",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementVpcFlowLogsConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

