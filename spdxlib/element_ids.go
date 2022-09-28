// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdxlib

import (
	"sort"

	"github.com/codenotary/spdx-tools/spdx"
)

// SortElementIDs sorts and returns the given slice of ElementIDs
func SortElementIDs(eIDs []spdx.ElementID) []spdx.ElementID {
	sort.Slice(eIDs, func(i, j int) bool {
		return eIDs[i] < eIDs[j]
	})

	return eIDs
}
