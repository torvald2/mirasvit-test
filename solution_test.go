package mirasvittest

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type Tests []struct {
	Solution int               `json:"solution"`
	Data     []map[string]bool `json:"data"`
}

func TestSolution(t *testing.T) {
	//55
	payload, err := ioutil.ReadFile("solutions.json")
	if err != nil {
		t.Fatal(err)
		return
	}
	var testData Tests
	if err := json.Unmarshal(payload, &testData); err != nil {
		t.Fatal(err)
		return
	}

	for _, test := range testData {
		sol := FindSolution(test.Data)
		if sol != test.Solution {
			t.Errorf("Solution not found: GOT %v Expect %v", sol, test.Solution)
		}
	}

}

func BenchmarkSolution(b *testing.B) {
	var blocks = []map[string]bool{
		{
			"school": true,
			"gym":    false,
			"store":  false,
		},
		{
			"school": false,
			"gym":    true,
			"store":  false,
		},
		{
			"school": true,
			"gym":    true,
			"store":  false,
		},
		{ // max distance: 1
			"school": true,
			"gym":    false,
			"store":  false,
		},
		{
			"school": true,
			"gym":    false,
			"store":  true,
		},
	}

	for i := 0; i < b.N; i++ {
		FindSolution(blocks)
	}
}

//BenchmarkSolution-8   	 1267094	       918.5 ns/op	     288 B/op	       9 allocs/op Start
//BenchmarkSolution-8   	 1326927	       893.9 ns/op	     264 B/op	       8 allocs/op Iter 1
//BenchmarkSolution-8   	 1409508	       804.9 ns/op	     168 B/op	       4 allocs/op
