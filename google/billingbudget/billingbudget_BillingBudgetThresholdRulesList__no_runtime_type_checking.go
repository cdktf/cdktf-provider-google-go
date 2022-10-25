//go:build no_runtime_type_checking

package billingbudget

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BillingBudgetThresholdRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BillingBudgetThresholdRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BillingBudgetThresholdRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BillingBudgetThresholdRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BillingBudgetThresholdRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BillingBudgetThresholdRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBillingBudgetThresholdRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

