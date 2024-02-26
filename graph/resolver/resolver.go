package resolver

import "unjammed/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MemberStore       map[string]model.Member
	OrganizationStore map[string]model.Organization
	EmploymentStore   map[string]model.Employment
}
