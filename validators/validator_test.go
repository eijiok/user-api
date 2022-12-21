package validators

import (
	"reflect"
	"testing"
	"time"
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
		{
			name: "Invalid password",
			args: args{123},
			want: "not a string value",
		},
		{
			name: "Empty string is not validated",
			args: args{""},
			want: "",
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
		{
			name: "nil string is not validated",
			args: args{&valueTwo, &valueThree, nil},
			want: "",
		},
		{
			name: "numbers are invalid values",
			args: args{&valueTwo, &valueThree, 123},
			want: "not a string value",
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

func TestDateTimeValidator(t *testing.T) {
	type args struct {
		minTime *time.Time
		maxTime *time.Time
		value   any
	}
	yesterday := time.Date(2022, 12, 17, 23, 12, 30, 0, time.UTC)
	today := time.Date(2022, 12, 18, 23, 12, 30, 0, time.UTC)
	tomorrow := time.Date(2022, 12, 19, 23, 12, 30, 0, time.UTC)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test DateTime between dates",
			args: args{&yesterday, &tomorrow, today},
			want: "",
		},
		{
			name: "test DateTime after max date",
			args: args{&yesterday, &today, tomorrow},
			want: "the date time Mon Dec 19 23:12:30 2022 cannot be after than Sun Dec 18 23:12:30 2022",
		},
		{
			name: "test DateTime before min date",
			args: args{&today, &tomorrow, yesterday},
			want: "the date time Sat Dec 17 23:12:30 2022 cannot be before than Sun Dec 18 23:12:30 2022",
		},
		{
			name: "nil is not validated",
			args: args{&yesterday, &today, nil},
			want: "",
		},
		{
			name: "zero time is not validated",
			args: args{&yesterday, &today, time.Time{}},
			want: "",
		},
		{
			name: "numbers are invalid",
			args: args{&yesterday, &today, 123},
			want: "not a valid time format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateTimeValidator(tt.args.minTime, tt.args.maxTime)(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTimeValidator() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestIsNullOrEmpty(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test nil",
			args: args{nil},
			want: true,
		},
		{
			name: "test empty string",
			args: args{""},
			want: true,
		},
		{
			name: "test non empty string",
			args: args{"asdf"},
			want: false,
		},
		{
			name: "test number",
			args: args{0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNullOrEmpty(tt.args.value); got != tt.want {
				t.Errorf("IsNullOrEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatorEmailFormat(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test email validator",
			args: args{"eiji.ok@gmail.com"},
			want: "",
		},
		{
			name: "test email validator",
			args: args{"eiji.okgmail.com"},
			want: "invalid email format",
		},
		{
			name: "test email validator",
			args: args{1234},
			want: "not a string value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatorEmailFormat(tt.args.value); got != tt.want {
				t.Errorf("ValidatorEmailFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneralValidator(t *testing.T) {
	type args struct {
		message            string
		validationFunction func(any) bool
		value              any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test generic invalid",
			args: args{"generic error", func(v any) bool { return false }, 1},
			want: "generic error",
		},
		{
			name: "test generic valid",
			args: args{"generic error", func(v any) bool { return true }, 1},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeneralValidator(tt.args.message, tt.args.validationFunction)(tt.args.value); got != tt.want {
				t.Errorf("GeneralValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
