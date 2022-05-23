package newsfeed

import (
	"testing"
)

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{"An item", "with body!"})
	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{"An item", "with body!"})
	feed.Add(Item{"Another", "With something else"})
	x := feed.GetAll()
	if len(x) != 2 {
		t.Errorf("Items were not added correctly. Expected = %v, Actual = %v", 2, len(x))
	}
}
