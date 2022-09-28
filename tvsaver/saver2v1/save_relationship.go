// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v1

import (
	"fmt"
	"io"

	"github.com/codenotary/spdx-tools/spdx/common"
	"github.com/codenotary/spdx-tools/spdx/v2_1"
)

func renderRelationship2_1(rln *v2_1.Relationship, w io.Writer) error {
	rlnAStr := common.RenderDocElementID(rln.RefA)
	rlnBStr := common.RenderDocElementID(rln.RefB)
	if rlnAStr != "SPDXRef-" && rlnBStr != "SPDXRef-" && rln.Relationship != "" {
		fmt.Fprintf(w, "Relationship: %s %s %s\n", rlnAStr, rln.Relationship, rlnBStr)
	}
	if rln.RelationshipComment != "" {
		fmt.Fprintf(w, "RelationshipComment: %s\n", textify(rln.RelationshipComment))
	}

	return nil
}
