package billingbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/billingbudget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BillingBudgetBudgetFilterOutputReference interface {
	cdktf.ComplexObject
	CalendarPeriod() *string
	SetCalendarPeriod(val *string)
	CalendarPeriodInput() *string
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
	CreditTypes() *[]*string
	SetCreditTypes(val *[]*string)
	CreditTypesInput() *[]*string
	CreditTypesTreatment() *string
	SetCreditTypesTreatment(val *string)
	CreditTypesTreatmentInput() *string
	CustomPeriod() BillingBudgetBudgetFilterCustomPeriodOutputReference
	CustomPeriodInput() *BillingBudgetBudgetFilterCustomPeriod
	// Experimental.
	Fqn() *string
	InternalValue() *BillingBudgetBudgetFilter
	SetInternalValue(val *BillingBudgetBudgetFilter)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Projects() *[]*string
	SetProjects(val *[]*string)
	ProjectsInput() *[]*string
	ResourceAncestors() *[]*string
	SetResourceAncestors(val *[]*string)
	ResourceAncestorsInput() *[]*string
	Services() *[]*string
	SetServices(val *[]*string)
	ServicesInput() *[]*string
	Subaccounts() *[]*string
	SetSubaccounts(val *[]*string)
	SubaccountsInput() *[]*string
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
	PutCustomPeriod(value *BillingBudgetBudgetFilterCustomPeriod)
	ResetCalendarPeriod()
	ResetCreditTypes()
	ResetCreditTypesTreatment()
	ResetCustomPeriod()
	ResetLabels()
	ResetProjects()
	ResetResourceAncestors()
	ResetServices()
	ResetSubaccounts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BillingBudgetBudgetFilterOutputReference
type jsiiProxy_BillingBudgetBudgetFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CalendarPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CalendarPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CreditTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creditTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CreditTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creditTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CreditTypesTreatment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creditTypesTreatment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CreditTypesTreatmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creditTypesTreatmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CustomPeriod() BillingBudgetBudgetFilterCustomPeriodOutputReference {
	var returns BillingBudgetBudgetFilterCustomPeriodOutputReference
	_jsii_.Get(
		j,
		"customPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) CustomPeriodInput() *BillingBudgetBudgetFilterCustomPeriod {
	var returns *BillingBudgetBudgetFilterCustomPeriod
	_jsii_.Get(
		j,
		"customPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) InternalValue() *BillingBudgetBudgetFilter {
	var returns *BillingBudgetBudgetFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) Projects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResourceAncestors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceAncestors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResourceAncestorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceAncestorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) Services() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) Subaccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subaccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) SubaccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subaccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBillingBudgetBudgetFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BillingBudgetBudgetFilterOutputReference {
	_init_.Initialize()

	if err := validateNewBillingBudgetBudgetFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingBudgetBudgetFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.billingBudget.BillingBudgetBudgetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBillingBudgetBudgetFilterOutputReference_Override(b BillingBudgetBudgetFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.billingBudget.BillingBudgetBudgetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetCalendarPeriod(val *string) {
	if err := j.validateSetCalendarPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calendarPeriod",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetCreditTypes(val *[]*string) {
	if err := j.validateSetCreditTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creditTypes",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetCreditTypesTreatment(val *string) {
	if err := j.validateSetCreditTypesTreatmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creditTypesTreatment",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetInternalValue(val *BillingBudgetBudgetFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetProjects(val *[]*string) {
	if err := j.validateSetProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projects",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetResourceAncestors(val *[]*string) {
	if err := j.validateSetResourceAncestorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceAncestors",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetServices(val *[]*string) {
	if err := j.validateSetServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"services",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetSubaccounts(val *[]*string) {
	if err := j.validateSetSubaccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subaccounts",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) PutCustomPeriod(value *BillingBudgetBudgetFilterCustomPeriod) {
	if err := b.validatePutCustomPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomPeriod",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetCalendarPeriod() {
	_jsii_.InvokeVoid(
		b,
		"resetCalendarPeriod",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetCreditTypes() {
	_jsii_.InvokeVoid(
		b,
		"resetCreditTypes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetCreditTypesTreatment() {
	_jsii_.InvokeVoid(
		b,
		"resetCreditTypesTreatment",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetCustomPeriod() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomPeriod",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		b,
		"resetLabels",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetProjects() {
	_jsii_.InvokeVoid(
		b,
		"resetProjects",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetResourceAncestors() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceAncestors",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetServices() {
	_jsii_.InvokeVoid(
		b,
		"resetServices",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ResetSubaccounts() {
	_jsii_.InvokeVoid(
		b,
		"resetSubaccounts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

