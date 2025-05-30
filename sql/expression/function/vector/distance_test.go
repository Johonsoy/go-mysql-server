// Copyright 2024 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vector

import (
	"testing"

	assert "github.com/stretchr/testify/require"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/dolthub/go-mysql-server/sql/types/jsontests"
)

func jsonExpression(t *testing.T, val interface{}) sql.Expression {
	return expression.NewLiteral(jsontests.ConvertToJson(t, val), types.JSON)
}

func TestDistance(t *testing.T) {
	ctx := sql.NewEmptyContext()
	distance := NewDistance(DistanceL2Squared{}, jsonExpression(t, "[0.0, 0.0]"), jsonExpression(t, "[3.0, 4.0]"))
	result, err := distance.Eval(ctx, nil)
	assert.NoError(t, err)
	assert.InEpsilon(t, 25.0, result, 0.1)
}
