package codesandboxsdk

import (
	"fmt"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

type Time struct {
	time.Time
}

// UnmarshalJSON implements the custom unmarshaller for Time without timezone.
func (t *Time) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "null" || s == "" {
		return nil
	}
	parsedTime, err := dateparse.ParseAny(s)
	if err != nil {
		return fmt.Errorf("failed to parse time %q: %w", s, err)
	}

	t.Time = parsedTime
	return nil
}
