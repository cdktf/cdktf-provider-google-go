package healthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/healthcarefhirstore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference interface {
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
	DatasetUri() *string
	SetDatasetUri(val *string)
	DatasetUriInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HealthcareFhirStoreStreamConfigsBigqueryDestination
	SetInternalValue(val *HealthcareFhirStoreStreamConfigsBigqueryDestination)
	SchemaConfig() HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference
	SchemaConfigInput() *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig
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
	PutSchemaConfig(value *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference
type jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) DatasetUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) DatasetUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) InternalValue() *HealthcareFhirStoreStreamConfigsBigqueryDestination {
	var returns *HealthcareFhirStoreStreamConfigsBigqueryDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) SchemaConfig() HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference {
	var returns HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference
	_jsii_.Get(
		j,
		"schemaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) SchemaConfigInput() *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	var returns *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig
	_jsii_.Get(
		j,
		"schemaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewHealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference_Override(h HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference)SetDatasetUri(val *string) {
	if err := j.validateSetDatasetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetUri",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference)SetInternalValue(val *HealthcareFhirStoreStreamConfigsBigqueryDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) PutSchemaConfig(value *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig) {
	if err := h.validatePutSchemaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putSchemaConfig",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

