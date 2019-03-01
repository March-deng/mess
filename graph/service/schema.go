package service

import (
	"errors"

	"github.com/graphql-go/graphql"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Location string `json:"location"`
	Phone    string `json:"phone"`
}

var UserInfo = graphql.NewObject(graphql.ObjectConfig{})

type Root struct {
	object *graphql.Object
}

type Resolver struct {
	s store
}

func (r *Resolver) ResolveUser(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)
	if ok {
		u, err := r.s.GetUserInfo(p.Context, name)
		return u, err
	}
	return nil, errors.New("name invalid")
}

func NewRoot(r *Resolver) *Root {
	return &Root{
		object: graphql.NewObject(graphql.ObjectConfig{
			Name: "query",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Name: "userInfos",
					Type: graphql.NewList(UserInfo),
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: r.ResolveUser,
				},
			},
		}),
	}
}
