/* *
 * Test Model
 */

package model

import (
	"context"
	"fmt"

	"github.com/extra-time-zone/xgin/example/model/dao"
	"github.com/extra-time-zone/xgin/xerror"
)

type TestModel struct {
	ctx context.Context
	dao *dao.TestDao
}

func NewTestModel(ctx context.Context) *TestModel {
	return &TestModel{
		ctx: ctx,
		dao: dao.NewTestDao(ctx),
	}
}

func (m *TestModel) GetData(x int) (string, xerror.Error) {
	data, xerr := m.dao.GetData(x)
	if xerr != nil {
		return "", xerror.Wrap(xerr, "test-model-error")
	}

	return fmt.Sprintf("%s_%s", data, "test-model"), nil
}
