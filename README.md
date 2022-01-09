# prettyout

## Installation

```shell
go get -u github.com/walkersumida/prettyout
```

## Usage

```go
package jaker

import (
	"testing"

	"github.com/walkersumida/prettyout" // <-
)

func TestProfile(t *testing.T) {
	expected := "Walker Sumida"

	t.Run("english fullname", func(t *testing.T) {
		got := "ウォーカー スミダ"
		if got != expected {
			t.Errorf(prettyout.Serror(expected, got)) // <-
		}
	})
}
```

```
=== RUN   TestProfile/english_fullname
    jaker_test.go:15:
        -expected +got:
          string(
        - 	"Walker Sumida",
        + 	"ウォーカー スミダ",
          )
--- FAIL: TestProfile (0.00s)
    --- FAIL: TestProfile/english_fullname (0.00s)
```