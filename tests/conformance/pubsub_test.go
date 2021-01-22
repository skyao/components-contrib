// +build conftests

package conformance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPubsubConformance(t *testing.T) {
	tc, err := NewTestConfiguration("../config/pubsub/tests.yml")
	assert.NoError(t, err)
	assert.NotNil(t, tc)
	tc.Run(t)
}
