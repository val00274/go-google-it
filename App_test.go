package googleit

import "testing"

func TestApp_QueryEscape(t *testing.T) {
	cases := []struct {
		Input  []string
		Expect string
	}{
		{[]string{"hello"}, "hello"},
		{[]string{"hello world"}, "hello+world"},
		{[]string{"こんにちは"}, "%E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%81%AF"},
	}

	app := NewApp()
	for _, c := range cases {
		if ret := app.QueryEscape(c.Input...); ret != c.Expect {
			t.Errorf(`expected "%s", but it was actually "%s"`, c.Expect, ret)
		}
	}
}
