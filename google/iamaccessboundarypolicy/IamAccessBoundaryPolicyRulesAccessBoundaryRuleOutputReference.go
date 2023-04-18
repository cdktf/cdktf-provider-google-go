package iamaccessboundarypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/iamaccessboundarypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference interface {
	cdktf.ComplexObject
	AvailabilityCondition() IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionOutputReference
	AvailabilityConditionInput() *IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition
	AvailablePermissions() *[]*string
	SetAvailablePermissions(val *[]*string)
	AvailablePermissionsInput() *[]*string
	AvailableResource() *string
	SetAvailableResource(val *string)
	AvailableResourceInput() *string
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
	InternalValue() *IamAccessBoundaryPolicyRulesAccessBoundaryRule
	SetInternalValue(val *IamAccessBoundaryPolicyRulesAccessBoundaryRule)
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
	PutAvailabilityCondition(value *IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition)
	ResetAvailabilityCondition()
	ResetAvailablePermissions()
	ResetAvailableResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference
type jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailabilityCondition() IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionOutputReference {
	var returns IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionOutputReference
	_jsii_.Get(
		j,
		"availabilityCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailabilityConditionInput() *IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition {
	var returns *IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition
	_jsii_.Get(
		j,
		"availabilityConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailablePermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availablePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailablePermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availablePermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailableResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailableResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) InternalValue() *IamAccessBoundaryPolicyRulesAccessBoundaryRule {
	var returns *IamAccessBoundaryPolicyRulesAccessBoundaryRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference {
	_init_.Initialize()

	if err := validateNewIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamAccessBoundaryPolicy.IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference_Override(i IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamAccessBoundaryPolicy.IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetAvailablePermissions(val *[]*string) {
	if err := j.validateSetAvailablePermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availablePermissions",
		val,
	)
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetAvailableResource(val *string) {
	if err := j.validateSetAvailableResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableResource",
		val,
	)
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetInternalValue(val *IamAccessBoundaryPolicyRulesAccessBoundaryRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) PutAvailabilityCondition(value *IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition) {
	if err := i.validatePutAvailabilityConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAvailabilityCondition",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ResetAvailabilityCondition() {
	_jsii_.InvokeVoid(
		i,
		"resetAvailabilityCondition",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ResetAvailablePermissions() {
	_jsii_.InvokeVoid(
		i,
		"resetAvailablePermissions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ResetAvailableResource() {
	_jsii_.InvokeVoid(
		i,
		"resetAvailableResource",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

