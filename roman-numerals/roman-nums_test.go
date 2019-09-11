package roman

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets contverted to I", 1, "I"},
		{"2 gets convdrted to II", 2, "II"},
		{"3 gets convdrted to III", 3, "III"},
		{"4 gets convdrted to IV (can't repeat more than 3 time))", 4, "IV"},
		{"5 gets convdrted to V )", 5, "V"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
