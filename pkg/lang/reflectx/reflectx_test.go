package reflectx

import (
	"reflect"
	"testing"
)

func TestInterface2Array(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "test1",
			args: args{s: []interface{}{"a", "b", "c"}},
			want: []interface{}{"a", "b", "c"},
		},
		{
			name: "test2",
			args: args{s: []interface{}{"a", 1, "b", 2, "c", 3}},
			want: []interface{}{"a", 1, "b", 2, "c", 3},
		},
		{
			name: "test3",
			args: args{s: []interface{}{}},
			want: nil,
		},
		{
			name: "test4",
			args: args{s: "a"},
			want: []interface{}{"a"},
		},
		{
			name: "test5",
			args: args{s: []interface{}{1}},
			want: []interface{}{1},
		},
		{
			name: "test6",
			args: args{s: true},
			want: []interface{}{true},
		},
		{
			name: "test7",
			args: args{s: []interface{}{1.1}},
			want: []interface{}{1.1},
		},
		{
			name: "test8",
			args: args{s: []interface{}{[]int{1, 2, 3}}},
			want: []interface{}{[]int{1, 2, 3}},
		},
		{
			name: "test9",
			args: args{s: []interface{}{map[int]string{1: "a", 2: "b", 3: "c"}}},
			want: []interface{}{map[int]string{1: "a", 2: "b", 3: "c"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interface2Array(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface2Array() = %v, want %v", got, tt.want)
			}
		})
	}
}
