package ladon

// EqualsSubjectCondition is a condition which is fulfilled if the request's subject is equal to the given value string
type EqualsSubjectCondition struct{}

// Fulfills returns true if the request's subject is equal to the given value string
func (c *EqualsSubjectCondition) Fulfills(value interface{}, r *Request) bool {
	s, ok := value.(string)
	if !ok {
		return ok
	}

	for _, sub := range r.Subjects {
		if sub == s {
			return true
		}
	}
	//return ok && s == r.Subject
	return false
}

// GetName returns the condition's name.
func (c *EqualsSubjectCondition) GetName() string {
	return "EqualsSubjectCondition"
}
