package gotest

import "testing"

type RunStrategy func(string, func(TWrapper))

func GetRealRunStrategy(t TWrapper) RunStrategy {
	return func(name string, testBody func(TWrapper)) {
		t.(*testing.T).Run(name, func(tt *testing.T) {
			testBody(tt)
		})
	}
}
