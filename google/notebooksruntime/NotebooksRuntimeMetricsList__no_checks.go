//go:build no_runtime_type_checking

package notebooksruntime

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotebooksRuntimeMetricsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotebooksRuntimeMetricsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotebooksRuntimeMetricsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotebooksRuntimeMetricsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotebooksRuntimeMetricsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotebooksRuntimeMetricsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

