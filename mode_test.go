package viamstreamdeck

import (
	"testing"

	"go.viam.com/test"
)

func TestExtractMode(t *testing.T) {
	t.Run("present string", func(t *testing.T) {
		mode, ok := extractMode(map[string]interface{}{"mode": "menu"})
		test.That(t, ok, test.ShouldBeTrue)
		test.That(t, mode, test.ShouldEqual, "menu")
	})

	t.Run("missing key", func(t *testing.T) {
		mode, ok := extractMode(map[string]interface{}{"other": "menu"})
		test.That(t, ok, test.ShouldBeFalse)
		test.That(t, mode, test.ShouldEqual, "")
	})

	t.Run("non-string value", func(t *testing.T) {
		mode, ok := extractMode(map[string]interface{}{"mode": 1.0})
		test.That(t, ok, test.ShouldBeFalse)
		test.That(t, mode, test.ShouldEqual, "")
	})

	t.Run("empty string value", func(t *testing.T) {
		mode, ok := extractMode(map[string]interface{}{"mode": ""})
		test.That(t, ok, test.ShouldBeFalse)
		test.That(t, mode, test.ShouldEqual, "")
	})

	t.Run("nil response", func(t *testing.T) {
		mode, ok := extractMode(nil)
		test.That(t, ok, test.ShouldBeFalse)
		test.That(t, mode, test.ShouldEqual, "")
	})
}
