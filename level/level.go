package level

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

// sentinal error that can be used for debugging when
// the input value to AsLevel does not find a match.
type ErrInvalidLevel[T ~string] struct {
	OriginalValue  T
	SanitizedValue string
}

func (err ErrInvalidLevel[T]) Error() string {
	return fmt.Sprintf(
		"could not match a level for `%s` to a valid zerolog level",
		err.OriginalValue,
	)
}

// Will take a given string-like value and convert it to
// a zerolog level. The fallback value will be returned if
// none match.
func AsLevelElse[T ~string](value T, fallback zerolog.Level) zerolog.Level {
	level, err := AsLevel(value)
	if err != nil {
		level = fallback
	}
	return level
}

// Will take the given string-like value and convert it to
// a zerolog level. Will return an error if none match.
func AsLevel[T ~string](value T) (zerolog.Level, error) {
	sanitizedValue := strings.TrimSpace(
		strings.ToLower(
			string(value),
		),
	)
	switch sanitizedValue {
	case "debug":
		return zerolog.DebugLevel, nil
	case "info":
		return zerolog.InfoLevel, nil
	case "warn":
		return zerolog.WarnLevel, nil
	case "error":
		return zerolog.ErrorLevel, nil
	case "fatal":
		return zerolog.FatalLevel, nil
	case "no":
		return zerolog.NoLevel, nil
	case "disabled":
		return zerolog.Disabled, nil
	default:
		return zerolog.NoLevel, ErrInvalidLevel[T]{
			OriginalValue:  value,
			SanitizedValue: sanitizedValue,
		}
	}
}
