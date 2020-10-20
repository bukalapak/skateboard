package element

func compute(a *API) {
	for i := range a.ResourceGroups {
		for j := range a.ResourceGroups[i].Resources {
			computeResource(a.ResourceGroups[i].Resources[j])
		}
	}

	for i := range a.Resources {
		computeResource(a.Resources[i])
	}
}

func computeResource(resource Resource) {
	for k := range resource.Transitions {
		computeTransition(&resource.Transitions[k], resource)
	}
}

func computeTransition(transition *Transition, r Resource) {
	transition.Method = computeMethod(*transition)
	transition.Href = computeHref(*transition, r)
}

func computeMethod(t Transition) string {
	for _, x := range t.Transactions {
		if x.Request.Method != "" {
			return x.Request.Method
		}
	}

	return ""
}

func computeHref(t Transition, r Resource) Href {
	if t.Href.Path == "" {
		return r.Href
	}

	return t.Href
}
