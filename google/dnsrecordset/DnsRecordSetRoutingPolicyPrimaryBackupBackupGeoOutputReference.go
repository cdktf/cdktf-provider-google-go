// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsrecordset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dnsrecordset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HealthCheckedTargets() DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference
	HealthCheckedTargetsInput() *DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Rrdatas() *[]*string
	SetRrdatas(val *[]*string)
	RrdatasInput() *[]*string
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
	PutHealthCheckedTargets(value *DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets)
	ResetHealthCheckedTargets()
	ResetRrdatas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference
type jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) HealthCheckedTargets() DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference {
	var returns DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference
	_jsii_.Get(
		j,
		"healthCheckedTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) HealthCheckedTargetsInput() *DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets {
	var returns *DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets
	_jsii_.Get(
		j,
		"healthCheckedTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) Rrdatas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rrdatas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) RrdatasInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rrdatasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference {
	_init_.Initialize()

	if err := validateNewDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dnsRecordSet.DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference_Override(d DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dnsRecordSet.DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetRrdatas(val *[]*string) {
	if err := j.validateSetRrdatasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rrdatas",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) PutHealthCheckedTargets(value *DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets) {
	if err := d.validatePutHealthCheckedTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHealthCheckedTargets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ResetHealthCheckedTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthCheckedTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ResetRrdatas() {
	_jsii_.InvokeVoid(
		d,
		"resetRrdatas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

