package mongo

import (
	"context"
	"my_project/graph/service"
	"my_project/store"
)

type MgoStoreConn struct {
	m *store.Mgo
}

func (m *MgoStoreConn) GetUserInfo(ctx context.Context, name string) (service.User, error) {
	return service.User{}, nil
}
