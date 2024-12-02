package main

import (
	"fmt"
	"strings"
	"testing"
)

type testpair struct {
	prog       string
	wantResult string
}

func TestExecute(t *testing.T) {
	pairs := []testpair{
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}

	for _, p := range pairs {
		res, _ := execute(format(p.prog))
		isResult := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), ","), "[]")
		if isResult != p.wantResult {
			t.Errorf("Wanted %q but got %q", p.wantResult, isResult)
		}
	}

}
