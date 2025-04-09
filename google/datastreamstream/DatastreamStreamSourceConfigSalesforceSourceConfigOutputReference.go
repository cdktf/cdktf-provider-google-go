// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/datastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference interface {
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
	ExcludeObjects() DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjectsOutputReference
	ExcludeObjectsInput() *DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjects
	// Experimental.
	Fqn() *string
	IncludeObjects() DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjectsOutputReference
	IncludeObjectsInput() *DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjects
	InternalValue() *DatastreamStreamSourceConfigSalesforceSourceConfig
	SetInternalValue(val *DatastreamStreamSourceConfigSalesforceSourceConfig)
	PollingInterval() *string
	SetPollingInterval(val *string)
	PollingIntervalInput() *string
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
	PutExcludeObjects(value *DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjects)
	PutIncludeObjects(value *DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjects)
	ResetExcludeObjects()
	ResetIncludeObjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference
type jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) ExcludeObjects() DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjectsOutputReference {
	var returns DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjectsOutputReference
	_jsii_.Get(
		j,
		"excludeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) ExcludeObjectsInput() *DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjects {
	var returns *DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjects
	_jsii_.Get(
		j,
		"excludeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) IncludeObjects() DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjectsOutputReference {
	var returns DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjectsOutputReference
	_jsii_.Get(
		j,
		"includeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) IncludeObjectsInput() *DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjects {
	var returns *DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjects
	_jsii_.Get(
		j,
		"includeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) InternalValue() *DatastreamStreamSourceConfigSalesforceSourceConfig {
	var returns *DatastreamStreamSourceConfigSalesforceSourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) PollingInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) PollingIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamSourceConfigSalesforceSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamSourceConfigSalesforceSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamSourceConfigSalesforceSourceConfigOutputReference_Override(d DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference)SetInternalValue(val *DatastreamStreamSourceConfigSalesforceSourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference)SetPollingInterval(val *string) {
	if err := j.validateSetPollingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pollingInterval",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) PutExcludeObjects(value *DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjects) {
	if err := d.validatePutExcludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludeObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) PutIncludeObjects(value *DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjects) {
	if err := d.validatePutIncludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIncludeObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) ResetExcludeObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) ResetIncludeObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSalesforceSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

