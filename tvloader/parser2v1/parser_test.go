// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package parser2v1

import (
	"testing"

	"github.com/codenotary/spdx-tools/tvloader/reader"
)

// ===== Parser exported entry point tests =====
func TestParser2_1CanParseTagValues(t *testing.T) {
	var tvPairs []reader.TagValuePair

	// create some pairs
	tvPair1 := reader.TagValuePair{Tag: "SPDXVersion", Value: "SPDX-2.1"}
	tvPairs = append(tvPairs, tvPair1)
	tvPair2 := reader.TagValuePair{Tag: "DataLicense", Value: "CC0-1.0"}
	tvPairs = append(tvPairs, tvPair2)
	tvPair3 := reader.TagValuePair{Tag: "SPDXID", Value: "SPDXRef-DOCUMENT"}
	tvPairs = append(tvPairs, tvPair3)

	// now parse them
	doc, err := ParseTagValues(tvPairs)
	if err != nil {
		t.Errorf("got error when calling ParseTagValues: %v", err)
	}
	if doc.SPDXVersion != "SPDX-2.1" {
		t.Errorf("expected SPDXVersion to be SPDX-2.1, got %v", doc.SPDXVersion)
	}
	if doc.DataLicense != "CC0-1.0" {
		t.Errorf("expected DataLicense to be CC0-1.0, got %v", doc.DataLicense)
	}
	if doc.SPDXIdentifier != "DOCUMENT" {
		t.Errorf("expected SPDXIdentifier to be DOCUMENT, got %v", doc.SPDXIdentifier)
	}

}

// ===== Parser initialization tests =====
func TestParser2_1InitCreatesResetStatus(t *testing.T) {
	parser := tvParser2_1{}
	if parser.st != psStart2_1 {
		t.Errorf("parser did not begin in start state")
	}
	if parser.doc != nil {
		t.Errorf("parser did not begin with nil document")
	}
}

func TestParser2_1HasDocumentAfterCallToParseFirstTag(t *testing.T) {
	parser := tvParser2_1{}
	err := parser.parsePair2_1("SPDXVersion", "SPDX-2.1")
	if err != nil {
		t.Errorf("got error when calling parsePair2_1: %v", err)
	}
	if parser.doc == nil {
		t.Errorf("doc is still nil after parsing first pair")
	}
}

func TestParser2_1StartFailsToParseIfInInvalidState(t *testing.T) {
	parser := tvParser2_1{st: psReview2_1}
	err := parser.parsePairFromStart2_1("SPDXVersion", "SPDX-2.1")
	if err == nil {
		t.Errorf("expected non-nil error, got nil")
	}
}

func TestParser2_1FilesWithoutSpdxIdThrowErrorAtCompleteParse(t *testing.T) {
	// case: checks the last file
	// Last unpackaged file no packages in doc
	// Last file of last package in the doc
	tvPairs := []reader.TagValuePair{
		{Tag: "SPDXVersion", Value: "SPDX-2.1"},
		{Tag: "DataLicense", Value: "CC0-1.0"},
		{Tag: "SPDXID", Value: "SPDXRef-DOCUMENT"},
		{Tag: "FileName", Value: "f1"},
	}
	_, err := ParseTagValues(tvPairs)
	if err == nil {
		t.Errorf("file without SPDX Identifier getting accepted")
	}
}

func TestParser2_1PackageWithoutSpdxIdThrowErrorAtCompleteParse(t *testing.T) {
	// case: Checks the last package
	tvPairs := []reader.TagValuePair{
		{Tag: "SPDXVersion", Value: "SPDX-2.1"},
		{Tag: "DataLicense", Value: "CC0-1.0"},
		{Tag: "SPDXID", Value: "SPDXRef-DOCUMENT"},
		{Tag: "PackageName", Value: "p1"},
	}
	_, err := ParseTagValues(tvPairs)
	if err == nil {
		t.Errorf("package without SPDX Identifier getting accepted")
	}
}
