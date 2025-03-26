// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerbranchrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/securesourcemanagerbranchrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/secure_source_manager_branch_rule google_secure_source_manager_branch_rule}.
type SecureSourceManagerBranchRule interface {
	cdktf.TerraformResource
	AllowStaleReviews() interface{}
	SetAllowStaleReviews(val interface{})
	AllowStaleReviewsInput() interface{}
	BranchRuleId() *string
	SetBranchRuleId(val *string)
	BranchRuleIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludePattern() *string
	SetIncludePattern(val *string)
	IncludePatternInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MinimumApprovalsCount() *float64
	SetMinimumApprovalsCount(val *float64)
	MinimumApprovalsCountInput() *float64
	MinimumReviewsCount() *float64
	SetMinimumReviewsCount(val *float64)
	MinimumReviewsCountInput() *float64
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RepositoryId() *string
	SetRepositoryId(val *string)
	RepositoryIdInput() *string
	RequireCommentsResolved() interface{}
	SetRequireCommentsResolved(val interface{})
	RequireCommentsResolvedInput() interface{}
	RequireLinearHistory() interface{}
	SetRequireLinearHistory(val interface{})
	RequireLinearHistoryInput() interface{}
	RequirePullRequest() interface{}
	SetRequirePullRequest(val interface{})
	RequirePullRequestInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SecureSourceManagerBranchRuleTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *SecureSourceManagerBranchRuleTimeouts)
	ResetAllowStaleReviews()
	ResetDisabled()
	ResetId()
	ResetMinimumApprovalsCount()
	ResetMinimumReviewsCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRequireCommentsResolved()
	ResetRequireLinearHistory()
	ResetRequirePullRequest()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SecureSourceManagerBranchRule
type jsiiProxy_SecureSourceManagerBranchRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) AllowStaleReviews() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStaleReviews",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) AllowStaleReviewsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStaleReviewsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) BranchRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) BranchRuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) IncludePattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) IncludePatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) MinimumApprovalsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumApprovalsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) MinimumApprovalsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumApprovalsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) MinimumReviewsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumReviewsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) MinimumReviewsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumReviewsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RepositoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RepositoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RequireCommentsResolved() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCommentsResolved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RequireCommentsResolvedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCommentsResolvedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RequireLinearHistory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLinearHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RequireLinearHistoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLinearHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RequirePullRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) RequirePullRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Timeouts() SecureSourceManagerBranchRuleTimeoutsOutputReference {
	var returns SecureSourceManagerBranchRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerBranchRule) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/secure_source_manager_branch_rule google_secure_source_manager_branch_rule} Resource.
func NewSecureSourceManagerBranchRule(scope constructs.Construct, id *string, config *SecureSourceManagerBranchRuleConfig) SecureSourceManagerBranchRule {
	_init_.Initialize()

	if err := validateNewSecureSourceManagerBranchRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecureSourceManagerBranchRule{}

	_jsii_.Create(
		"@cdktf/provider-google.secureSourceManagerBranchRule.SecureSourceManagerBranchRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/secure_source_manager_branch_rule google_secure_source_manager_branch_rule} Resource.
func NewSecureSourceManagerBranchRule_Override(s SecureSourceManagerBranchRule, scope constructs.Construct, id *string, config *SecureSourceManagerBranchRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.secureSourceManagerBranchRule.SecureSourceManagerBranchRule",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetAllowStaleReviews(val interface{}) {
	if err := j.validateSetAllowStaleReviewsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStaleReviews",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetBranchRuleId(val *string) {
	if err := j.validateSetBranchRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchRuleId",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetIncludePattern(val *string) {
	if err := j.validateSetIncludePatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePattern",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetMinimumApprovalsCount(val *float64) {
	if err := j.validateSetMinimumApprovalsCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumApprovalsCount",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetMinimumReviewsCount(val *float64) {
	if err := j.validateSetMinimumReviewsCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumReviewsCount",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetRepositoryId(val *string) {
	if err := j.validateSetRepositoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryId",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetRequireCommentsResolved(val interface{}) {
	if err := j.validateSetRequireCommentsResolvedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCommentsResolved",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetRequireLinearHistory(val interface{}) {
	if err := j.validateSetRequireLinearHistoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLinearHistory",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerBranchRule)SetRequirePullRequest(val interface{}) {
	if err := j.validateSetRequirePullRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirePullRequest",
		val,
	)
}

// Generates CDKTF code for importing a SecureSourceManagerBranchRule resource upon running "cdktf plan <stack-name>".
func SecureSourceManagerBranchRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSecureSourceManagerBranchRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.secureSourceManagerBranchRule.SecureSourceManagerBranchRule",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SecureSourceManagerBranchRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecureSourceManagerBranchRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.secureSourceManagerBranchRule.SecureSourceManagerBranchRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecureSourceManagerBranchRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecureSourceManagerBranchRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.secureSourceManagerBranchRule.SecureSourceManagerBranchRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecureSourceManagerBranchRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecureSourceManagerBranchRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.secureSourceManagerBranchRule.SecureSourceManagerBranchRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecureSourceManagerBranchRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.secureSourceManagerBranchRule.SecureSourceManagerBranchRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) PutTimeouts(value *SecureSourceManagerBranchRuleTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetAllowStaleReviews() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowStaleReviews",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetDisabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDisabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetMinimumApprovalsCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMinimumApprovalsCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetMinimumReviewsCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMinimumReviewsCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetProject() {
	_jsii_.InvokeVoid(
		s,
		"resetProject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetRequireCommentsResolved() {
	_jsii_.InvokeVoid(
		s,
		"resetRequireCommentsResolved",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetRequireLinearHistory() {
	_jsii_.InvokeVoid(
		s,
		"resetRequireLinearHistory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetRequirePullRequest() {
	_jsii_.InvokeVoid(
		s,
		"resetRequirePullRequest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerBranchRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

