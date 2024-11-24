package split

import "log"

func SplitMessage(data []byte, chunkSize int) [][]byte {
	if len(data) == 0 {
		log.Println("Data is empty, nothing to split")
		return nil
	}

	chunks := make([][]byte, 0, (len(data)+chunkSize-1)/chunkSize)
	for start := 0; start < len(data); start += chunkSize {
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[start:end])
	}
	log.Printf("Split into %d chunks", len(chunks))
	return chunks
}
