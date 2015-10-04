package population

import "testing"

func TestListCountries(t *testing.T) {
	ts, c := testServerAndClientWithJSON(200, []byte(`{
		"countries":["Afghanistan", "Albania", "Algeria"]
	}`))
	defer ts.Close()

	list, err := c.ListCountries()
	if err != nil {
		t.Fatalf(`unexpected error %v`, err)
	}

	if got, want := len(list.Countries), 3; got != want {
		t.Fatalf(`len(list.Countries) = %d, want %d`, got, want)
	}
}
