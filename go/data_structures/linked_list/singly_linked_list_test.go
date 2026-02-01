package linked_list

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	list := New()
	if list.Head != nil {
		t.Error("Expected Head to be nil")
	}
	if list.Size != 0 {
		t.Errorf("Expected Size to be 0, got %d", list.Size)
	}
}

func TestAppend(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	if list.Size != 3 {
		t.Errorf("Expected Size to be 3, got %d", list.Size)
	}

	expected := []int{1, 2, 3}
	actual := list.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPrepend(t *testing.T) {
	list := New()
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	if list.Size != 3 {
		t.Errorf("Expected Size to be 3, got %d", list.Size)
	}

	expected := []int{3, 2, 1}
	actual := list.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsertAt(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		index    int
		value    int
		expected []int
		success  bool
	}{
		{
			name:     "Insert at beginning",
			initial:  []int{1, 2, 3},
			index:    0,
			value:    0,
			expected: []int{0, 1, 2, 3},
			success:  true,
		},
		{
			name:     "Insert at middle",
			initial:  []int{1, 2, 4},
			index:    2,
			value:    3,
			expected: []int{1, 2, 3, 4},
			success:  true,
		},
		{
			name:     "Insert at end",
			initial:  []int{1, 2, 3},
			index:    3,
			value:    4,
			expected: []int{1, 2, 3, 4},
			success:  true,
		},
		{
			name:     "Insert at invalid index",
			initial:  []int{1, 2, 3},
			index:    5,
			value:    4,
			expected: []int{1, 2, 3},
			success:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New()
			for _, v := range tt.initial {
				list.Append(v)
			}

			result := list.InsertAt(tt.index, tt.value)
			if result != tt.success {
				t.Errorf("Expected InsertAt to return %v, got %v", tt.success, result)
			}

			actual := list.ToSlice()
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		delete   int
		expected []int
		found    bool
	}{
		{
			name:     "Delete from middle",
			initial:  []int{1, 2, 3, 4},
			delete:   3,
			expected: []int{1, 2, 4},
			found:    true,
		},
		{
			name:     "Delete from head",
			initial:  []int{1, 2, 3},
			delete:   1,
			expected: []int{2, 3},
			found:    true,
		},
		{
			name:     "Delete from tail",
			initial:  []int{1, 2, 3},
			delete:   3,
			expected: []int{1, 2},
			found:    true,
		},
		{
			name:     "Delete non-existent",
			initial:  []int{1, 2, 3},
			delete:   5,
			expected: []int{1, 2, 3},
			found:    false,
		},
		{
			name:     "Delete from empty list",
			initial:  []int{},
			delete:   1,
			expected: []int{},
			found:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New()
			for _, v := range tt.initial {
				list.Append(v)
			}

			result := list.Delete(tt.delete)
			if result != tt.found {
				t.Errorf("Expected Delete to return %v, got %v", tt.found, result)
			}

			actual := list.ToSlice()
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestDeleteAt(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		index    int
		expected []int
		success  bool
	}{
		{
			name:     "Delete at beginning",
			initial:  []int{1, 2, 3},
			index:    0,
			expected: []int{2, 3},
			success:  true,
		},
		{
			name:     "Delete at middle",
			initial:  []int{1, 2, 3},
			index:    1,
			expected: []int{1, 3},
			success:  true,
		},
		{
			name:     "Delete at end",
			initial:  []int{1, 2, 3},
			index:    2,
			expected: []int{1, 2},
			success:  true,
		},
		{
			name:     "Delete at invalid index",
			initial:  []int{1, 2, 3},
			index:    5,
			expected: []int{1, 2, 3},
			success:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New()
			for _, v := range tt.initial {
				list.Append(v)
			}

			result := list.DeleteAt(tt.index)
			if result != tt.success {
				t.Errorf("Expected DeleteAt to return %v, got %v", tt.success, result)
			}

			actual := list.ToSlice()
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestFind(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	tests := []struct {
		name  string
		value int
		found bool
	}{
		{"Find existing value", 2, true},
		{"Find non-existing value", 5, false},
		{"Find head value", 1, true},
		{"Find tail value", 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := list.Find(tt.value)
			if tt.found && node == nil {
				t.Error("Expected to find node, got nil")
			}
			if !tt.found && node != nil {
				t.Errorf("Expected not to find node, got %v", node)
			}
			if tt.found && node.Value != tt.value {
				t.Errorf("Expected node value to be %d, got %d", tt.value, node.Value)
			}
		})
	}
}

func TestGet(t *testing.T) {
	list := New()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	tests := []struct {
		name     string
		index    int
		expected int
		found    bool
	}{
		{"Get at index 0", 0, 10, true},
		{"Get at index 1", 1, 20, true},
		{"Get at index 2", 2, 30, true},
		{"Get at invalid index", 5, 0, false},
		{"Get at negative index", -1, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := list.Get(tt.index)
			if found != tt.found {
				t.Errorf("Expected found to be %v, got %v", tt.found, found)
			}
			if tt.found && value != tt.expected {
				t.Errorf("Expected value to be %d, got %d", tt.expected, value)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected []int
	}{
		{
			name:     "Reverse multiple elements",
			initial:  []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Reverse single element",
			initial:  []int{1},
			expected: []int{1},
		},
		{
			name:     "Reverse two elements",
			initial:  []int{1, 2},
			expected: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New()
			for _, v := range tt.initial {
				list.Append(v)
			}

			list.Reverse()
			actual := list.ToSlice()
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	list := New()
	if !list.IsEmpty() {
		t.Error("Expected new list to be empty")
	}

	list.Append(1)
	if list.IsEmpty() {
		t.Error("Expected list with elements to not be empty")
	}

	list.Delete(1)
	if !list.IsEmpty() {
		t.Error("Expected list after deleting all elements to be empty")
	}
}

func TestClear(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Clear()
	if !list.IsEmpty() {
		t.Error("Expected list to be empty after Clear()")
	}
	if list.Size != 0 {
		t.Errorf("Expected Size to be 0 after Clear(), got %d", list.Size)
	}
}
