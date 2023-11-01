package relevance

import (
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/segment"
	"github.com/go-ego/gse/hmm/stop_word"
)

// Relevance easily scalable Relevance calculations (for idf, tf-idf, bm25 and so on)
type Relevance interface {
	// AddToken add text, frequency, position on obj
	AddToken(text string, freq float64, pos ...string) error

	// LoadDict load file from incoming parameters,
	// if incoming params no exist, will load file from default file path
	LoadDict(files ...string) error

	// LoadDictStr loading dict file by file path
	LoadDictStr(pathStr string) error

	// LoadStopWord loading word file by filename
	LoadStopWord(fileName ...string) error

	// Freq find the frequency, position, existence information of the key
	Freq(key string) (float64, string, bool)

	// NumTokens  the number of tokens in the dictionary
	NumTokens() int

	// TotalFreq the total number of tokens in the dictionary
	TotalFreq() float64

	// FreqMap get frequency map
	// key: word, value: frequency
	FreqMap(text string) map[string]float64

	// GetSeg Get the segmenter of Relevance algorithms
	GetSeg() gse.Segmenter

	// ConstructSeg return the segment with weight
	ConstructSeg(text string) segment.Segments
}

type Base struct {
	// loading some stop words
	StopWord *stop_word.StopWord

	// loading segmenter for cut word
	Seg gse.Segmenter
}
