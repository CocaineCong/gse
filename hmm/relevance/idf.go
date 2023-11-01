// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package relevance

import (
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/stop_word"
)

// IDF type a dictionary for all words with the
// IDFs(Inverse Document Frequency).
type IDF struct {
	// median of word frequencies for calculate the weight of backup
	median float64

	// the list of word frequencies
	freqs []float64

	Base
}

// NewIDF create a new Relevance to implementing IDF
func NewIDF() Relevance {
	idf := &IDF{
		freqs: make([]float64, 0),
	}

	idf.StopWord = stop_word.NewStopWord()

	return Relevance(idf)
}

// AddToken add a new word with IDF into the dictionary.
func (i *IDF) AddToken(text string, freq float64, pos ...string) error {
	err := i.Seg.AddToken(text, freq, pos...)

	i.freqs = append(i.freqs, freq)
	sort.Float64s(i.freqs)
	i.median = i.freqs[len(i.freqs)/2]
	return err
}

// LoadDict load the idf dictionary
func (i *IDF) LoadDict(files ...string) error {
	if len(files) <= 0 {
		files = i.Seg.GetIdfPath(files...)
	}

	return i.Seg.LoadDict(files...)
}

// Freq return the IDF of the word
func (i *IDF) Freq(key string) (float64, string, bool) {
	return i.Seg.Find(key)
}

// NumTokens return the IDF tokens' num
func (i *IDF) NumTokens() int {
	return i.Seg.Dict.NumTokens()
}

// TotalFreq return the IDF total frequency
func (i *IDF) TotalFreq() float64 {
	return i.Seg.Dict.TotalFreq()
}

// GetFreqMap return the IDF freq map
func (i *IDF) GetFreqMap(text string) map[string]float64 {
	freqMap := make(map[string]float64)

	for _, w := range i.Seg.Cut(text, true) {
		w = strings.TrimSpace(w)
		if utf8.RuneCountInString(w) < 2 {
			continue
		}
		if i.StopWord.IsStopWord(w) {
			continue
		}

		if f, ok := freqMap[w]; ok {
			freqMap[w] = f + 1.0
		} else {
			freqMap[w] = 1.0
		}
	}

	total := 0.0
	for _, freq := range freqMap {
		total += freq
	}

	for k, v := range freqMap {
		freqMap[k] = v / total
	}

	return freqMap
}

// CalculateWeight calculate the word's weight by IDF
func (i *IDF) CalculateWeight(k string, v float64) float64 {
	if freq, _, ok := i.Freq(k); ok {
		return freq * v
	}

	return i.median * v
}

// GetSeg get IDF Segmenter
func (i *IDF) GetSeg() gse.Segmenter {
	return i.Seg
}

// LoadDictStr load dict for IDF seg
func (i *IDF) LoadDictStr(dictStr string) error {
	return i.Seg.LoadDictStr(dictStr)
}

// LoadStopWord load stop word for IDF
func (i *IDF) LoadStopWord(fileName ...string) error {
	return i.StopWord.LoadDict(fileName...)
}
