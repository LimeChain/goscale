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

func (d Dictionary[K, V]) Encode(buffer *bytes.Buffer) error {
	err := ToCompact(len(d)).Encode(buffer)
	if err != nil {
		return err
	}

	keys := make([]K, 0)
	for k := range d {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, k := range keys {
		v := d[k]
		err := k.Encode(buffer)
		if err != nil {
			return err
		}

		err = v.Encode(buffer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d Dictionary[K, V]) Bytes() []byte {
	return EncodedBytes(d)
}

func DecodeDictionary[K Comparable, V Encodable](buffer *bytes.Buffer) (Dictionary[K, V], error) {
	result := Dictionary[K, V]{}

	v, err := DecodeCompact[BigNumbers](buffer)
	if err != nil {
		return nil, err
	}
	size := int(v.ToBigInt().Int64())

	for i := 0; i < size; i++ {
		key, err := decodeByType(*new(K), buffer)
		if err != nil {
			return Dictionary[K, V]{}, err
		}
		value, err := decodeByType(*new(V), buffer)
		if err != nil {
			return Dictionary[K, V]{}, err
		}
		result[key.(K)] = value.(V)
	}

	return result, nil
}
