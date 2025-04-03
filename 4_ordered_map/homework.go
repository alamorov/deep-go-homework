package main

import "golang.org/x/exp/constraints"

type node[K constraints.Ordered, V any] struct {
	key   K
	value V
	left  *node[K, V]
	right *node[K, V]
}

type OrderedMap[K constraints.Ordered, V any] struct {
	root *node[K, V]
	size int
}

func NewOrderedMap[K constraints.Ordered, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{
		root: nil,
		size: 0,
	}
}

func (m *OrderedMap[K, V]) Insert(key K, value V) {
	m.root = m.insert(m.root, key, value)
}

func (m *OrderedMap[K, V]) insert(root *node[K, V], key K, value V) *node[K, V] {
	if root == nil {
		m.size++
		return &node[K, V]{key: key, value: value}
	}

	if key < root.key {
		root.left = m.insert(root.left, key, value)
	} else if key > root.key {
		root.right = m.insert(root.right, key, value)
	} else {
		root.value = value
	}
	return root
}

func (m *OrderedMap[K, V]) Erase(key K) {
	var ok bool
	m.root, ok = m.erase(m.root, key)
	if ok {
		m.size--
	}
}

func (m *OrderedMap[K, V]) erase(root *node[K, V], key K) (*node[K, V], bool) {
	if root == nil {
		return nil, false
	}

	var ok bool
	if key < root.key {
		root.left, ok = m.erase(root.left, key)
	} else if key > root.key {
		root.right, ok = m.erase(root.right, key)
	} else {
		if root.left == nil {
			return root.right, true
		} else if root.right == nil {
			return root.left, true
		}

		minNode := root.right
		for minNode.left != nil {
			minNode = minNode.left
		}
		root.key = minNode.key
		root.value = minNode.value
		root.right, _ = m.erase(root.right, minNode.key)
		return root, true
	}
	return root, ok
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	return m.contains(m.root, key)
}

func (m *OrderedMap[K, V]) contains(root *node[K, V], key K) bool {
	if root == nil {
		return false
	}

	if key < root.key {
		return m.contains(root.left, key)
	} else if key > root.key {
		return m.contains(root.right, key)
	} else {
		return true
	}
}

func (m *OrderedMap[K, V]) Size() int {
	return m.size
}

func (m *OrderedMap[K, V]) ForEach(action func(K, V)) {
	m.forEach(m.root, action)
}

func (m *OrderedMap[K, V]) forEach(root *node[K, V], action func(K, V)) {
	if root == nil {
		return
	}

	m.forEach(root.left, action)
	action(root.key, root.value)
	m.forEach(root.right, action)
}
