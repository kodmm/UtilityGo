package sample

func TestConsoleOut(t *testing.T) {
	var b bytes.Buffer
	DumpUserTo(&b, &User{Name: "クライド・バロウ"})
	if b.String() != "クライド・バロウ(住所不定)" {
		t.Errorf("error (expected: 'クライド・バロウ(住所不定)', actual='%s')", b.String())
	}
}
