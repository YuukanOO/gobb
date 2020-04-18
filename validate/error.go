package validate

import (
	"fmt"
	"sort"
	"strings"
)

// Errors maps field names to a validation error string.
type Errors map[string]string

func (e *Errors) Error() string {
	// Let's concatenate field errors to make it easier to read when outputting
	// this error.
	var builder strings.Builder

	// And sort field by names
	keys := make([]string, 0, len(*e))

	for field := range *e {
		keys = append(keys, field)
	}

	sort.Strings(keys)

	for i, key := range keys {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("%s: %s", key, (*e)[key]))
	}

	return builder.String()
}
