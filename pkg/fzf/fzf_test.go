package fzf

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"testing"
)

type testItem struct {
	title string
}

func (ti testItem) Id() string {
	return ti.title
}

func (ti testItem) String() string {
	return fmt.Sprintf("%s", ti.title)
}

type testItems []testItem

func (tis testItems) ListString() []string {
	var strs []string
	for _, i := range tis {
		strs = append(strs, i.String())
	}
	return strs
}

func (tis testItems) EachItem(handler func(Item)) {
	for _, ti := range tis {
		handler(ti)
	}
}

func TestFzf_setFilteredItems(t *testing.T) {
	items := testItems{
		{
			title: "unittest1",
		},
		{
			title: "unittest2",
		},
		{
			title: "unittest3",
		},
	}
	expect := []testItem{
		{
			title: "unittest2",
		},
	}

	f := New(items)
	f.searchingText = "2"
	f.setFilteredItems()

	if len(f.filteredItems) != len(expect) {
		t.Fatalf("filtered items length is not equal with expected")
	}
	for i, item := range f.filteredItems {
		if item.Id() != expect[i].Id() {
			t.Fatalf("filtered items elements is not queal with expected")
		}
	}
}

func TestFzf_isQuitInput(t *testing.T) {
	tests := []struct {
		key    keyboard.Key
		expect bool
	}{
		{
			key:    keyboard.KeyEsc,
			expect: true,
		},
		{
			key:    keyboard.KeyArrowRight,
			expect: false,
		},
	}

	f := New(nil)
	for _, test := range tests {
		if test.expect != f.isQuitInput(test.key) {
			t.Fatalf("is quit input result is not expected %v", test.expect)
		}
	}
}

func TestFzf_isIgnoreInput(t *testing.T) {
	tests := []struct {
		key    keyboard.Key
		expect bool
	}{
		{
			key:    keyboard.KeyArrowLeft,
			expect: true,
		},
		{
			key:    keyboard.KeyArrowUp,
			expect: false,
		},
	}

	f := New(nil)
	for _, test := range tests {
		if test.expect != f.isIgnoreInput(test.key) {
			t.Fatalf("is quit input result is not expected %v", test.expect)
		}
	}
}

func TestFzf_deleteSearchingText(t *testing.T) {
	tests := []struct {
		searchingText string
		expect        string
	}{
		{
			searchingText: "",
			expect:        "",
		},
		{
			searchingText: "네이버",
			expect:        "네이",
		},
		{
			searchingText: "NAVER",
			expect:        "NAVE",
		},
	}

	for _, test := range tests {
		f := Fzf{searchingText: test.searchingText}
		f.deleteSearchingText()

		if test.expect != f.searchingText {
			t.Fatalf("deleting searching text result %s is not expected %s", f.searchingText, test.expect)
		}
	}
}

func TestFzf_decreaseIndexInFilteredItems(t *testing.T) {
	tests := []struct {
		searchingIndex int
		expect         int
	}{
		{
			searchingIndex: 0,
			expect:         0,
		},
		{
			searchingIndex: 1,
			expect:         0,
		},
	}

	for _, test := range tests {
		f := Fzf{searchingIndex: test.searchingIndex}
		f.decreaseIndexInFilteredItems()

		if test.expect != f.searchingIndex {
			t.Fatalf("decresing index result %d is not expected %d", f.searchingIndex, test.expect)
		}
	}
}

func TestFzf_increaseIndexInFilteredItems(t *testing.T) {
	items := testItems{
		{
			title: "unittest1",
		},
		{
			title: "unittest2",
		},
	}
	tests := []struct {
		searchingIndex int
		expect         int
	}{
		{
			searchingIndex: 0,
			expect:         1,
		},
		{
			searchingIndex: 1,
			expect:         1,
		},
	}

	for _, test := range tests {
		f := New(items)
		f.searchingIndex = test.searchingIndex
		f.increaseIndexInFilteredItems()
		if test.expect != f.searchingIndex {
			t.Fatalf("decresing index result %d is not expected %d", f.searchingIndex, test.expect)
		}
	}
}
