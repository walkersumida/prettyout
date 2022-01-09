package prettyout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Error(t *testing.T, expected, got interface{}, opts ...cmp.Option) {
	if diff := cmp.Diff(expected, got, opts...); diff != "" {
		t.Errorf("-expected +got: %s", diff)
	}
}
