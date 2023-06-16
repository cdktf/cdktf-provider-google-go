package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/dialogflowcxpage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxPageTransitionRoutesOutputReference interface {
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
	Condition() *string
	SetCondition(val *string)
	ConditionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Intent() *string
	SetIntent(val *string)
	IntentInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	TargetFlow() *string
	SetTargetFlow(val *string)
	TargetFlowInput() *string
	TargetPage() *string
	SetTargetPage(val *string)
	TargetPageInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TriggerFulfillment() DialogflowCxPageTransitionRoutesTriggerFulfillmentOutputReference
	TriggerFulfillmentInput() *DialogflowCxPageTransitionRoutesTriggerFulfillment
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
	PutTriggerFulfillment(value *DialogflowCxPageTransitionRoutesTriggerFulfillment)
	ResetCondition()
	ResetIntent()
	ResetTargetFlow()
	ResetTargetPage()
	ResetTriggerFulfillment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxPageTransitionRoutesOutputReference
type jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) Intent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) IntentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) TargetFlow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) TargetFlowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) TargetPage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) TargetPageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) TriggerFulfillment() DialogflowCxPageTransitionRoutesTriggerFulfillmentOutputReference {
	var returns DialogflowCxPageTransitionRoutesTriggerFulfillmentOutputReference
	_jsii_.Get(
		j,
		"triggerFulfillment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) TriggerFulfillmentInput() *DialogflowCxPageTransitionRoutesTriggerFulfillment {
	var returns *DialogflowCxPageTransitionRoutesTriggerFulfillment
	_jsii_.Get(
		j,
		"triggerFulfillmentInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxPageTransitionRoutesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCxPageTransitionRoutesOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageTransitionRoutesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageTransitionRoutesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCxPageTransitionRoutesOutputReference_Override(d DialogflowCxPageTransitionRoutesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageTransitionRoutesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetCondition(val *string) {
	if err := j.validateSetConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetIntent(val *string) {
	if err := j.validateSetIntentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intent",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetTargetFlow(val *string) {
	if err := j.validateSetTargetFlowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetFlow",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetTargetPage(val *string) {
	if err := j.validateSetTargetPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPage",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) PutTriggerFulfillment(value *DialogflowCxPageTransitionRoutesTriggerFulfillment) {
	if err := d.validatePutTriggerFulfillmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTriggerFulfillment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		d,
		"resetCondition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ResetIntent() {
	_jsii_.InvokeVoid(
		d,
		"resetIntent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ResetTargetFlow() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetFlow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ResetTargetPage() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetPage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ResetTriggerFulfillment() {
	_jsii_.InvokeVoid(
		d,
		"resetTriggerFulfillment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

