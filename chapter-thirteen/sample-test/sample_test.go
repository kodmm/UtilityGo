package sample

import (
	"net/http"
	"testing"
)

func TestConsoleOut(t *testing.T) {
	var b bytes.Buffer
	DumpUserTo(&b, &User{Name: "クライド・バロウ"})
	if b.String() != "クライド・バロウ(住所不定)" {
		t.Errorf("error (expected: 'クライド・バロウ(住所不定)', actual='%s')", b.String())
	}
}

type testTransport struct {
	req **http.Request
	res *http.Response
	err error
}

func (t *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	*(t.req) = req
	return t.res, t.err
}

var _ http.RoundTripper = &testTransport{}

func newTransport(req **http.Request, res *http.Response, err error) http.RoundTripper {
	return &testTransport{
		req: req,
		res: res,
		err: err,
	}
}

// テストコード
func TestHTTPRequest(t *testing.T) {
	var req *http.Request
	res := httptest.NewRecorder()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.WriteString(`{"ranking": ["Back to the Future", "Rambo"]}`)

	c := http.Client{
		Transport: newTransport(&req, res.Result(), nil)
	}
	r, err := c.Get("http://exmaple.com/movies/1985")
	assert.Nil(t, err)
	assert.Equal(t, 200, r.StatusCode)
}

type contextTimeKey string

const timeKey contextTimeKey = "timeKey"

func CurrentTime(ctx context.Context) time.Time {
	v := ctx.Value(timeKey)
	if t, ok := v.(time.Time); ok {
		return t
	}
	return time.Now()
}

func SetFixTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, timeKey, t)
}

// テスト対象
func NextMonth(ctx context.Context) time.Month {
	now := CurrenTime(ctx)
	return now.AddDate(0, 1, 0).Month()
}

func TestNextMonth(t *testing.T) {
	ctx := SetFixTime(context.Background(), time.Date(1980, time.December, 1, 0, 0, 0, 0, time.Local))
	assert.Equal(t, time.January, NextMonth(ctx))
}


