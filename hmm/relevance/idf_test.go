package relevance_test

import (
	"fmt"
	"testing"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/extracker"

	"github.com/vcaesar/tt"
)

var (
	segs, _ = gse.New()
	text    = "那里湖面总是澄清, 那里空气充满宁静"
)

func TestExtAndRank(t *testing.T) {
	var te extracker.TagExtracter
	te.WithGse(segs)
	err := te.LoadIDF()
	tt.Nil(t, err)

	segments := te.ExtractTags(text, 5)
	fmt.Println("segments: ", len(segments), segments)

	var tr extracker.TextRanker
	tr.WithGse(segs)

	results := tr.TextRank(text, 5)
	fmt.Println("results: ", results)
}