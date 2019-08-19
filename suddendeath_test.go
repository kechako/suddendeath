package suddendeath

import "testing"

var tests = map[string]string{
	// Neutral (10)
	"🍣🍺🍖🍷🍸": `＿人人人人人人人＿
＞　🍣🍺🍖🍷🍸　＜
￣Y^Y^Y^Y^Y^Y^Y￣`,
	// Ambiguous (10)
	"※○△□☆": `＿人人人人人人人＿
＞　※○△□☆　＜
￣Y^Y^Y^Y^Y^Y^Y￣`,
	// Wide (10)
	"あいうえお": `＿人人人人人人人＿
＞　あいうえお　＜
￣Y^Y^Y^Y^Y^Y^Y￣`,
	// Narrow (5)
	"aaaaa": `＿人人人人人＿
＞　aaaaa 　＜
￣Y^Y^Y^Y^Y￣`,
	// Fullwidth (10)
	"ａｂｃｄｅ": `＿人人人人人人人＿
＞　ａｂｃｄｅ　＜
￣Y^Y^Y^Y^Y^Y^Y￣`,
	// Halfwidth (5)
	"ｱｲｳｴｵ": `＿人人人人人＿
＞　ｱｲｳｴｵ 　＜
￣Y^Y^Y^Y^Y￣`,
	// Mixed (10)
	"🍣※あaａｱ": `＿人人人人人人人＿
＞　🍣※あaａｱ　＜
￣Y^Y^Y^Y^Y^Y^Y￣`,
}

func TestGenerate(t *testing.T) {
	for text, expect := range tests {
		text := text
		expect := expect
		t.Run(text, func(t *testing.T) {
			got := Generate(text)
			if got != expect {
				t.Errorf("want:\n%v\ngot:\n%v\n", expect, got)
			}
		})
	}
}
