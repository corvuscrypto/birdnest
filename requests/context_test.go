package requests

import "testing"

func TestContext(T *testing.T) {
	ctx := make(Context)
	if ctx.Get("test") != nil {
		T.Errorf("Received incorrect value")
	}
	if ctx.Get("test", 100) != 100 {
		T.Errorf("Received incorrect value")
	}

	ctx.Set("test", 12)
	if ctx.Get("test") != 12 {
		T.Errorf("Received incorrect value")
	}
	ctx.Set("test", 13)
	if ctx.Get("test") != 13 {
		T.Errorf("Received incorrect value")
	}
	ctx.Set("test", 100, true)
	if ctx.Get("test") != 13 {
		T.Errorf("Received incorrect value")
	}
}
