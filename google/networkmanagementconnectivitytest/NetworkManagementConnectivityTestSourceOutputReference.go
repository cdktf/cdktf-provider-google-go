// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagementconnectivitytest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/networkmanagementconnectivitytest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkManagementConnectivityTestSourceOutputReference interface {
	cdktf.ComplexObject
	AppEngineVersion() NetworkManagementConnectivityTestSourceAppEngineVersionOutputReference
	AppEngineVersionInput() *NetworkManagementConnectivityTestSourceAppEngineVersion
	CloudFunction() NetworkManagementConnectivityTestSourceCloudFunctionOutputReference
	CloudFunctionInput() *NetworkManagementConnectivityTestSourceCloudFunction
	CloudRunRevision() NetworkManagementConnectivityTestSourceCloudRunRevisionOutputReference
	CloudRunRevisionInput() *NetworkManagementConnectivityTestSourceCloudRunRevision
	CloudSqlInstance() *string
	SetCloudSqlInstance(val *string)
	CloudSqlInstanceInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GkeMasterCluster() *string
	SetGkeMasterCluster(val *string)
	GkeMasterClusterInput() *string
	Instance() *string
	SetInstance(val *string)
	InstanceInput() *string
	InternalValue() *NetworkManagementConnectivityTestSource
	SetInternalValue(val *NetworkManagementConnectivityTestSource)
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAppEngineVersion(value *NetworkManagementConnectivityTestSourceAppEngineVersion)
	PutCloudFunction(value *NetworkManagementConnectivityTestSourceCloudFunction)
	PutCloudRunRevision(value *NetworkManagementConnectivityTestSourceCloudRunRevision)
	ResetAppEngineVersion()
	ResetCloudFunction()
	ResetCloudRunRevision()
	ResetCloudSqlInstance()
	ResetGkeMasterCluster()
	ResetInstance()
	ResetIpAddress()
	ResetNetwork()
	ResetNetworkType()
	ResetPort()
	ResetProjectId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkManagementConnectivityTestSourceOutputReference
type jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) AppEngineVersion() NetworkManagementConnectivityTestSourceAppEngineVersionOutputReference {
	var returns NetworkManagementConnectivityTestSourceAppEngineVersionOutputReference
	_jsii_.Get(
		j,
		"appEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) AppEngineVersionInput() *NetworkManagementConnectivityTestSourceAppEngineVersion {
	var returns *NetworkManagementConnectivityTestSourceAppEngineVersion
	_jsii_.Get(
		j,
		"appEngineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) CloudFunction() NetworkManagementConnectivityTestSourceCloudFunctionOutputReference {
	var returns NetworkManagementConnectivityTestSourceCloudFunctionOutputReference
	_jsii_.Get(
		j,
		"cloudFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) CloudFunctionInput() *NetworkManagementConnectivityTestSourceCloudFunction {
	var returns *NetworkManagementConnectivityTestSourceCloudFunction
	_jsii_.Get(
		j,
		"cloudFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) CloudRunRevision() NetworkManagementConnectivityTestSourceCloudRunRevisionOutputReference {
	var returns NetworkManagementConnectivityTestSourceCloudRunRevisionOutputReference
	_jsii_.Get(
		j,
		"cloudRunRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) CloudRunRevisionInput() *NetworkManagementConnectivityTestSourceCloudRunRevision {
	var returns *NetworkManagementConnectivityTestSourceCloudRunRevision
	_jsii_.Get(
		j,
		"cloudRunRevisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) CloudSqlInstance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSqlInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) CloudSqlInstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSqlInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GkeMasterCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeMasterCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GkeMasterClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeMasterClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) Instance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) InstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) InternalValue() *NetworkManagementConnectivityTestSource {
	var returns *NetworkManagementConnectivityTestSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkManagementConnectivityTestSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkManagementConnectivityTestSourceOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkManagementConnectivityTestSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkManagementConnectivityTest.NetworkManagementConnectivityTestSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkManagementConnectivityTestSourceOutputReference_Override(n NetworkManagementConnectivityTestSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkManagementConnectivityTest.NetworkManagementConnectivityTestSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetCloudSqlInstance(val *string) {
	if err := j.validateSetCloudSqlInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudSqlInstance",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetGkeMasterCluster(val *string) {
	if err := j.validateSetGkeMasterClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gkeMasterCluster",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetInstance(val *string) {
	if err := j.validateSetInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instance",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetInternalValue(val *NetworkManagementConnectivityTestSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) PutAppEngineVersion(value *NetworkManagementConnectivityTestSourceAppEngineVersion) {
	if err := n.validatePutAppEngineVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAppEngineVersion",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) PutCloudFunction(value *NetworkManagementConnectivityTestSourceCloudFunction) {
	if err := n.validatePutCloudFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCloudFunction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) PutCloudRunRevision(value *NetworkManagementConnectivityTestSourceCloudRunRevision) {
	if err := n.validatePutCloudRunRevisionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCloudRunRevision",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetAppEngineVersion() {
	_jsii_.InvokeVoid(
		n,
		"resetAppEngineVersion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetCloudFunction() {
	_jsii_.InvokeVoid(
		n,
		"resetCloudFunction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetCloudRunRevision() {
	_jsii_.InvokeVoid(
		n,
		"resetCloudRunRevision",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetCloudSqlInstance() {
	_jsii_.InvokeVoid(
		n,
		"resetCloudSqlInstance",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetGkeMasterCluster() {
	_jsii_.InvokeVoid(
		n,
		"resetGkeMasterCluster",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetInstance() {
	_jsii_.InvokeVoid(
		n,
		"resetInstance",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		n,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		n,
		"resetNetwork",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetNetworkType() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		n,
		"resetPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ResetProjectId() {
	_jsii_.InvokeVoid(
		n,
		"resetProjectId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementConnectivityTestSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

