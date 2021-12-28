package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//Table for testing
var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "no error rendering non=existent go template, when one is expected"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering non=existent jet template, when one is expected"},
	{"invalid_render_engine", "foo", "home", true, "no error rendering non-existent template"},
}

func TestRender_Page(t *testing.T) {
	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()
		testRender.Renderer = e.renderer
		testRender.RootPath = "./testdata"
		err = testRender.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}
	// changed to table test
	//testRender.Renderer = "jet"
	//err = testRender.Page(w, r, "home", nil, nil)
	//if err != nil {
	//	t.Error("Error jet rendering page", err)
	//}
	//
	//testRender.Renderer = ""
	//err = testRender.Page(w, r, "home", nil, nil)
	//if err == nil {
	//	t.Error("Error jet rendering page", err)
	//}

}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}
	testRender.Renderer = "go"
	testRender.RootPath = "./testdata"
	err = testRender.GoPage(w, r, "home", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}
	testRender.Renderer = "jet"
	testRender.RootPath = "./testdata"
	err = testRender.JetPage(w, r, "home", nil, nil)
	if err != nil {
		t.Error(err)
	}
}
