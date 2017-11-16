package suddendeath

import "testing"

func TestTextWidth(t *testing.T) {
	pattern := map[string]int{
		// Neutral
		"🍣🍺🍖🍷🍸": 10,
		// Ambiguous
		"※○△□☆": 10,
		// Wide
		"あいうえお": 10,
		// Narrow
		"aaaaa": 5,
		// Fullwidth
		"ａｂｃｄｅ": 10,
		// Halfwidth
		"ｱｲｳｴｵ": 5,
		// Mixed
		"🍣※あaａｱ": 10,
	}

	for str, expect := range pattern {
		w := textWidth(str)

		if w != expect {
			t.Errorf("want %v\ngot %v", expect, w)
		}
	}
}

func TestGenerate(t *testing.T) {
	expect := `＿人人人人人人＿
＞　突然の死　＜
￣Y^Y^Y^Y^Y^Y￣`
	got := Generate("突然の死")
	if got != expect {
		t.Errorf("want:\n%v\ngot:\n%v\n", expect, got)
	}

	expect = `＿人人人人＿
＞　abc 　＜
￣Y^Y^Y^Y￣`
	got = Generate("abc")
	if got != expect {
		t.Errorf("want:\n%v\ngot:\n%v\n", expect, got)
	}

	expect = `＿人人人人人人人人＿
＞　突 然 の 死 　＜
￣Y^Y^Y^Y^Y^Y^Y^Y￣`
	got = Generate("突 然 の 死")
	if got != expect {
		t.Errorf("want:\n%v\ngot:\n%v\n", expect, got)
	}
}
