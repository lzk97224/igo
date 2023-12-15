package imap

import (
	"reflect"
	"sort"
)

type order interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

func asc[K order](k1, k2 K) bool {
	return k1 < k2
}

type sortByValueFunc[V any] func(v1, v2 V) bool

type sortMap[K order, V any] struct {
	sortByValue sortByValueFunc[V]
	keys        []K
	data        map[K]V
}

func NewSortMap[K order, V any]() *sortMap[K, V] {
	nsm := &sortMap[K, V]{
		keys: make([]K, 0),
		data: make(map[K]V),
	}
	return nsm
}

func NewSortMapByValue[K order, V any](sort sortByValueFunc[V]) *sortMap[K, V] {
	nsm := NewSortMap[K, V]()
	nsm.sortByValue = sort
	return nsm
}

func (m *sortMap[K, V]) Len() int {
	return len(m.keys)
}

func (m *sortMap[K, V]) Less(i, j int) bool {
	if m.sortByValue != nil {
		return m.sortByValue(m.data[m.keys[i]], m.data[m.keys[j]])
	}
	return asc(m.keys[i], m.keys[j])
}

func (m *sortMap[K, V]) Swap(i, j int) {
	m.keys[i], m.keys[j] = m.keys[j], m.keys[i]
}

func (m *sortMap[K, V]) Sort() {
	sort.Sort(m)
}

func (m *sortMap[K, V]) ReverseSort() {
	sort.Sort(sort.Reverse(m))
}
func (m *sortMap[K, V]) Range(f func(key K, value V) bool) {
	for _, k := range m.keys {
		if !f(k, m.data[k]) {
			break
		}
	}
}

func (m *sortMap[K, V]) Set(key K, value V) {
	if _, ok := m.data[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.data[key] = value
}

func (m *sortMap[K, V]) GetNoNil(key K) V {
	v, ok := m.data[key]
	if !ok {
		var r V
		t := reflect.TypeOf(r)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			nv := reflect.New(t)
			nvi := nv.Interface()
			return nvi.(V)
		}
		return r
	}
	return v
}

func (m *sortMap[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

func (m *sortMap[K, V]) Delete(key K) {
	delete(m.data, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

func (m *sortMap[K, V]) Keys() []K {
	return m.keys
}

func (m *sortMap[K, V]) Values() []V {
	values := make([]V, 0)
	for _, k := range m.keys {
		values = append(values, m.data[k])
	}
	return values
}
