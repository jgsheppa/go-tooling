package counter

import (
	"bytes"
	"testing"
)

func TestNextCounterToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}


	cases := []struct {
		description string
		current int
		iterations int
		want int
	}{
		{
			description: "should equal 5",
			current: 0,
			iterations: 5,
			want: 5,
		},
		{
			description: "should equal 2",
			current: 1,
			iterations: 2,
			want: 2,
		},
	}

	for _, tt := range cases {
		c := &Counter{
			Current: tt.current,
			Output:  fakeTerminal,
		}

        t.Run(tt.description, func(t *testing.T){
			for i := tt.current; i < tt.iterations; i++ {
				c.Next()
			}
            
            if c.Current != tt.want {
                t.Errorf("want %d, but got %d", tt.want, c.Current)
            }
        })
    }
}
