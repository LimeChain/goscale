package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-dictionary

	SCALE Dictionary type translates to Go's map type.
*/

import (
	"bytes"
	"sort"
)

type Comparable interface {
	Encodable
	Ordered
}

type Dictionary[K Comparable, V Encodable] map[K]V

func (d Dictionary[K, V]) Encode(buffer *bytes.Buffer) {
	Compact(len(d)).Encode(buffer)

	keys := make([]K, 0)
	for k := range d {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, k := range keys {
		v := d[k]
		k.Encode(buffer)
		v.Encode(buffer)
	}
}

func DecodeDictionary[K Comparable, V Encodable](buffer *bytes.Buffer) Dictionary[K, V] {
	result := Dictionary[K, V]{}

	size := int(DecodeCompact(buffer))

	for i := 0; i < size; i++ {
		key := decodeByType(*new(K), buffer)
		value := decodeByType(*new(V), buffer)
		result[key.(K)] = value.(V)
	}

	return result
}
