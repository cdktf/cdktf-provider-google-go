package cloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/cloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildTriggerBuildSourceRepoSourceOutputReference interface {
	cdktf.ComplexObject
	BranchName() *string
	SetBranchName(val *string)
	BranchNameInput() *string
	CommitSha() *string
	SetCommitSha(val *string)
	CommitShaInput() *string
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
	Dir() *string
	SetDir(val *string)
	DirInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudbuildTriggerBuildSourceRepoSource
	SetInternalValue(val *CloudbuildTriggerBuildSourceRepoSource)
	InvertRegex() interface{}
	SetInvertRegex(val interface{})
	InvertRegexInput() interface{}
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	RepoName() *string
	SetRepoName(val *string)
	RepoNameInput() *string
	Substitutions() *map[string]*string
	SetSubstitutions(val *map[string]*string)
	SubstitutionsInput() *map[string]*string
	TagName() *string
	SetTagName(val *string)
	TagNameInput() *string
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
	ResetBranchName()
	ResetCommitSha()
	ResetDir()
	ResetInvertRegex()
	ResetProjectId()
	ResetSubstitutions()
	ResetTagName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudbuildTriggerBuildSourceRepoSourceOutputReference
type jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) BranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) CommitSha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) CommitShaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitShaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) Dir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) DirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) InternalValue() *CloudbuildTriggerBuildSourceRepoSource {
	var returns *CloudbuildTriggerBuildSourceRepoSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) InvertRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) InvertRegexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) RepoName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) RepoNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) Substitutions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) SubstitutionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) TagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) TagNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudbuildTriggerBuildSourceRepoSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudbuildTriggerBuildSourceRepoSourceOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildTriggerBuildSourceRepoSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBuildSourceRepoSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildTriggerBuildSourceRepoSourceOutputReference_Override(c CloudbuildTriggerBuildSourceRepoSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBuildSourceRepoSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetBranchName(val *string) {
	if err := j.validateSetBranchNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetCommitSha(val *string) {
	if err := j.validateSetCommitShaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitSha",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetDir(val *string) {
	if err := j.validateSetDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dir",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetInternalValue(val *CloudbuildTriggerBuildSourceRepoSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetInvertRegex(val interface{}) {
	if err := j.validateSetInvertRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invertRegex",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetRepoName(val *string) {
	if err := j.validateSetRepoNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repoName",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetSubstitutions(val *map[string]*string) {
	if err := j.validateSetSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"substitutions",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetTagName(val *string) {
	if err := j.validateSetTagNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagName",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ResetBranchName() {
	_jsii_.InvokeVoid(
		c,
		"resetBranchName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ResetCommitSha() {
	_jsii_.InvokeVoid(
		c,
		"resetCommitSha",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ResetDir() {
	_jsii_.InvokeVoid(
		c,
		"resetDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ResetInvertRegex() {
	_jsii_.InvokeVoid(
		c,
		"resetInvertRegex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ResetProjectId() {
	_jsii_.InvokeVoid(
		c,
		"resetProjectId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ResetSubstitutions() {
	_jsii_.InvokeVoid(
		c,
		"resetSubstitutions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ResetTagName() {
	_jsii_.InvokeVoid(
		c,
		"resetTagName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildSourceRepoSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

