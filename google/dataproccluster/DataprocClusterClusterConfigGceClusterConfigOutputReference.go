// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataproccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterClusterConfigGceClusterConfigOutputReference interface {
	cdktf.ComplexObject
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
	ConfidentialInstanceConfig() DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfigOutputReference
	ConfidentialInstanceConfigInput() *DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalIpOnly() interface{}
	SetInternalIpOnly(val interface{})
	InternalIpOnlyInput() interface{}
	InternalValue() *DataprocClusterClusterConfigGceClusterConfig
	SetInternalValue(val *DataprocClusterClusterConfigGceClusterConfig)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NodeGroupAffinity() DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityOutputReference
	NodeGroupAffinityInput() *DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity
	ReservationAffinity() DataprocClusterClusterConfigGceClusterConfigReservationAffinityOutputReference
	ReservationAffinityInput() *DataprocClusterClusterConfigGceClusterConfigReservationAffinity
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceAccountScopes() *[]*string
	SetServiceAccountScopes(val *[]*string)
	ServiceAccountScopesInput() *[]*string
	ShieldedInstanceConfig() DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutConfidentialInstanceConfig(value *DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfig)
	PutNodeGroupAffinity(value *DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity)
	PutReservationAffinity(value *DataprocClusterClusterConfigGceClusterConfigReservationAffinity)
	PutShieldedInstanceConfig(value *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig)
	ResetConfidentialInstanceConfig()
	ResetInternalIpOnly()
	ResetMetadata()
	ResetNetwork()
	ResetNodeGroupAffinity()
	ResetReservationAffinity()
	ResetServiceAccount()
	ResetServiceAccountScopes()
	ResetShieldedInstanceConfig()
	ResetSubnetwork()
	ResetTags()
	ResetZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigGceClusterConfigOutputReference
type jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ConfidentialInstanceConfig() DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfigOutputReference {
	var returns DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"confidentialInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ConfidentialInstanceConfigInput() *DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfig {
	var returns *DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfig
	_jsii_.Get(
		j,
		"confidentialInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) InternalIpOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) InternalIpOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) InternalValue() *DataprocClusterClusterConfigGceClusterConfig {
	var returns *DataprocClusterClusterConfigGceClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) NodeGroupAffinity() DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityOutputReference {
	var returns DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityOutputReference
	_jsii_.Get(
		j,
		"nodeGroupAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) NodeGroupAffinityInput() *DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	var returns *DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity
	_jsii_.Get(
		j,
		"nodeGroupAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ReservationAffinity() DataprocClusterClusterConfigGceClusterConfigReservationAffinityOutputReference {
	var returns DataprocClusterClusterConfigGceClusterConfigReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ReservationAffinityInput() *DataprocClusterClusterConfigGceClusterConfigReservationAffinity {
	var returns *DataprocClusterClusterConfigGceClusterConfigReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ServiceAccountScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ServiceAccountScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ShieldedInstanceConfig() DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference {
	var returns DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ShieldedInstanceConfigInput() *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig {
	var returns *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


func NewDataprocClusterClusterConfigGceClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigGceClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterClusterConfigGceClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigGceClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigGceClusterConfigOutputReference_Override(d DataprocClusterClusterConfigGceClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigGceClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetInternalIpOnly(val interface{}) {
	if err := j.validateSetInternalIpOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIpOnly",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetInternalValue(val *DataprocClusterClusterConfigGceClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetServiceAccountScopes(val *[]*string) {
	if err := j.validateSetServiceAccountScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountScopes",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) PutConfidentialInstanceConfig(value *DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfig) {
	if err := d.validatePutConfidentialInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfidentialInstanceConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) PutNodeGroupAffinity(value *DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity) {
	if err := d.validatePutNodeGroupAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNodeGroupAffinity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) PutReservationAffinity(value *DataprocClusterClusterConfigGceClusterConfigReservationAffinity) {
	if err := d.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) PutShieldedInstanceConfig(value *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig) {
	if err := d.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetConfidentialInstanceConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetConfidentialInstanceConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetInternalIpOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetInternalIpOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		d,
		"resetMetadata",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetNodeGroupAffinity() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeGroupAffinity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		d,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetServiceAccountScopes() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountScopes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ResetZone() {
	_jsii_.InvokeVoid(
		d,
		"resetZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

