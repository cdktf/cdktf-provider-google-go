// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/bigquerytable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryTableExternalCatalogTableOptionsOutputReference interface {
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
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BigqueryTableExternalCatalogTableOptions
	SetInternalValue(val *BigqueryTableExternalCatalogTableOptions)
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	StorageDescriptor() BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference
	StorageDescriptorInput() *BigqueryTableExternalCatalogTableOptionsStorageDescriptor
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
	PutStorageDescriptor(value *BigqueryTableExternalCatalogTableOptionsStorageDescriptor)
	ResetConnectionId()
	ResetParameters()
	ResetStorageDescriptor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryTableExternalCatalogTableOptionsOutputReference
type jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) InternalValue() *BigqueryTableExternalCatalogTableOptions {
	var returns *BigqueryTableExternalCatalogTableOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) StorageDescriptor() BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference {
	var returns BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference
	_jsii_.Get(
		j,
		"storageDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) StorageDescriptorInput() *BigqueryTableExternalCatalogTableOptionsStorageDescriptor {
	var returns *BigqueryTableExternalCatalogTableOptionsStorageDescriptor
	_jsii_.Get(
		j,
		"storageDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBigqueryTableExternalCatalogTableOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryTableExternalCatalogTableOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryTableExternalCatalogTableOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalCatalogTableOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryTableExternalCatalogTableOptionsOutputReference_Override(b BigqueryTableExternalCatalogTableOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalCatalogTableOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference)SetInternalValue(val *BigqueryTableExternalCatalogTableOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) PutStorageDescriptor(value *BigqueryTableExternalCatalogTableOptionsStorageDescriptor) {
	if err := b.validatePutStorageDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStorageDescriptor",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ResetConnectionId() {
	_jsii_.InvokeVoid(
		b,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ResetStorageDescriptor() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageDescriptor",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

