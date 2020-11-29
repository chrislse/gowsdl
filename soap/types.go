package soap

import (
	"time"
)

//Time extended time parsing types
type SafeTime time.Time

func (t *SafeTime) UnmarshalText(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t1, err := time.Parse(time.RFC3339, string(data))
	if err != nil {
		t = &SafeTime{}
		return nil
	}
	*t = SafeTime(t1)
	return err
}
