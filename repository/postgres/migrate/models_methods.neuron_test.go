// Code generated by neuron/generator. DO NOT EDIT.
// This file was generated at:
// Tue, 19 May 2020 10:04:54 +0200

package migrate

import (
	"github.com/neuronlabs/neuron/errors"
	"github.com/neuronlabs/neuron/mapping"
	"time"
)

// Compile time check if Model implements mapping.Model interface.
var _ mapping.Model = &Model{}

// NeuronCollectionName implements mapping.Model interface method.
// Returns the name of the collection for the 'Model'.
func (m *Model) NeuronCollectionName() string {
	return "models"
}

// IsPrimaryKeyZero implements query.Model interface method.
func (m *Model) IsPrimaryKeyZero() bool {
	return m.ID == 0
}

// GetPrimaryKeyValue implements query.Model interface method.
func (m *Model) GetPrimaryKeyValue() interface{} {
	return m.ID
}

// GetPrimaryKeyAddress implements query.Model interface method.
func (m *Model) GetPrimaryKeyAddress() interface{} {
	return &m.ID
}

// GetPrimaryKeyHashableValue implements query.Model interface method.
func (m *Model) GetPrimaryKeyHashableValue() interface{} {
	return m.ID
}

// GetPrimaryKeyZeroValue implements query.Model interface method.
func (m *Model) GetPrimaryKeyZeroValue() interface{} {
	return 0
}

// SetPrimaryKey implements query.Model interface method.
func (m *Model) SetPrimaryKeyValue(value interface{}) error {
	if v, ok := value.(int); ok {
		m.ID = v
		return nil
	}
	// Check alternate types for given field.
	switch valueType := value.(type) {
	case int8:
		m.ID = int(valueType)
	case int16:
		m.ID = int(valueType)
	case int32:
		m.ID = int(valueType)
	case int64:
		m.ID = int(valueType)
	case uint:
		m.ID = int(valueType)
	case uint8:
		m.ID = int(valueType)
	case uint16:
		m.ID = int(valueType)
	case uint32:
		m.ID = int(valueType)
	case uint64:
		m.ID = int(valueType)
	case float32:
		m.ID = int(valueType)
	case float64:
		m.ID = int(valueType)
	default:
		return errors.Newf(mapping.ClassFieldValue, "provided invalid value: '%T' for the primary field for model: 'Model'", value)
	}
	return nil
}

// Compile time check if Model implements mapping.Fielder interface.
var _ mapping.Fielder = &Model{}

// GetFieldsAddress gets the address of provided 'field'.
func (m *Model) GetFieldsAddress(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return &m.ID, nil
	case 1: // Attr
		return &m.Attr, nil
	case 2: // SnakeCased
		return &m.SnakeCased, nil
	case 3: // CreatedAt
		return &m.CreatedAt, nil
	case 4: // UpdatedAt
		return &m.UpdatedAt, nil
	case 5: // DeletedAt
		return &m.DeletedAt, nil
	}
	return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for given model: Model'", field.Name())
}

// GetFieldZeroValue implements mapping.Fielder interface.s
func (m *Model) GetFieldZeroValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return 0, nil
	case 1: // Attr
		return "", nil
	case 2: // SnakeCased
		return "", nil
	case 3: // CreatedAt
		return time.Time{}, nil
	case 4: // UpdatedAt
		return time.Time{}, nil
	case 5: // DeletedAt
		return (*time.Time)(nil), nil
	default:
		return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
	}
}

// IsFieldZero implements mapping.Fielder interface.
func (m *Model) IsFieldZero(field *mapping.StructField) (bool, error) {
	switch field.Index[0] {
	case 0: // ID
		return m.ID == 0, nil
	case 1: // Attr
		return m.Attr == "", nil
	case 2: // SnakeCased
		return m.SnakeCased == "", nil
	case 3: // CreatedAt
		return m.CreatedAt == time.Time{}, nil
	case 4: // UpdatedAt
		return m.UpdatedAt == time.Time{}, nil
	case 5: // DeletedAt
		return m.DeletedAt == nil, nil
	}
	return false, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
}

// SetFieldZeroValue implements mapping.Fielder interface.s
func (m *Model) SetFieldZeroValue(field *mapping.StructField) error {
	switch field.Index[0] {
	case 0: // ID
		m.ID = 0
	case 1: // Attr
		m.Attr = ""
	case 2: // SnakeCased
		m.SnakeCased = ""
	case 3: // CreatedAt
		m.CreatedAt = time.Time{}
	case 4: // UpdatedAt
		m.UpdatedAt = time.Time{}
	case 5: // DeletedAt
		m.DeletedAt = nil
	default:
		return errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
	}
	return nil
}

// GetHashableFieldValue implements mapping.Fielder interface.
func (m *Model) GetHashableFieldValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return m.ID, nil
	case 1: // Attr
		return m.Attr, nil
	case 2: // SnakeCased
		return m.SnakeCased, nil
	case 3: // CreatedAt
		return m.CreatedAt, nil
	case 4: // UpdatedAt
		return m.UpdatedAt, nil
	case 5: // DeletedAt
		if m.DeletedAt == nil {
			return m.DeletedAt, nil
		}
		return *m.DeletedAt, nil
	}
	return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for given model: 'Model'", field.Name())
}

// GetFieldValue implements mapping.Fielder interface.
func (m *Model) GetFieldValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return m.ID, nil
	case 1: // Attr
		return m.Attr, nil
	case 2: // SnakeCased
		return m.SnakeCased, nil
	case 3: // CreatedAt
		return m.CreatedAt, nil
	case 4: // UpdatedAt
		return m.UpdatedAt, nil
	case 5: // DeletedAt
		return m.DeletedAt, nil
	}
	return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for given model: Model'", field.Name())
}

// SetFieldValue implements mapping.Fielder interface.
func (m *Model) SetFieldValue(field *mapping.StructField, value interface{}) (err error) {
	switch field.Index[0] {
	case 0: // ID
		if v, ok := value.(int); ok {
			m.ID = v
			return nil
		}
		switch v := value.(type) {
		case int8:
			m.ID = int(v)
		case int16:
			m.ID = int(v)
		case int32:
			m.ID = int(v)
		case int64:
			m.ID = int(v)
		case uint:
			m.ID = int(v)
		case uint8:
			m.ID = int(v)
		case uint16:
			m.ID = int(v)
		case uint32:
			m.ID = int(v)
		case uint64:
			m.ID = int(v)
		case float32:
			m.ID = int(v)
		case float64:
			m.ID = int(v)
		default:
			return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
		}
		return nil
	case 1: // Attr
		if v, ok := value.(string); ok {
			m.Attr = v
			return nil
		}
		// Check alternate types for the Attr.
		if v, ok := value.([]byte); ok {
			m.Attr = string(v)
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 2: // SnakeCased
		if v, ok := value.(string); ok {
			m.SnakeCased = v
			return nil
		}
		// Check alternate types for the SnakeCased.
		if v, ok := value.([]byte); ok {
			m.SnakeCased = string(v)
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 3: // CreatedAt
		if v, ok := value.(time.Time); ok {
			m.CreatedAt = v
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 4: // UpdatedAt
		if v, ok := value.(time.Time); ok {
			m.UpdatedAt = v
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 5: // DeletedAt
		if value == nil {
			m.DeletedAt = nil
			return nil
		}
		if v, ok := value.(*time.Time); ok {
			m.DeletedAt = v
			return nil
		}
		// Check if it is non-pointer value.
		if v, ok := value.(time.Time); ok {
			m.DeletedAt = &v
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	default:
		return errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for the model: 'Model'", field.Name())
	}
}

// Compile time check if BasicModel implements mapping.Model interface.
var _ mapping.Model = &BasicModel{}

// NeuronCollectionName implements mapping.Model interface method.
// Returns the name of the collection for the 'BasicModel'.
func (b *BasicModel) NeuronCollectionName() string {
	return "basic_models"
}

// IsPrimaryKeyZero implements query.Model interface method.
func (b *BasicModel) IsPrimaryKeyZero() bool {
	return b.ID == 0
}

// GetPrimaryKeyValue implements query.Model interface method.
func (b *BasicModel) GetPrimaryKeyValue() interface{} {
	return b.ID
}

// GetPrimaryKeyAddress implements query.Model interface method.
func (b *BasicModel) GetPrimaryKeyAddress() interface{} {
	return &b.ID
}

// GetPrimaryKeyHashableValue implements query.Model interface method.
func (b *BasicModel) GetPrimaryKeyHashableValue() interface{} {
	return b.ID
}

// GetPrimaryKeyZeroValue implements query.Model interface method.
func (b *BasicModel) GetPrimaryKeyZeroValue() interface{} {
	return 0
}

// SetPrimaryKey implements query.Model interface method.
func (b *BasicModel) SetPrimaryKeyValue(value interface{}) error {
	if v, ok := value.(int); ok {
		b.ID = v
		return nil
	}
	// Check alternate types for given field.
	switch valueType := value.(type) {
	case int8:
		b.ID = int(valueType)
	case int16:
		b.ID = int(valueType)
	case int32:
		b.ID = int(valueType)
	case int64:
		b.ID = int(valueType)
	case uint:
		b.ID = int(valueType)
	case uint8:
		b.ID = int(valueType)
	case uint16:
		b.ID = int(valueType)
	case uint32:
		b.ID = int(valueType)
	case uint64:
		b.ID = int(valueType)
	case float32:
		b.ID = int(valueType)
	case float64:
		b.ID = int(valueType)
	default:
		return errors.Newf(mapping.ClassFieldValue, "provided invalid value: '%T' for the primary field for model: 'BasicModel'", value)
	}
	return nil
}

// Compile time check if BasicModel implements mapping.Fielder interface.
var _ mapping.Fielder = &BasicModel{}

// GetFieldsAddress gets the address of provided 'field'.
func (b *BasicModel) GetFieldsAddress(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return &b.ID, nil
	case 1: // String
		return &b.String, nil
	case 2: // Timed
		return &b.Timed, nil
	case 3: // PtrTime
		return &b.PtrTime, nil
	case 4: // Int
		return &b.Int, nil
	case 5: // Int16
		return &b.Int16, nil
	case 6: // Varchar20
		return &b.Varchar20, nil
	case 7: // Float32
		return &b.Float32, nil
	}
	return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for given model: BasicModel'", field.Name())
}

// GetFieldZeroValue implements mapping.Fielder interface.s
func (b *BasicModel) GetFieldZeroValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return 0, nil
	case 1: // String
		return "", nil
	case 2: // Timed
		return time.Time{}, nil
	case 3: // PtrTime
		return (*time.Time)(nil), nil
	case 4: // Int
		return 0, nil
	case 5: // Int16
		return 0, nil
	case 6: // Varchar20
		return "", nil
	case 7: // Float32
		return 0, nil
	default:
		return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
	}
}

// IsFieldZero implements mapping.Fielder interface.
func (b *BasicModel) IsFieldZero(field *mapping.StructField) (bool, error) {
	switch field.Index[0] {
	case 0: // ID
		return b.ID == 0, nil
	case 1: // String
		return b.String == "", nil
	case 2: // Timed
		return b.Timed == time.Time{}, nil
	case 3: // PtrTime
		return b.PtrTime == nil, nil
	case 4: // Int
		return b.Int == 0, nil
	case 5: // Int16
		return b.Int16 == 0, nil
	case 6: // Varchar20
		return b.Varchar20 == "", nil
	case 7: // Float32
		return b.Float32 == 0, nil
	}
	return false, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
}

// SetFieldZeroValue implements mapping.Fielder interface.s
func (b *BasicModel) SetFieldZeroValue(field *mapping.StructField) error {
	switch field.Index[0] {
	case 0: // ID
		b.ID = 0
	case 1: // String
		b.String = ""
	case 2: // Timed
		b.Timed = time.Time{}
	case 3: // PtrTime
		b.PtrTime = nil
	case 4: // Int
		b.Int = 0
	case 5: // Int16
		b.Int16 = 0
	case 6: // Varchar20
		b.Varchar20 = ""
	case 7: // Float32
		b.Float32 = 0
	default:
		return errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
	}
	return nil
}

// GetHashableFieldValue implements mapping.Fielder interface.
func (b *BasicModel) GetHashableFieldValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return b.ID, nil
	case 1: // String
		return b.String, nil
	case 2: // Timed
		return b.Timed, nil
	case 3: // PtrTime
		if b.PtrTime == nil {
			return b.PtrTime, nil
		}
		return *b.PtrTime, nil
	case 4: // Int
		return b.Int, nil
	case 5: // Int16
		return b.Int16, nil
	case 6: // Varchar20
		return b.Varchar20, nil
	case 7: // Float32
		return b.Float32, nil
	}
	return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for given model: 'BasicModel'", field.Name())
}

// GetFieldValue implements mapping.Fielder interface.
func (b *BasicModel) GetFieldValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 0: // ID
		return b.ID, nil
	case 1: // String
		return b.String, nil
	case 2: // Timed
		return b.Timed, nil
	case 3: // PtrTime
		return b.PtrTime, nil
	case 4: // Int
		return b.Int, nil
	case 5: // Int16
		return b.Int16, nil
	case 6: // Varchar20
		return b.Varchar20, nil
	case 7: // Float32
		return b.Float32, nil
	}
	return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for given model: BasicModel'", field.Name())
}

// SetFieldValue implements mapping.Fielder interface.
func (b *BasicModel) SetFieldValue(field *mapping.StructField, value interface{}) (err error) {
	switch field.Index[0] {
	case 0: // ID
		if v, ok := value.(int); ok {
			b.ID = v
			return nil
		}
		switch v := value.(type) {
		case int8:
			b.ID = int(v)
		case int16:
			b.ID = int(v)
		case int32:
			b.ID = int(v)
		case int64:
			b.ID = int(v)
		case uint:
			b.ID = int(v)
		case uint8:
			b.ID = int(v)
		case uint16:
			b.ID = int(v)
		case uint32:
			b.ID = int(v)
		case uint64:
			b.ID = int(v)
		case float32:
			b.ID = int(v)
		case float64:
			b.ID = int(v)
		default:
			return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
		}
		return nil
	case 1: // String
		if v, ok := value.(string); ok {
			b.String = v
			return nil
		}
		// Check alternate types for the String.
		if v, ok := value.([]byte); ok {
			b.String = string(v)
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 2: // Timed
		if v, ok := value.(time.Time); ok {
			b.Timed = v
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 3: // PtrTime
		if value == nil {
			b.PtrTime = nil
			return nil
		}
		if v, ok := value.(*time.Time); ok {
			b.PtrTime = v
			return nil
		}
		// Check if it is non-pointer value.
		if v, ok := value.(time.Time); ok {
			b.PtrTime = &v
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 4: // Int
		if v, ok := value.(int); ok {
			b.Int = v
			return nil
		}
		switch v := value.(type) {
		case int8:
			b.Int = int(v)
		case int16:
			b.Int = int(v)
		case int32:
			b.Int = int(v)
		case int64:
			b.Int = int(v)
		case uint:
			b.Int = int(v)
		case uint8:
			b.Int = int(v)
		case uint16:
			b.Int = int(v)
		case uint32:
			b.Int = int(v)
		case uint64:
			b.Int = int(v)
		case float32:
			b.Int = int(v)
		case float64:
			b.Int = int(v)
		default:
			return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
		}
		return nil
	case 5: // Int16
		if v, ok := value.(int16); ok {
			b.Int16 = v
			return nil
		}
		switch v := value.(type) {
		case int:
			b.Int16 = int16(v)
		case int8:
			b.Int16 = int16(v)
		case int32:
			b.Int16 = int16(v)
		case int64:
			b.Int16 = int16(v)
		case uint:
			b.Int16 = int16(v)
		case uint8:
			b.Int16 = int16(v)
		case uint16:
			b.Int16 = int16(v)
		case uint32:
			b.Int16 = int16(v)
		case uint64:
			b.Int16 = int16(v)
		case float32:
			b.Int16 = int16(v)
		case float64:
			b.Int16 = int16(v)
		default:
			return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
		}
		return nil
	case 6: // Varchar20
		if v, ok := value.(string); ok {
			b.Varchar20 = v
			return nil
		}
		// Check alternate types for the Varchar20.
		if v, ok := value.([]byte); ok {
			b.Varchar20 = string(v)
			return nil
		}
		return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	case 7: // Float32
		if v, ok := value.(float32); ok {
			b.Float32 = v
			return nil
		}
		switch v := value.(type) {
		case int:
			b.Float32 = float32(v)
		case int8:
			b.Float32 = float32(v)
		case int16:
			b.Float32 = float32(v)
		case int32:
			b.Float32 = float32(v)
		case int64:
			b.Float32 = float32(v)
		case uint:
			b.Float32 = float32(v)
		case uint8:
			b.Float32 = float32(v)
		case uint16:
			b.Float32 = float32(v)
		case uint32:
			b.Float32 = float32(v)
		case uint64:
			b.Float32 = float32(v)
		case float64:
			b.Float32 = float32(v)
		default:
			return errors.Newf(mapping.ClassFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
		}
		return nil
	default:
		return errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for the model: 'BasicModel'", field.Name())
	}
}
