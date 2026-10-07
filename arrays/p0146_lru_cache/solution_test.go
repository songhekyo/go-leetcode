package p0146lrucache

import "testing"

type op struct {
	name  string // "put" atau "get"
	key   int
	value int // untuk put
	want  int // untuk get
}

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []op
	}{
		{
			name:     "example_1",
			capacity: 2,
			ops: []op{
				{name: "put", key: 1, value: 1},
				{name: "put", key: 2, value: 2},
				{name: "get", key: 1, want: 1},
				{name: "put", key: 3, value: 3}, // buang key 2
				{name: "get", key: 2, want: -1},
				{name: "put", key: 4, value: 4}, // buang key 1
				{name: "get", key: 1, want: -1},
				{name: "get", key: 3, want: 3},
				{name: "get", key: 4, want: 4},
			},
		},
		{
			name:     "update_existing_key",
			capacity: 2,
			ops: []op{
				{name: "put", key: 1, value: 1},
				{name: "put", key: 2, value: 2},
				{name: "put", key: 1, value: 10}, // update, 1 jadi paling baru
				{name: "put", key: 3, value: 3},  // buang key 2
				{name: "get", key: 1, want: 10},
				{name: "get", key: 2, want: -1},
				{name: "get", key: 3, want: 3},
			},
		},
		{
			name:     "capacity_one",
			capacity: 1,
			ops: []op{
				{name: "put", key: 1, value: 1},
				{name: "put", key: 2, value: 2}, // buang key 1
				{name: "get", key: 1, want: -1},
				{name: "get", key: 2, want: 2},
			},
		},
		{
			name:     "get_missing_on_empty",
			capacity: 2,
			ops: []op{
				{name: "get", key: 5, want: -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor(tt.capacity)
			for i, o := range tt.ops {
				switch o.name {
				case "put":
					c.Put(o.key, o.value)
				case "get":
					if got := c.Get(o.key); got != o.want {
						t.Errorf("op #%d: Get(%d) = %d, want %d", i, o.key, got, o.want)
					}
				}
			}
		})
	}
}
