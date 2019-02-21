package slugify

import "testing"

func TestSlugify(t *testing.T) {
	var tests = []struct{ in, out string }{
		{"simple test", "simple-test"},
		{"Simple Test", "simple-test"},
		{"I'm go developer", "i-m-go-developer"},
		{"Simples código em go", "simples-codigo-em-go"},
		{"日本語の手紙をテスト", "ri-ben-yu-noshou-zhi-wotesuto"},
		{"北京kožuščekł", "bei-jing-kozuscekl"},
	}

	for _, test := range tests {
		if out := Slugify(test.in); out != test.out {
			t.Errorf("%q: %q != %q", test.in, out, test.out)
		}
	}
}

func TestIDify(t *testing.T) {
	var tests = []struct{ in, out string }{
		{"simple test", "simple-test"},
		{"Simple Test", "Simple-Test"},
		{"I'm go developer", "I-m-go-developer"},
		{"Simples código em go", "Simples-codigo-em-go"},
		{"日本語の手紙をテスト", "Ri-Ben-Yu-noShou-Zhi-wotesuto"},
		{"北京kožuščekł", "Bei-Jing-kozuscekl"},
	}

	for _, test := range tests {
		if out := IDify(test.in); out != test.out {
			t.Errorf("%q: %q != %q", test.in, out, test.out)
		}
	}
}