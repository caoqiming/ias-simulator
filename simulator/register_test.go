package simulator

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMemoryAddressRegister(t *testing.T) {
	test := []struct {
		want int
	}{
		{
			want: 10,
		},
		{
			want: 24,
		},
		{
			want: 256,
		},
		{
			want: 4095,
		},
		{
			want: 2024,
		},
		{
			want: 2025,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i),
			func(t *testing.T) {
				MAR.SetAddr(tt.want)
				// fmt.Printf("tt.want %d,MAR.data %v,MAR.data[0] %v,MAR.data[1] %v,\n", tt.want, MAR.data, MAR.data[0], MAR.data[1])
				got := MAR.GetAddr()
				assert.Equal(t, got, tt.want)
			})
	}
}
