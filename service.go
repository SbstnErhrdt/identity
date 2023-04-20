package identity

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/SbstnErhrdt/identity/identity_interface_graphql"
	"github.com/graphql-go/graphql"
)

type Service struct {
	identity_controllers.IdentityService
	gqlRootObject         *graphql.Object
	gqlRootMutationObject *graphql.Object
}

func NewIdentityService(controllerService identity_controllers.IdentityService) *Service {
	return &Service{
		IdentityService: controllerService,
	}
}

// SetGraphQLQueryInterface sets the graphql query interface
func (s *Service) SetGraphQLQueryInterface(rootQueryObject *graphql.Object) *Service {
	s.gqlRootObject = rootQueryObject
	// init queries
	q := identity_interface_graphql.InitGraphQlQueries(s)
	// connect to root query object
	q.GenerateQueryObjects(s.gqlRootObject)
	return s
}

// SetGraphQLMutationInterface sets the graphql mutation interface
func (s *Service) SetGraphQLMutationInterface(rootMutationObject *graphql.Object) *Service {
	s.gqlRootMutationObject = rootMutationObject
	// init mutations
	q := identity_interface_graphql.InitMutations(s)
	// connect to root mutation object
	q.GenerateMutationObjects(rootMutationObject)
	return s
}
