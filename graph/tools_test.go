package graph

import (
	"reflect"
	"testing"
)

func Test_removeFromSlice(t *testing.T) {
	type args struct {
		l    []int64
		item int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "remove from slice",
			args: args{
				l:    []int64{1, 2, 3},
				item: 3,
			},
			want: []int64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeFromSlice(tt.args.l, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
