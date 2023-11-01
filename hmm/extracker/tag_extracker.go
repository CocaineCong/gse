package extracker

import (
	"sort"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/relevance"
	"github.com/go-ego/gse/hmm/segment"
)

// TagExtracter is extract tags struct.
type TagExtracter struct {
	seg gse.Segmenter

	// calculate weight by Relevance(including IDF,TF-IDF,BM25 and so on)
	Relevance relevance.Relevance

	// stopWord *StopWord
}

// WithGse register the gse segmenter
func (t *TagExtracter) WithGse(segs gse.Segmenter) {
	// t.stopWord = NewStopWord()
	t.seg = segs
}

// LoadDict load and create a new dictionary from the file
func (t *TagExtracter) LoadDict(fileName ...string) error {
	// t.stopWord = NewStopWord()
	return t.seg.LoadDict(fileName...)
}

// LoadIDF load and create a new IDF dictionary from the file.
func (t *TagExtracter) LoadIDF(fileName ...string) error {
	t.Relevance = relevance.NewIDF()
	return t.Relevance.LoadDict(fileName...)
}

// LoadIDFStr load and create a new IDF dictionary from the string.
func (t *TagExtracter) LoadIDFStr(str string) error {
	t.Relevance = relevance.NewIDF()
	return t.Relevance.LoadDictStr(str)
}

// LoadStopWords load and create a new StopWord dictionary from the file.
func (t *TagExtracter) LoadStopWords(fileName ...string) error {
	return t.Relevance.LoadStopWord(fileName...)
}

// ExtractTags extract the topK keywords from text.
func (t *TagExtracter) ExtractTags(text string, topK int) (tags segment.Segments) {
	if t.Relevance == nil {
		// If no correlation algorithm, we will set the idf for default.
		t.Relevance = relevance.NewIDF()
	}

	// handler text to construct segment with weight
	weighSeg := t.Relevance.ConstructSeg(text)

	// sort by weight desc
	sort.Sort(sort.Reverse(weighSeg))

	// choose the top keywords if length of weightSeg bigger than topK
	if len(weighSeg) > topK {
		tags = weighSeg[:topK]
		return
	}

	tags = weighSeg
	return
}
