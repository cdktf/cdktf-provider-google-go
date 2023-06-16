package vertexaifeaturestoreentitytype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/vertexaifeaturestoreentitytype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference interface {
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
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis
	SetInternalValue(val *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis)
	MonitoringIntervalDays() *float64
	SetMonitoringIntervalDays(val *float64)
	MonitoringIntervalDaysInput() *float64
	StalenessDays() *float64
	SetStalenessDays(val *float64)
	StalenessDaysInput() *float64
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
	ResetDisabled()
	ResetMonitoringIntervalDays()
	ResetStalenessDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference
type jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) InternalValue() *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis {
	var returns *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) MonitoringIntervalDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) MonitoringIntervalDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) StalenessDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stalenessDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) StalenessDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stalenessDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiFeaturestoreEntitytype.VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference_Override(v VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiFeaturestoreEntitytype.VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference)SetInternalValue(val *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference)SetMonitoringIntervalDays(val *float64) {
	if err := j.validateSetMonitoringIntervalDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringIntervalDays",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference)SetStalenessDays(val *float64) {
	if err := j.validateSetStalenessDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stalenessDays",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		v,
		"resetDisabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) ResetMonitoringIntervalDays() {
	_jsii_.InvokeVoid(
		v,
		"resetMonitoringIntervalDays",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) ResetStalenessDays() {
	_jsii_.InvokeVoid(
		v,
		"resetStalenessDays",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

