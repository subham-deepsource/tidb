// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package cascades

import (
	plannercore "github.com/pingcap/tidb/planner/core"
	"github.com/pingcap/tidb/planner/memo"
	"github.com/pingcap/tidb/planner/transformation"
	"github.com/pingcap/tidb/sessionctx"
)

// Transformation defines the interface for the transformation rules.
type Transformation interface {
	GetPattern() *memo.Pattern
	Match(expr *memo.ExprIter) (matched bool)
	OnTransform(sctx sessionctx.Context, old *memo.ExprIter) (new *memo.GroupExpr, eraseOld bool, err error)
}

// GetTransformationRules gets the all the candidate transformation rules based
// on the logical plan node.
func GetTransformationRules(node plannercore.LogicalPlan) []Transformation {
	return transformationMap[memo.GetOperand(node)]
}

var transformationMap = map[memo.Operand][]Transformation{
	memo.OperandSelection: {
		transformation.NewFilterAggregateTransposeRule(),
	},
	memo.OperandProjection: nil,
}
