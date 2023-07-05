package rancher_gatekeeper

import (
	"testing"

	"github.com/rancher/hull/pkg/test"
)

func TestChart(t *testing.T) {
	opts := test.GetRancherOptions()
	opts.Coverage.IncludeSubcharts = true
	opts.Coverage.Disabled = false
	opts.YAMLLint.Enabled = false
	suite.Run(t, opts)
}
