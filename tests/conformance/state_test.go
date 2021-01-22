// +build conftests

package conformance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateConformance(t *testing.T) {
	tc, err := NewTestConfiguration("../config/state/tests.yml")
	assert.NoError(t, err)
	assert.NotNil(t, tc)
	errs := tc.Run(t)
	if len(errs) != 0 {
		for _, err = range errs {
			t.Log(err)
		}
		assert.Fail(t, "some of tests failed")
	}
}