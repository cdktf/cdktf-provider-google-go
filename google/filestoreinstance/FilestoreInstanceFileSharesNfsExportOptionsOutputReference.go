package filestoreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/filestoreinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FilestoreInstanceFileSharesNfsExportOptionsOutputReference interface {
	cdktf.ComplexObject
	AccessMode() *string
	SetAccessMode(val *string)
	AccessModeInput() *string
	AnonGid() *float64
	SetAnonGid(val *float64)
	AnonGidInput() *float64
	AnonUid() *float64
	SetAnonUid(val *float64)
	AnonUidInput() *float64
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpRanges() *[]*string
	SetIpRanges(val *[]*string)
	IpRangesInput() *[]*string
	SquashMode() *string
	SetSquashMode(val *string)
	SquashModeInput() *string
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
	ResetAccessMode()
	ResetAnonGid()
	ResetAnonUid()
	ResetIpRanges()
	ResetSquashMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FilestoreInstanceFileSharesNfsExportOptionsOutputReference
type jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) AccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) AccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) AnonGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) AnonGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) AnonUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) AnonUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) IpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) IpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) SquashMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) SquashModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFilestoreInstanceFileSharesNfsExportOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FilestoreInstanceFileSharesNfsExportOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewFilestoreInstanceFileSharesNfsExportOptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.filestoreInstance.FilestoreInstanceFileSharesNfsExportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFilestoreInstanceFileSharesNfsExportOptionsOutputReference_Override(f FilestoreInstanceFileSharesNfsExportOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.filestoreInstance.FilestoreInstanceFileSharesNfsExportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetAccessMode(val *string) {
	if err := j.validateSetAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessMode",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetAnonGid(val *float64) {
	if err := j.validateSetAnonGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonGid",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetAnonUid(val *float64) {
	if err := j.validateSetAnonUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonUid",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetIpRanges(val *[]*string) {
	if err := j.validateSetIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRanges",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetSquashMode(val *string) {
	if err := j.validateSetSquashModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squashMode",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetAccessMode() {
	_jsii_.InvokeVoid(
		f,
		"resetAccessMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetAnonGid() {
	_jsii_.InvokeVoid(
		f,
		"resetAnonGid",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetAnonUid() {
	_jsii_.InvokeVoid(
		f,
		"resetAnonUid",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetIpRanges() {
	_jsii_.InvokeVoid(
		f,
		"resetIpRanges",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetSquashMode() {
	_jsii_.InvokeVoid(
		f,
		"resetSquashMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FilestoreInstanceFileSharesNfsExportOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

