package split

import (
	"strings"
	"testing"
)

func TestSplitMessage(t *testing.T) {
	data := []byte(strings.Repeat("a", 2048)) // Создаём 2048 байт данных
	chunks := SplitMessage(data, 512)

	if len(chunks) != 4 {
		t.Errorf("Expected 4 chunks, got %d", len(chunks))
	}

	for _, chunk := range chunks {
		if len(chunk) > 512 {
			t.Errorf("Chunk size exceeded limit: %d bytes", len(chunk))
		}
	}
}
