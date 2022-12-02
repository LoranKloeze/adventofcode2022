// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import "testing"

func TestRunExistingChallenge(t *testing.T) {

	called := false
	callMe := func() {
		called = true
	}

	challenges := map[string]func(){
		"day1-1": callMe,
	}

	runChallenge("day1-1", challenges)

	if !called {
		t.Errorf("Expected callMe to be called, but it wasn't")
	}

}

func TestRunMissingChallenge(t *testing.T) {
	challenges := map[string]func(){
		"day1-1": func() {},
	}

	err := runChallenge("day3-5", challenges)

	if err == nil {
		t.Errorf("Expected error running a missing challenge, but got none")
	}
}
