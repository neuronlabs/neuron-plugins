package memory

import (
	"context"
	"os"
	"strings"

	"github.com/patrickmn/go-cache"

	"github.com/neuronlabs/neuron/errors"
	"github.com/neuronlabs/neuron/log"
	"github.com/neuronlabs/neuron/store"
)

// Compile time check if memory implements store interface.
var _ store.Store = &Memory{}

// Memory is a in-memory store implementation for neuron store.
type Memory struct {
	cache   *cache.Cache
	Options *store.Options
}

// New creates new in-memory store.
func New(options ...store.Option) (*Memory, error) {
	m := &Memory{
		Options: store.DefaultOptions(),
	}
	for _, option := range options {
		option(m.Options)
	}
	m.cache = cache.New(m.Options.DefaultExpiration, m.Options.CleanupInterval)
	if m.Options.FileName != "" {
		if err := m.cache.LoadFile(m.Options.FileName); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return nil, err
			}
		}
	}
	return m, nil
}

// Set implements store.Store interface.
func (m *Memory) Set(_ context.Context, record *store.Record, options ...store.SetOption) error {
	o := &store.SetOptions{}
	for _, option := range options {
		option(o)
	}
	ttl := m.Options.DefaultExpiration
	if o.TTL != 0 {
		ttl = o.TTL
	}
	key := m.key(record.Key)
	if record.ExpiresAt.IsZero() {
		record.ExpiresAt = m.Options.TimeFunc().Add(ttl)
	}
	cp := make([]byte, len(record.Value))
	copy(cp, record.Value)

	m.cache.Set(key, cp, ttl)
	return nil
}

// Get implements store.Store interface.
func (m *Memory) Get(_ context.Context, key string) (*store.Record, error) {
	v, expiration, found := m.cache.GetWithExpiration(m.key(key))
	if !found {
		return nil, store.ErrRecordNotFound
	}
	value, ok := v.([]byte)
	if !ok {
		log.Errorf("Malformed store value type: %T", v)
		return nil, errors.Wrap(store.ErrInternal, "malformed record type")
	}
	return &store.Record{
		Key:       key,
		Value:     value,
		ExpiresAt: expiration,
	}, nil
}

// Delete implements store.Store interface.
func (m *Memory) Delete(_ context.Context, key string) error {
	_, found := m.cache.Get(m.key(key))
	if !found {
		return store.ErrRecordNotFound
	}
	m.cache.Delete(m.key(key))
	return nil
}

// Find implements store.Store.
func (m *Memory) Find(_ context.Context, options ...store.FindOption) ([]*store.Record, error) {
	findOptions := &store.FindPattern{}
	for _, option := range options {
		option(findOptions)
	}
	items := m.cache.Items()
	var records []*store.Record
	prefix := m.Options.Prefix + findOptions.Prefix
	suffix := m.Options.Suffix + findOptions.Suffix
	for k, v := range items {
		if prefix != "" && !strings.HasPrefix(k, prefix) {
			continue
		}
		if suffix != "" && !strings.HasSuffix(k, suffix) {
			continue
		}
		if findOptions.Offset != 0 {
			findOptions.Offset--
			continue
		}
		rec, ok := v.Object.(*store.Record)
		if !ok {
			return nil, errors.Wrap(store.ErrInternal, "a record is not store.Record")
		}
		records = append(records, rec.Copy())
		findOptions.Limit--
		if findOptions.Limit == 0 {
			break
		}
	}
	return records, nil
}

// Close implements io.Closer interface.
func (m *Memory) Close(context.Context) error {
	if m.Options.FileName != "" {
		if err := m.cache.SaveFile(m.Options.FileName); err != nil {
			return err
		}
		log.Debug2f("Memory store - file saved with success")
	}
	m.cache.Flush()
	log.Debug2f("Memory store closed with success")
	return nil
}

func (m *Memory) key(key string) string {
	return m.Options.Prefix + key + m.Options.Suffix
}
