// +build integrate

package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/neuronlabs/neuron-plugins/repository/postgres/internal"
	"github.com/neuronlabs/neuron-plugins/repository/postgres/migrate"
	"github.com/neuronlabs/neuron-plugins/repository/postgres/tests"
	"github.com/neuronlabs/neuron/mapping"
	"github.com/neuronlabs/neuron/orm"
)

// TestIntegrationDelete integration tests for the Delete processes.
func TestIntegrationDelete(t *testing.T) {
	c := testingController(t, true, testModels...)
	p := testingRepository(c)

	ctx := context.Background()

	mStruct, err := c.ModelStruct(&tests.SimpleModel{})
	require.NoError(t, err)

	defer func() {
		table, err := migrate.ModelsTable(mStruct)
		require.NoError(t, err)
		_ = internal.DropTables(ctx, p.ConnPool, table.Name, table.Schema)
	}()

	db := orm.New(c)
	newModel := func() *tests.SimpleModel {
		return &tests.SimpleModel{
			Attr: "Something",
		}
	}
	t.Run("WithFilter", func(t *testing.T) {
		model := newModel()
		model2 := newModel()
		models := []*tests.SimpleModel{model, model2}
		// Insert models.
		err = tests.SimpleModels.Query(db, models...).Insert()
		require.NoError(t, err)

		assert.Len(t, models, 2)

		affected, err := tests.SimpleModels.Query(db).Where("ID IN", model.ID, model2.ID).Delete()
		require.NoError(t, err)

		assert.Equal(t, int64(2), affected)
	})

	t.Run("Models", func(t *testing.T) {
		model := newModel()
		model2 := newModel()
		// Insert models.
		err = tests.SimpleModels.Query(db, model, model2).Insert()
		require.NoError(t, err)

		affected, err := tests.SimpleModels.Query(db, model, model2).Delete()
		require.NoError(t, err)

		assert.Equal(t, int64(2), affected)
	})
}

func TestSoftDelete(t *testing.T) {
	c := testingController(t, true, testModels...)
	p := testingRepository(c)

	ctx := context.Background()

	mStruct, err := c.ModelStruct(&tests.Model{})
	require.NoError(t, err)

	defer func() {
		table, err := migrate.ModelsTable(mStruct)
		require.NoError(t, err)
		_ = internal.DropTables(ctx, p.ConnPool, table.Name, table.Schema)
	}()

	db := orm.New(c)

	newModel := func() *tests.Model {
		return &tests.Model{
			AttrString: "Something",
		}
	}
	model := newModel()
	model2 := newModel()
	models := []mapping.Model{model, model2}
	// Insert models.
	err = db.Query(mStruct, models...).Insert()
	require.NoError(t, err)

	affected, err := db.Query(mStruct, model).Delete()
	require.NoError(t, err)

	assert.Equal(t, int64(1), affected)

	res, err := db.Query(mStruct).Where("ID = ", model.ID).Find()
	require.NoError(t, err)

	assert.Len(t, res, 0)
}
