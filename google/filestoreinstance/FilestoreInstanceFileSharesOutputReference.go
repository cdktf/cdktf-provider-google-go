// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestoreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/filestoreinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FilestoreInstanceFileSharesOutputReference interface {
	cdktf.ComplexObject
	CapacityGb() *float64
	SetCapacityGb(val *float64)
	CapacityGbInput() *float64
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
	InternalValue() *FilestoreInstanceFileShares
	SetInternalValue(val *FilestoreInstanceFileShares)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NfsExportOptions() FilestoreInstanceFileSharesNfsExportOptionsList
	NfsExportOptionsInput() interface{}
	SourceBackup() *string
	SetSourceBackup(val *string)
	SourceBackupInput() *string
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
	PutNfsExportOptions(value interface{})
	ResetNfsExportOptions()
	ResetSourceBackup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FilestoreInstanceFileSharesOutputReference
type jsiiProxy_FilestoreInstanceFileSharesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) CapacityGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) CapacityGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) InternalValue() *FilestoreInstanceFileShares {
	var returns *FilestoreInstanceFileShares
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) NfsExportOptions() FilestoreInstanceFileSharesNfsExportOptionsList {
	var returns FilestoreInstanceFileSharesNfsExportOptionsList
	_jsii_.Get(
		j,
		"nfsExportOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) NfsExportOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsExportOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) SourceBackup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) SourceBackupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFilestoreInstanceFileSharesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FilestoreInstanceFileSharesOutputReference {
	_init_.Initialize()

	if err := validateNewFilestoreInstanceFileSharesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FilestoreInstanceFileSharesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.filestoreInstance.FilestoreInstanceFileSharesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFilestoreInstanceFileSharesOutputReference_Override(f FilestoreInstanceFileSharesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.filestoreInstance.FilestoreInstanceFileSharesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference)SetCapacityGb(val *float64) {
	if err := j.validateSetCapacityGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityGb",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference)SetInternalValue(val *FilestoreInstanceFileShares) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference)SetSourceBackup(val *string) {
	if err := j.validateSetSourceBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBackup",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) PutNfsExportOptions(value interface{}) {
	if err := f.validatePutNfsExportOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putNfsExportOptions",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) ResetNfsExportOptions() {
	_jsii_.InvokeVoid(
		f,
		"resetNfsExportOptions",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) ResetSourceBackup() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceBackup",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

