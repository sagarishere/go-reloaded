package main

import "testing"

func TestBin2Dec(t *testing.T) {
	tests := []struct {
		binary string
		expect int
	}{
		{"0", 0},
		{"1", 1},
		{"10", 2},
		{"101", 5},
		{"1101", 13},
		{"11111111", 255},
	}

	for _, test := range tests {
		result := Bin2Dec(test.binary)
		if result != test.expect {
			t.Errorf("Bin2Dec(%q) = %d; want %d", test.binary, result, test.expect)
		}
	}
}

func TestHex2Dec(t *testing.T) {
	tests := []struct {
		hex    string
		expect int
	}{
		{"0", 0},
		{"1", 1},
		{"A", 10},
		{"10", 16},
		{"1A", 26},
		{"FF", 255},
	}

	for _, test := range tests {
		result := Hex2Dec(test.hex)
		if result != test.expect {
			t.Errorf("Hex2Dec(%q) = %d; want %d", test.hex, result, test.expect)
		}
	}
}

func TestCap(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"hello", "Hello"},
		{"h", "H"},
		{"", ""},
		{"Hello", "Hello"},
	}

	for _, test := range tests {
		str := test.input
		Cap(&str)
		if str != test.expect {
			t.Errorf("Cap(%q) = %q; want %q", test.input, str, test.expect)
		}
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"HELLO", "hello"},
		{"Hello", "hello"},
		{"", ""},
	}

	for _, test := range tests {
		str := test.input
		Lower(&str)
		if str != test.expect {
			t.Errorf("Lower(%q) = %q; want %q", test.input, str, test.expect)
		}
	}
}

func TestUpper(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"hello", "HELLO"},
		{"Hello", "HELLO"},
		{"", ""},
	}

	for _, test := range tests {
		str := test.input
		Upper(&str)
		if str != test.expect {
			t.Errorf("Upper(%q) = %q; want %q", test.input, str, test.expect)
		}
	}
}

func TestExtractDigit(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"2)", 2},
		{"3)", 3},
		{"10)", 10},
	}

	for _, test := range tests {
		result := extractDigit(test.input)
		if result != test.expect {
			t.Errorf("extractDigit(%q) = %d; want %d", test.input, result, test.expect)
		}
	}
}

func TestProcessText(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "hello world",
			expect: "hello world",
		},
		{
			input:  "A apple",
			expect: "An apple",
		},
		{
			input:  "a orange",
			expect: "an orange",
		},
		{
			input:  "She said ' hello ' world",
			expect: "She said 'hello' world",
		},
		{
			input:  "101 (bin)",
			expect: "5",
		},
		{
			input:  "1A (hex)",
			expect: "26",
		},
		{
			input:  "hello (cap)",
			expect: "Hello",
		},
		{
			input:  "HELLO (low)",
			expect: "hello",
		},
		{
			input:  "hello (up)",
			expect: "HELLO",
		},
		{
			input:  "this is a test (cap, 2)",
			expect: "this is A Test",
		},
		{
			input:  "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			expect: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			input:  "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			expect: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			input:  "Don not be sad ,because sad backwards is das . And das not good",
			expect: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			input:  "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			expect: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}

	for _, test := range tests {
		result := processText(test.input)
		if result != test.expect {
			t.Errorf("processText(%q) = \n%q;\n want\n %q", test.input, result, test.expect)
		}
	}
}
