package main

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct{
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English":{
			lang:"en",
			want:"Hello World",
		},
		"French":{
			lang:"fr",
			want:"Bonjour le Monde",
		},
		"Greek":{
			lang:"el",
			want:"Χαίρετε κόσμε",
		},
		"Swahili":{
			lang:"sw",
			want:"Habari dunia",
		},
		"Akkadian": {
			lang: "akk",
			want: `Unsupported language: "akk" `,
		},
		"Empty": {
			lang:"",
			want: `Unsupported language: ""`,
		},
	}

	for name, tc := range tests {
		t.Run(name,func(t *testing.T){
			got := greet(tc.lang)

			if got != tc.want{
				t.Errorf("expected %q, got: %q",tc.want,got)
			}
		})
	}
}