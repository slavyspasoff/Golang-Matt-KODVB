package hello

import "testing"

func TestSay(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"Matt", "Anne"},
			result: "Hello, Matt, Anne!",
		},
	}
	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("Want: %s (%v), got %s", st.result, st.items, s)
		}
	}

}
