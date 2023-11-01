package extracker

import (
	"sort"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/relevance"
)

// Segment type a word with weight.
type Segment struct {
	text   string
	weight float64
}

// Text return the segment's text.
func (s Segment) Text() string {
	return s.text
}

// Weight return the segment's weight.
func (s Segment) Weight() float64 {
	return s.weight
}

// Segments type a slice of Segment.
type Segments []Segment

func (ss Segments) Len() int {
	return len(ss)
}

func (ss Segments) Less(i, j int) bool {
	if ss[i].weight == ss[j].weight {
		return ss[i].text < ss[j].text
	}

	return ss[i].weight < ss[j].weight
}

func (ss Segments) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

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
func (t *TagExtracter) ExtractTags(text string, topK int) (tags Segments) {

	ws := make(Segments, 0)
	for k, v := range t.Relevance.GetFreqMap(text) {
		ws = append(ws, Segment{text: k, weight: t.Relevance.CalculateWeight(k, v)})
	}

	sort.Sort(sort.Reverse(ws))

	if len(ws) > topK {
		tags = ws[:topK]
		return
	}

	tags = ws
	return
}
