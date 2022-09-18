package graph

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_GetFollows(t *testing.T) {
	type fields struct {
		follow map[int64][]int64
	}
	type args struct {
		from int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int64
		wantErr bool
	}{
		{
			name:    "get follows 01",
			fields:  fields{follow: map[int64][]int64{1: {2, 3}}},
			args:    args{from: 1},
			want:    []int64{2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				follow: tt.fields.follow,
			}
			got, err := g.GetFollows(tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.GetFollows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.GetFollows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_GetFriendsOfFriends(t *testing.T) {
	type fields struct {
		follow map[int64][]int64
	}
	type args struct {
		from int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int64
		wantErr bool
	}{
		{
			name: "get fof 01",
			fields: fields{
				follow: map[int64][]int64{
					1: {2, 3},
					2: {3, 1},
					3: {5, 6},
				},
			},
			args:    args{from: 1},
			want:    []int64{5, 6},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				follow: tt.fields.follow,
			}
			got, err := g.GetFriendsOfFriends(tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.GetFriendsOfFriends() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.GetFriendsOfFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_RemoveFollow(t *testing.T) {
	g := NewGraph()
	g.AddFollow(1, 2)
	g.AddFollow(1, 3)
	g.AddFollow(1, 4)
	g.RemoveFollow(1, 3)
	r, err := g.GetFollows(1)
	if err != nil {
		t.Errorf("Graph.GetFriendsOfFriends() = %s", err.Error())
	}
	assert.Equal(t, r, []int64{2, 4})

}
