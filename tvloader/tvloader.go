// Package tvloader is used to load and parse SPDX tag-value documents
// into tools-golang data structures.
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package tvloader

import (
	"io"

	"github.com/codenotary/spdx-tools/spdx/v2_1"
	"github.com/codenotary/spdx-tools/spdx/v2_2"
	"github.com/codenotary/spdx-tools/tvloader/parser2v1"
	"github.com/codenotary/spdx-tools/tvloader/parser2v2"
	"github.com/codenotary/spdx-tools/tvloader/reader"
)

// Load2_1 takes an io.Reader and returns a fully-parsed SPDX Document
// (version 2.1) if parseable, or error if any error is encountered.
func Load2_1(content io.Reader) (*v2_1.Document, error) {
	tvPairs, err := reader.ReadTagValues(content)
	if err != nil {
		return nil, err
	}

	doc, err := parser2v1.ParseTagValues(tvPairs)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Load2_2 takes an io.Reader and returns a fully-parsed SPDX Document
// (version 2.2) if parseable, or error if any error is encountered.
func Load2_2(content io.Reader) (*v2_2.Document, error) {
	tvPairs, err := reader.ReadTagValues(content)
	if err != nil {
		return nil, err
	}

	doc, err := parser2v2.ParseTagValues(tvPairs)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
