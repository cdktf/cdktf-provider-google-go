package binaryauthorizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/binaryauthorizationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference interface {
	cdktf.ComplexObject
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
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
	EnforcementMode() *string
	SetEnforcementMode(val *string)
	EnforcementModeInput() *string
	EvaluationMode() *string
	SetEvaluationMode(val *string)
	EvaluationModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RequireAttestationsBy() *[]*string
	SetRequireAttestationsBy(val *[]*string)
	RequireAttestationsByInput() *[]*string
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
	ResetRequireAttestationsBy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference
type jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) EnforcementMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) EnforcementModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) EvaluationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) EvaluationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) RequireAttestationsBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requireAttestationsBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) RequireAttestationsByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requireAttestationsByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBinaryAuthorizationPolicyClusterAdmissionRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference {
	_init_.Initialize()

	if err := validateNewBinaryAuthorizationPolicyClusterAdmissionRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.binaryAuthorizationPolicy.BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBinaryAuthorizationPolicyClusterAdmissionRulesOutputReference_Override(b BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.binaryAuthorizationPolicy.BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetEnforcementMode(val *string) {
	if err := j.validateSetEnforcementModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcementMode",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetEvaluationMode(val *string) {
	if err := j.validateSetEvaluationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationMode",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetRequireAttestationsBy(val *[]*string) {
	if err := j.validateSetRequireAttestationsByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAttestationsBy",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) ResetRequireAttestationsBy() {
	_jsii_.InvokeVoid(
		b,
		"resetRequireAttestationsBy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BinaryAuthorizationPolicyClusterAdmissionRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

