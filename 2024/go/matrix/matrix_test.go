package matrix

import "testing"

func TestMatrixNew(t *testing.T) {
	m, err := New[rune](1, 1)

	if err != nil {
		t.Errorf("Expected err to be nil, got %q", err)
	}

	if m.width != 1 {
		t.Errorf("Expected width to equal %q, got %q", 1, m.width)
	}

	if m.height != 1 {
		t.Errorf("Expected height to equal %q, got %q", 1, m.height)
	}
}
