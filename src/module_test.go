package src

import (
	"testing"
)

func TestModuleName(t *testing.T) {
	if ProjectName() != "days-since" {
		t.Errorf("Project name `%s` incorrect", ProjectName())
	}
}
