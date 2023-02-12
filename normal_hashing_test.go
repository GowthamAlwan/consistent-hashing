package consistenthashing

import "testing"

func TestSetVal(t *testing.T) {
	setVal("name", "gowtham")
}

func TestGetVal(t *testing.T) {
	t.Run("404", func(t *testing.T) {
		getVal("name")
	})

	t.Run("200", func(t *testing.T) {
		setVal("name", "gowtham")
		getVal("name")
	})
}
