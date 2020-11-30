package soap

import (
	"testing"
	"time"
)

func TestTimeParsingWithoutZ(t *testing.T) {
	st := []byte("2020-11-30T02:22:27.693")

	safeTime := &SafeTime{}

	err := safeTime.UnmarshalText(st)

	if err != nil {
		t.Fatal("expect no error but got", err)
	}

	sTime := time.Time(*safeTime)

	if sTime.Second() == 0 {
		t.Fatal("expect time not equal to zero")
	}
}
