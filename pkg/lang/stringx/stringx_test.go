package stringx

import "testing"

func TestCamel2Snake(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "a"},
			want: "a",
		},
		{
			args: args{s: "A"},
			want: "a",
		},
		{
			args: args{s: "aA"},
			want: "a_a",
		},
		{
			args: args{s: "AA"},
			want: "a_a",
		},
		{
			args: args{s: "aAa"},
			want: "a_aa",
		},
		{
			args: args{s: "AAa"},
			want: "a_aa",
		},
		{
			args: args{s: "aAA"},
			want: "a_a_a",
		},
		{
			args: args{s: "AAA"},
			want: "a_a_a",
		},
		{
			args: args{s: "AAaA"},
			want: "a_aa_a",
		},
		{
			args: args{s: "aAAa"},
			want: "a_a_aa",
		},
		{
			args: args{s: "aAaA"},
			want: "a_aa_a",
		},
		{
			args: args{s: "XxYy"},
			want: "xx_yy",
		},
		{
			args: args{s: "XxYY"},
			want: "xx_y_y",
		},
		{
			args: args{s: "xxYy"},
			want: "xx_yy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel2Snake(tt.args.s); got != tt.want {
				t.Errorf("Camel2Snake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstLower(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "a"},
			want: "a",
		},
		{
			args: args{str: "A"},
			want: "a",
		},
		{
			args: args{str: "AA"},
			want: "aA",
		},
		{
			args: args{str: "Aa"},
			want: "aa",
		},
		{
			args: args{str: "aAa"},
			want: "aAa",
		},
		{
			args: args{str: "AAa"},
			want: "aAa",
		},
		{
			args: args{str: "aAA"},
			want: "aAA",
		},
		{
			args: args{str: "AAA"},
			want: "aAA",
		},
		{
			args: args{str: "AAaA"},
			want: "aAaA",
		},
		{
			args: args{str: "aAAa"},
			want: "aAAa",
		},
		{
			args: args{str: "aAaA"},
			want: "aAaA",
		},
		{
			args: args{str: "XxYy"},
			want: "xxYy",
		},
		{
			args: args{str: "XxYY"},
			want: "xxYY",
		},
		{
			args: args{str: "xxYy"},
			want: "xxYy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstLower(tt.args.str); got != tt.want {
				t.Errorf("FirstLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstUpper(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "a"},
			want: "A",
		},
		{
			args: args{str: "A"},
			want: "A",
		},
		{
			args: args{str: "AA"},
			want: "AA",
		},
		{
			args: args{str: "Aa"},
			want: "Aa",
		},
		{
			args: args{str: "aAa"},
			want: "AAa",
		},
		{
			args: args{str: "AAa"},
			want: "AAa",
		},
		{
			args: args{str: "aAA"},
			want: "AAA",
		},
		{
			args: args{str: "AAA"},
			want: "AAA",
		},
		{
			args: args{str: "AAaA"},
			want: "AAaA",
		},
		{
			args: args{str: "aAAa"},
			want: "AAAa",
		},
		{
			args: args{str: "aAaA"},
			want: "AAaA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstUpper(tt.args.str); got != tt.want {
				t.Errorf("FirstUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnake2BigCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "xx_yy",
			},
			want: "XxYy",
		},
		{
			name: "test2",
			args: args{
				s: "xx_y_y",
			},
			want: "XxYY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snake2BigCamel(tt.args.s); got != tt.want {
				t.Errorf("Snake2BigCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
