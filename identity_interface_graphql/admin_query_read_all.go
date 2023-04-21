package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/graphql-go/graphql"
)

// IdentitiesPaginationDTO is the dto for the identities pagination
type IdentitiesPaginationDTO struct {
	Amount  int64
	Results []*identity_models.Identity
}

// IdentitiesPaginationGraphQlModel is the graphql model for the identities pagination
var IdentitiesPaginationGraphQlModel = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Identities",
	Description: "An identities result object",
	Fields: graphql.Fields{
		"results": &graphql.Field{
			Type:        graphql.NewList(IdentityGraphQlModel),
			Description: "list of projects",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*IdentitiesPaginationDTO); ok {
					return obj.Results, nil
				}
				return nil, errors.New("can not cast object")
			},
		},
		"amount": &graphql.Field{
			Type:        graphql.Int,
			Description: "The total amount of projects",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*IdentitiesPaginationDTO); ok {
					return obj.Amount, nil
				}
				return nil, errors.New("can not cast object")
			},
		},
	},
})

// IdentitiesSearchField is the graphql field for the search identities
func IdentitiesSearchField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "IdentitiesSearch",
		Description: "Retrieve the identity of the current user",
		Type:        IdentitiesPaginationGraphQlModel,
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
				Description:  "The searchable keyword",
			},
			"from": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 0,
				Description:  "The start of the pagination of the page",
			},
			"size": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 30,
				Description:  "The size of the page of the pagination",
			},
			"orderBy": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
				Description:  "Order by operation of the pagination",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// params
			keyword := p.Args["keyword"].(string)
			from := p.Args["from"].(int)
			size := p.Args["size"].(int)
			orderBy := p.Args["orderBy"].(string)

			// from context
			userUID, err := GetUserUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			res, amount, err := identity_controllers.ReadAllUsers(service, userUID, keyword, from, size, orderBy)
			if err != nil {
				return nil, err
			}
			dto := IdentitiesPaginationDTO{
				Results: res,
				Amount:  amount,
			}
			return &dto, nil
		},
	}
	return &field
}
