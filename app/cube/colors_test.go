package cube

import (
	"testing"
)

func TestGetColor(t *testing.T) {
	for i := 0; i < 6; i++ {
		color := GetColor(i)
		if color.number != i {
			t.Errorf("Expected color number %d, got %d", i, color.number)
		}
	}
}
