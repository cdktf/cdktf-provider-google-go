//go:build no_runtime_type_checking

package firestoreindex

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirestoreIndexFieldsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FirestoreIndexFieldsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FirestoreIndexFieldsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FirestoreIndexFieldsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FirestoreIndexFieldsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FirestoreIndexFieldsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFirestoreIndexFieldsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

