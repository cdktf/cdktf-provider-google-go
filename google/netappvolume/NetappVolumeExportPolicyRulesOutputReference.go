// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/netappvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeExportPolicyRulesOutputReference interface {
	cdktf.ComplexObject
	AccessType() *string
	SetAccessType(val *string)
	AccessTypeInput() *string
	AllowedClients() *string
	SetAllowedClients(val *string)
	AllowedClientsInput() *string
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
	HasRootAccess() *string
	SetHasRootAccess(val *string)
	HasRootAccessInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Kerberos5IReadOnly() interface{}
	SetKerberos5IReadOnly(val interface{})
	Kerberos5IReadOnlyInput() interface{}
	Kerberos5IReadWrite() interface{}
	SetKerberos5IReadWrite(val interface{})
	Kerberos5IReadWriteInput() interface{}
	Kerberos5PReadOnly() interface{}
	SetKerberos5PReadOnly(val interface{})
	Kerberos5PReadOnlyInput() interface{}
	Kerberos5PReadWrite() interface{}
	SetKerberos5PReadWrite(val interface{})
	Kerberos5PReadWriteInput() interface{}
	Kerberos5ReadOnly() interface{}
	SetKerberos5ReadOnly(val interface{})
	Kerberos5ReadOnlyInput() interface{}
	Kerberos5ReadWrite() interface{}
	SetKerberos5ReadWrite(val interface{})
	Kerberos5ReadWriteInput() interface{}
	Nfsv3() interface{}
	SetNfsv3(val interface{})
	Nfsv3Input() interface{}
	Nfsv4() interface{}
	SetNfsv4(val interface{})
	Nfsv4Input() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAccessType()
	ResetAllowedClients()
	ResetHasRootAccess()
	ResetKerberos5IReadOnly()
	ResetKerberos5IReadWrite()
	ResetKerberos5PReadOnly()
	ResetKerberos5PReadWrite()
	ResetKerberos5ReadOnly()
	ResetKerberos5ReadWrite()
	ResetNfsv3()
	ResetNfsv4()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeExportPolicyRulesOutputReference
type jsiiProxy_NetappVolumeExportPolicyRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AllowedClients() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AllowedClientsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) HasRootAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hasRootAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) HasRootAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hasRootAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5IReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5IReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5IReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5IReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5PReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5PReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5PReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5PReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5ReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5ReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5ReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5ReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Nfsv3() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Nfsv3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Nfsv4() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Nfsv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetappVolumeExportPolicyRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetappVolumeExportPolicyRulesOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeExportPolicyRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeExportPolicyRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeExportPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetappVolumeExportPolicyRulesOutputReference_Override(n NetappVolumeExportPolicyRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeExportPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetAccessType(val *string) {
	if err := j.validateSetAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessType",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetAllowedClients(val *string) {
	if err := j.validateSetAllowedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetHasRootAccess(val *string) {
	if err := j.validateSetHasRootAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasRootAccess",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetKerberos5IReadOnly(val interface{}) {
	if err := j.validateSetKerberos5IReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5IReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetKerberos5IReadWrite(val interface{}) {
	if err := j.validateSetKerberos5IReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5IReadWrite",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetKerberos5PReadOnly(val interface{}) {
	if err := j.validateSetKerberos5PReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5PReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetKerberos5PReadWrite(val interface{}) {
	if err := j.validateSetKerberos5PReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5PReadWrite",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetKerberos5ReadOnly(val interface{}) {
	if err := j.validateSetKerberos5ReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5ReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetKerberos5ReadWrite(val interface{}) {
	if err := j.validateSetKerberos5ReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5ReadWrite",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetNfsv3(val interface{}) {
	if err := j.validateSetNfsv3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv3",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetNfsv4(val interface{}) {
	if err := j.validateSetNfsv4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv4",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetAccessType() {
	_jsii_.InvokeVoid(
		n,
		"resetAccessType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetAllowedClients() {
	_jsii_.InvokeVoid(
		n,
		"resetAllowedClients",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetHasRootAccess() {
	_jsii_.InvokeVoid(
		n,
		"resetHasRootAccess",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5IReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5IReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5IReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5IReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5PReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5PReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5PReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5PReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5ReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5ReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5ReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5ReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetNfsv3() {
	_jsii_.InvokeVoid(
		n,
		"resetNfsv3",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetNfsv4() {
	_jsii_.InvokeVoid(
		n,
		"resetNfsv4",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

