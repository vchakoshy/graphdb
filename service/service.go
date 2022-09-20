package service

import (
	"context"

	"github.com/vchakoshy/graphdb/graph"
)

var graphClient *graph.Graph

func SetGraphClient(g *graph.Graph) {
	graphClient = g
}

type ImplementedGraphdbServer struct{}

func (ImplementedGraphdbServer) GetFriendsOfFriends(c context.Context, u *User) (*UserSlice, error) {
	ul, err := graphClient.GetFriendsOfFriends(u.GetId())
	if err != nil {
		return nil, err
	}

	us := &UserSlice{}
	for _, u := range ul {
		us.Users = append(us.Users, &User{Id: u})
	}

	return us, nil
}

func (ImplementedGraphdbServer) AddFollow(c context.Context, f *Follow) (*Follow, error) {
	graphClient.AddFollow(f.GetFrom(), f.GetTo())
	return f, nil
}

func (ImplementedGraphdbServer) RemoveFollow(c context.Context, f *Follow) (*Follow, error) {
	graphClient.RemoveFollow(f.GetFrom(), f.GetTo())
	return f, nil
}

func (ImplementedGraphdbServer) mustEmbedUnimplementedGraphdbServer() {}
