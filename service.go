package identity

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/SbstnErhrdt/identity/identity_interface_graphql"
	"github.com/graphql-go/graphql"
)

type Service struct {
	identity_controllers.IdentityService
	gqlRootQueryObject         *graphql.Object
	gqlRootMutationObject      *graphql.Object
	gqlRootAdminQueryObject    *graphql.Object
	gqlRootAdminMutationObject *graphql.Object
}

func NewIdentityService(controllerService identity_controllers.IdentityService) *Service {
	s := &Service{
		IdentityService: controllerService,
	}
	return s
}

// SetGraphQLQueryInterface sets the graphql query interface
func (s *Service) SetGraphQLQueryInterface(rootQueryObject *graphql.Object) *Service {
	s.gqlRootQueryObject = rootQueryObject
	// init queries
	q := identity_interface_graphql.InitGraphQlQueries(s)
	// connect to root query object
	q.GenerateQueryObjects(s.gqlRootQueryObject)
	return s
}

// SetGraphQLMutationInterface sets the graphql mutation interface
func (s *Service) SetGraphQLMutationInterface(rootMutationObject *graphql.Object) *Service {
	s.gqlRootMutationObject = rootMutationObject
	// init mutations
	q := identity_interface_graphql.InitGraphQlMutations(s)
	// connect to root mutation object
	q.GenerateMutationObjects(rootMutationObject)
	return s
}

// SetGraphQLAdminQueryInterface sets the graphql query interface
func (s *Service) SetGraphQLAdminQueryInterface(rootQueryObject *graphql.Object) *Service {
	s.gqlRootQueryObject = rootQueryObject
	// init queries
	q := identity_interface_graphql.InitAdminGraphQlQueries(s)
	// connect to root query object
	q.GenerateQueryObjects(s.gqlRootQueryObject)
	return s
}

// SetGraphQLAdminMutationInterface sets the graphql mutation interface
func (s *Service) SetGraphQLAdminMutationInterface(rootMutationObject *graphql.Object) *Service {
	s.gqlRootMutationObject = rootMutationObject
	// init mutations
	q := identity_interface_graphql.InitAdminGraphQlMutations(s)
	// connect to root mutation object
	q.GenerateMutationObjects(rootMutationObject)
	return s
}
