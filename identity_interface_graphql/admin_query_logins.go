package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/graphql-go/graphql"
)

// LoginAttemptsPaginationDTO is the data transfer object for the login attempts pagination
type LoginAttemptsPaginationDTO struct {
	Results []*identity_models.IdentityLogin `json:"results"`
	Amount  int64                            `json:"amount"`
}

// LoginAttemptsPaginationGraphQlModel is the graphql model for the login attempts pagination
var LoginAttemptsPaginationGraphQlModel = graphql.NewObject(graphql.ObjectConfig{
	Name:        "LoginAttemptsField",
	Description: "Login attempts pagination",
	Fields: graphql.Fields{
		"results": &graphql.Field{
			Type:        graphql.NewList(LoginAttemptGraphQlModel),
			Description: "list of login attempts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*LoginAttemptsPaginationDTO); ok {
					return obj.Results, nil
				}
				return nil, errors.New("can not cast object")
			},
		},
		"amount": &graphql.Field{
			Type:        graphql.Int,
			Description: "The total amount of projects",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*LoginAttemptsPaginationDTO); ok {
					return obj.Amount, nil
				}
				return nil, errors.New("can not cast object")
			},
		},
	},
})

// LoginAttemptsField is the graphql field for searching the logins
func LoginAttemptsField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "LoginAttempts",
		Description: "Retrieve all login attempts of the system",
		Type:        LoginAttemptsPaginationGraphQlModel,
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
				Description:  "The searchable email address",
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
				DefaultValue: "created_at desc",
				Description:  "Order by operation of the pagination",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			err = CheckAdmin(service, &p)
			if err != nil {
				return nil, err
			}
			// params
			keyword := p.Args["keyword"].(string)
			from := p.Args["from"].(int)
			size := p.Args["size"].(int)
			orderBy := p.Args["orderBy"].(string)
			// get identities
			res, amount, err := identity_controllers.ReadIdentityLogins(service, keyword, from, size, orderBy)
			if err != nil {
				return nil, err
			}
			dto := LoginAttemptsPaginationDTO{
				Results: res,
				Amount:  amount,
			}
			return &dto, nil
		},
	}
	return &field
}
