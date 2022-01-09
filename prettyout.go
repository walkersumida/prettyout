package prettyout

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Serror(expected, got interface{}, opts ...cmp.Option) string {
	if diff := cmp.Diff(expected, got, opts...); diff != "" {
		return fmt.Sprintf("\n-expected +got: \n%s", diff)
	}
	return ""
}
