// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeaturestoreentitytype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/vertexaifeaturestoreentitytype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference interface {
	cdktf.ComplexObject
	AnomalyDetectionBaseline() *string
	SetAnomalyDetectionBaseline(val *string)
	AnomalyDetectionBaselineInput() *string
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
	InternalValue() *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis
	SetInternalValue(val *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis)
	State() *string
	SetState(val *string)
	StateInput() *string
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
	ResetAnomalyDetectionBaseline()
	ResetState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference
type jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) AnomalyDetectionBaseline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectionBaseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) AnomalyDetectionBaselineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectionBaselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) InternalValue() *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis {
	var returns *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiFeaturestoreEntitytype.VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference_Override(v VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiFeaturestoreEntitytype.VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference)SetAnomalyDetectionBaseline(val *string) {
	if err := j.validateSetAnomalyDetectionBaselineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anomalyDetectionBaseline",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference)SetInternalValue(val *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) ResetAnomalyDetectionBaseline() {
	_jsii_.InvokeVoid(
		v,
		"resetAnomalyDetectionBaseline",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		v,
		"resetState",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

