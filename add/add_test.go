package add

import "testing"

func TestTwoBigFloat(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Fail. Require param 01", args: args{a: "", b: "1"}, want: ""},
		{name: "Fail. Require param 02", args: args{a: "1", b: ""}, want: ""},
		{name: "Fail. Require params", args: args{a: "", b: ""}, want: ""},
		{name: "Fail. Invalid param 01 (alphabet)", args: args{a: "a.000", b: "0"}, want: ""},
		{name: "Fail. Invalid param 02 (alphabet)", args: args{a: "0.12", b: "b.84560"}, want: ""},
		{name: "Fail. Invalid 2 params (alphabet)", args: args{a: "a", b: "b"}, want: ""},

		{name: "Fail. Invalid param 01 (alphabet included)", args: args{a: "12a234.000", b: "0"}, want: ""},
		{name: "Fail. Invalid param 02 (alphabet included)", args: args{a: "0", b: "43123.1$20"}, want: ""},
		{name: "Fail. Invalid 2 params (alphabet included)", args: args{a: "67a565.75i5", b: "5656b565.U88*8"}, want: ""},

		{name: "Success. Two positive numbers (int format)", args: args{a: "987654321", b: "12345679"}, want: "1000000000.00000"},
		{name: "Success. Two positive numbers", args: args{a: "999999.00", b: "1.00"}, want: "1000000.00000"},
		{name: "Success. Two positive numbers. roundPrec 5", args: args{a: "99.56789", b: "0.432111111"}, want: "100.00000"},
		{name: "Success. Two positive numbers, very small value", args: args{a: "0.12345111111110001", b: "0.12345"}, want: "0.24690"},
		{name: "Success. Two positive numbers, very very small value - case ceil", args: args{a: "0.000010001", b: "0.02598300"}, want: "0.02599"},
		{name: "Success. Two positive numbers, very very small value - case floor", args: args{a: "0.000010001", b: "0.02598500"}, want: "0.02600"},

		{name: "Success. One positive, one negative (int format)", args: args{a: "9999999999999", b: "-1"}, want: "9999999999998.00000"},
		{name: "Success. One positive, one negative. roundPrec 5", args: args{a: "999.25864", b: "-9.25863"}, want: "990.00001"},

		{name: "Success. Two negative numbers (int format)", args: args{a: "-15", b: "-5"}, want: "-20.00000"},
		{name: "Success. Two negative numbers", args: args{a: "-15.00", b: "-5.0"}, want: "-20.00000"},
		{name: "Success. Two negative numbers roundPrec 5", args: args{a: "-100005.02568", b: "-5.06432"}, want: "-100010.09000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoBigFloat(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("TwoBigFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
