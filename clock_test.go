package emoji

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{
			name: "MST",
			t:    must("15:04:05 MST"),
			want: "🕒",
		},
		{
			name: "UTC",
			t:    must("15:04:05 UTC"),
			want: "🕒",
		},

		{
			name: "boundary-0",
			t:    must("00:04:05 EDT"),
			want: "🕛",
		},
		{
			name: "boundary-1",
			t:    must("01:04:05 EDT"),
			want: "🕐",
		},
		{
			name: "boundary-12",
			t:    must("12:04:05 EDT"),
			want: "🕛",
		},
		{
			name: "boundary-23",
			t:    must("23:59:59 EDT"),
			want: "🕚",
		},

		// {
		// 	name: "EEST",
		// 	t:    must("15:04:05 EEST"),
		// 	want: "🕒",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.t.String())

			got := Clock(tt.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

func must(value string) time.Time {
	const layout = "15:04:05 MST"

	ts, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}

	return ts
}
