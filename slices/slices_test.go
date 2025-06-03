package slices

import (
	"math"
	"strconv"
	"testing"
)

func TestPaginate(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}

	tests := []struct {
		page     int
		size     int
		want     []string
		wantPage int
	}{
		{1, 3, []string{"a", "b", "c"}, 1},
		{2, 3, []string{"d", "e", "f"}, 2},
		{3, 3, []string{"g", "h", "i"}, 3},
		{4, 3, []string{}, 4},                                             // out of range
		{1, 10, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}, 1}, // all
	}

	for _, tt := range tests {
		got := Paginate(items, tt.page, tt.size)

		if got.Page != tt.wantPage {
			t.Errorf("Paginate: got page %d, want %d", got.Page, tt.wantPage)
		}

		if got.Size != tt.size {
			t.Errorf("Paginate: got size %d, want %d", got.Size, tt.size)
		}

		if got.NumberOfElements != len(tt.want) {
			t.Errorf("Paginate: got NumberOfElements %d, want %d", got.NumberOfElements, len(tt.want))
		}

		if got.TotalElements != len(items) {
			t.Errorf("Paginate: got TotalElements %d, want %d", got.TotalElements, len(items))
		}

		expectedPages := int(math.Ceil(float64(len(items)) / float64(tt.size)))
		if got.TotalPages != expectedPages {
			t.Errorf("Paginate: got TotalPages %d, want %d", got.TotalPages, expectedPages)
		}

		if len(got.Content) != len(tt.want) {
			t.Errorf("Paginate: unexpected content length, got %d, want %d", len(got.Content), len(tt.want))
		}

		for i := range got.Content {
			if got.Content[i] != tt.want[i] {
				t.Errorf("Paginate: content[%d] = %s, want %s", i, got.Content[i], tt.want[i])
			}
		}
	}
}

func TestMap(t *testing.T) {
	ints := []int{1, 2, 3}
	toStr := func(i int) string {
		return strconv.Itoa(i)
	}

	got := Map(ints, toStr)
	want := []string{"1", "2", "3"}

	if len(got) != len(want) {
		t.Fatalf("unexpected length: got %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Map[%d]: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFilter(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6}
	isEven := func(i int) bool {
		return i%2 == 0
	}

	got := Filter(ints, isEven)
	want := []int{2, 4, 6}

	if len(got) != len(want) {
		t.Fatalf("unexpected length: got %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Filter[%d]: got %d, want %d", i, got[i], want[i])
		}
	}
}
