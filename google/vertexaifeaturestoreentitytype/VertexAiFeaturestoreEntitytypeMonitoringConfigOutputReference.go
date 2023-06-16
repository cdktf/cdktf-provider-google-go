package vertexaifeaturestoreentitytype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/vertexaifeaturestoreentitytype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference interface {
	cdktf.ComplexObject
	CategoricalThresholdConfig() VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfigOutputReference
	CategoricalThresholdConfigInput() *VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig
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
	ImportFeaturesAnalysis() VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference
	ImportFeaturesAnalysisInput() *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis
	InternalValue() *VertexAiFeaturestoreEntitytypeMonitoringConfig
	SetInternalValue(val *VertexAiFeaturestoreEntitytypeMonitoringConfig)
	NumericalThresholdConfig() VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfigOutputReference
	NumericalThresholdConfigInput() *VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig
	SnapshotAnalysis() VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference
	SnapshotAnalysisInput() *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis
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
	PutCategoricalThresholdConfig(value *VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig)
	PutImportFeaturesAnalysis(value *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis)
	PutNumericalThresholdConfig(value *VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig)
	PutSnapshotAnalysis(value *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis)
	ResetCategoricalThresholdConfig()
	ResetImportFeaturesAnalysis()
	ResetNumericalThresholdConfig()
	ResetSnapshotAnalysis()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference
type jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) CategoricalThresholdConfig() VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfigOutputReference {
	var returns VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfigOutputReference
	_jsii_.Get(
		j,
		"categoricalThresholdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) CategoricalThresholdConfigInput() *VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig {
	var returns *VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig
	_jsii_.Get(
		j,
		"categoricalThresholdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ImportFeaturesAnalysis() VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference {
	var returns VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference
	_jsii_.Get(
		j,
		"importFeaturesAnalysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ImportFeaturesAnalysisInput() *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis {
	var returns *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis
	_jsii_.Get(
		j,
		"importFeaturesAnalysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) InternalValue() *VertexAiFeaturestoreEntitytypeMonitoringConfig {
	var returns *VertexAiFeaturestoreEntitytypeMonitoringConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) NumericalThresholdConfig() VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfigOutputReference {
	var returns VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfigOutputReference
	_jsii_.Get(
		j,
		"numericalThresholdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) NumericalThresholdConfigInput() *VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig {
	var returns *VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig
	_jsii_.Get(
		j,
		"numericalThresholdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) SnapshotAnalysis() VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference {
	var returns VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference
	_jsii_.Get(
		j,
		"snapshotAnalysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) SnapshotAnalysisInput() *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis {
	var returns *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis
	_jsii_.Get(
		j,
		"snapshotAnalysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiFeaturestoreEntitytypeMonitoringConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiFeaturestoreEntitytype.VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference_Override(v VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiFeaturestoreEntitytype.VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference)SetInternalValue(val *VertexAiFeaturestoreEntitytypeMonitoringConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) PutCategoricalThresholdConfig(value *VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig) {
	if err := v.validatePutCategoricalThresholdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCategoricalThresholdConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) PutImportFeaturesAnalysis(value *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis) {
	if err := v.validatePutImportFeaturesAnalysisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putImportFeaturesAnalysis",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) PutNumericalThresholdConfig(value *VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig) {
	if err := v.validatePutNumericalThresholdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNumericalThresholdConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) PutSnapshotAnalysis(value *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis) {
	if err := v.validatePutSnapshotAnalysisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSnapshotAnalysis",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ResetCategoricalThresholdConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetCategoricalThresholdConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ResetImportFeaturesAnalysis() {
	_jsii_.InvokeVoid(
		v,
		"resetImportFeaturesAnalysis",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ResetNumericalThresholdConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetNumericalThresholdConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ResetSnapshotAnalysis() {
	_jsii_.InvokeVoid(
		v,
		"resetSnapshotAnalysis",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

