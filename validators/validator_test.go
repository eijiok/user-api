package validators

import (
	"reflect"
	"testing"
)

func TestValidateRequired(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test empty string value",
			args: args{""},
			want: "field required",
		},
		{
			name: "test nil value",
			args: args{nil},
			want: "field required",
		},
		{
			name: "test non empty string",
			args: args{"a"},
			want: "",
		},
		{
			name: "test number",
			args: args{1},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateRequired(tt.args.value); got != tt.want {
				t.Errorf("ValidateRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatorPassword(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Validate password",
			args: args{"asdfv"},
			want: "",
		},
		{
			name: "Invalid password",
			args: args{"asdf"},
			want: "the password should have more than 5 characters",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatorPassword(tt.args.value); got != tt.want {
				t.Errorf("ValidatorPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateStringLength(t *testing.T) {
	type args struct {
		minSize *int
		maxSize *int
		value   any
	}
	valueOne := 1
	valueTwo := 2
	valueThree := 3
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "validate string length ok",
			args: args{&valueOne, &valueOne, "a"},
			want: "",
		},
		{
			name: "validate string greater than defined",
			args: args{&valueOne, &valueTwo, "asd"},
			want: "the length cannot be greater than 2",
		},
		{
			name: "validate string shorter than defined",
			args: args{&valueTwo, &valueThree, "a"},
			want: "the length cannot be lower than 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateStringLength(tt.args.minSize, tt.args.maxSize)(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
