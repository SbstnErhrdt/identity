package identity_interface_graphql

import (
	"errors"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func GetUserUIDFromContext(p *graphql.ResolveParams) (uid uuid.UUID, err error) {
	uid, ok := p.Context.Value("USER_UID").(uuid.UUID)
	if !ok {
		err = errors.New("can not get USER_UID from context")
		return
	}
	return
}
