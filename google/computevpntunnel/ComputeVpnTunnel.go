package computevpntunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v8/computevpntunnel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.75.0/docs/resources/compute_vpn_tunnel google_compute_vpn_tunnel}.
type ComputeVpnTunnel interface {
	cdktf.TerraformResource
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DetailedStatus() *string
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
	IkeVersion() *float64
	SetIkeVersion(val *float64)
	IkeVersionInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalTrafficSelector() *[]*string
	SetLocalTrafficSelector(val *[]*string)
	LocalTrafficSelectorInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerExternalGateway() *string
	SetPeerExternalGateway(val *string)
	PeerExternalGatewayInput() *string
	PeerExternalGatewayInterface() *float64
	SetPeerExternalGatewayInterface(val *float64)
	PeerExternalGatewayInterfaceInput() *float64
	PeerGcpGateway() *string
	SetPeerGcpGateway(val *string)
	PeerGcpGatewayInput() *string
	PeerIp() *string
	SetPeerIp(val *string)
	PeerIpInput() *string
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
	RemoteTrafficSelector() *[]*string
	SetRemoteTrafficSelector(val *[]*string)
	RemoteTrafficSelectorInput() *[]*string
	Router() *string
	SetRouter(val *string)
	RouterInput() *string
	SelfLink() *string
	SharedSecret() *string
	SetSharedSecret(val *string)
	SharedSecretHash() *string
	SharedSecretInput() *string
	TargetVpnGateway() *string
	SetTargetVpnGateway(val *string)
	TargetVpnGatewayInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeVpnTunnelTimeoutsOutputReference
	TimeoutsInput() interface{}
	TunnelId() *string
	VpnGateway() *string
	SetVpnGateway(val *string)
	VpnGatewayInput() *string
	VpnGatewayInterface() *float64
	SetVpnGatewayInterface(val *float64)
	VpnGatewayInterfaceInput() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *ComputeVpnTunnelTimeouts)
	ResetDescription()
	ResetId()
	ResetIkeVersion()
	ResetLocalTrafficSelector()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerExternalGateway()
	ResetPeerExternalGatewayInterface()
	ResetPeerGcpGateway()
	ResetPeerIp()
	ResetProject()
	ResetRegion()
	ResetRemoteTrafficSelector()
	ResetRouter()
	ResetTargetVpnGateway()
	ResetTimeouts()
	ResetVpnGateway()
	ResetVpnGatewayInterface()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeVpnTunnel
type jsiiProxy_ComputeVpnTunnel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeVpnTunnel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) DetailedStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detailedStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) IkeVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) IkeVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) LocalTrafficSelector() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localTrafficSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) LocalTrafficSelectorInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localTrafficSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) PeerExternalGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerExternalGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) PeerExternalGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerExternalGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) PeerExternalGatewayInterface() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerExternalGatewayInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) PeerExternalGatewayInterfaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerExternalGatewayInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) PeerGcpGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerGcpGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) PeerGcpGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerGcpGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) PeerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) PeerIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) RemoteTrafficSelector() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteTrafficSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) RemoteTrafficSelectorInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteTrafficSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Router() *string {
	var returns *string
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) RouterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) SharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) SharedSecretHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecretHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) SharedSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) TargetVpnGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVpnGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) TargetVpnGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVpnGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) Timeouts() ComputeVpnTunnelTimeoutsOutputReference {
	var returns ComputeVpnTunnelTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) TunnelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) VpnGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) VpnGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) VpnGatewayInterface() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnGatewayInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnel) VpnGatewayInterfaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnGatewayInterfaceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.75.0/docs/resources/compute_vpn_tunnel google_compute_vpn_tunnel} Resource.
func NewComputeVpnTunnel(scope constructs.Construct, id *string, config *ComputeVpnTunnelConfig) ComputeVpnTunnel {
	_init_.Initialize()

	if err := validateNewComputeVpnTunnelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeVpnTunnel{}

	_jsii_.Create(
		"@cdktf/provider-google.computeVpnTunnel.ComputeVpnTunnel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.75.0/docs/resources/compute_vpn_tunnel google_compute_vpn_tunnel} Resource.
func NewComputeVpnTunnel_Override(c ComputeVpnTunnel, scope constructs.Construct, id *string, config *ComputeVpnTunnelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeVpnTunnel.ComputeVpnTunnel",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetIkeVersion(val *float64) {
	if err := j.validateSetIkeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeVersion",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetLocalTrafficSelector(val *[]*string) {
	if err := j.validateSetLocalTrafficSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localTrafficSelector",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetPeerExternalGateway(val *string) {
	if err := j.validateSetPeerExternalGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerExternalGateway",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetPeerExternalGatewayInterface(val *float64) {
	if err := j.validateSetPeerExternalGatewayInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerExternalGatewayInterface",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetPeerGcpGateway(val *string) {
	if err := j.validateSetPeerGcpGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerGcpGateway",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetPeerIp(val *string) {
	if err := j.validateSetPeerIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIp",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetRemoteTrafficSelector(val *[]*string) {
	if err := j.validateSetRemoteTrafficSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteTrafficSelector",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetRouter(val *string) {
	if err := j.validateSetRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"router",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetSharedSecret(val *string) {
	if err := j.validateSetSharedSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedSecret",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetTargetVpnGateway(val *string) {
	if err := j.validateSetTargetVpnGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVpnGateway",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetVpnGateway(val *string) {
	if err := j.validateSetVpnGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGateway",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnel)SetVpnGatewayInterface(val *float64) {
	if err := j.validateSetVpnGatewayInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGatewayInterface",
		val,
	)
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
func ComputeVpnTunnel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeVpnTunnel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeVpnTunnel.ComputeVpnTunnel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeVpnTunnel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeVpnTunnel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeVpnTunnel.ComputeVpnTunnel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeVpnTunnel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeVpnTunnel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeVpnTunnel.ComputeVpnTunnel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeVpnTunnel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeVpnTunnel.ComputeVpnTunnel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeVpnTunnel) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeVpnTunnel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeVpnTunnel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeVpnTunnel) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeVpnTunnel) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeVpnTunnel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeVpnTunnel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeVpnTunnel) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeVpnTunnel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeVpnTunnel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeVpnTunnel) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) PutTimeouts(value *ComputeVpnTunnelTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetIkeVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetIkeVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetLocalTrafficSelector() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalTrafficSelector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetPeerExternalGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetPeerExternalGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetPeerExternalGatewayInterface() {
	_jsii_.InvokeVoid(
		c,
		"resetPeerExternalGatewayInterface",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetPeerGcpGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetPeerGcpGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetPeerIp() {
	_jsii_.InvokeVoid(
		c,
		"resetPeerIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetRemoteTrafficSelector() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteTrafficSelector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetRouter() {
	_jsii_.InvokeVoid(
		c,
		"resetRouter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetTargetVpnGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetVpnGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetVpnGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetVpnGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) ResetVpnGatewayInterface() {
	_jsii_.InvokeVoid(
		c,
		"resetVpnGatewayInterface",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

