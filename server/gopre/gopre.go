package gopre

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
)

func init() {
	router()
}

func router() {
	http.HandleFunc("/", topHandler)
	http.HandleFunc("/index", topHandler)
}

func Display(w http.ResponseWriter, r *http.Request, name []string, handler func(http.ResponseWriter, *http.Request, map[string]interface{})) error {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	path := "gopre/templates/"

	templates := make([]string, len(name)+1)
	templates[0] = path + "layout.tmpl"
	for i, v := range name {
		templates[i+1] = path + v
	}

	var listTmpl = template.Must(template.ParseFiles(templates...))

	tc := make(map[string]interface{})
	c := appengine.NewContext(r)
	u := user.Current(c)
	tc["User"] = u
	if u != nil {
		url, _ := user.LogoutURL(c, "/")
		tc["LogoutURL"] = url
	}

	if handler != nil {
		handler(w, r, tc)
	}

	if err := listTmpl.Execute(w, tc); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return err
	}
	return nil
}

func topHandler(w http.ResponseWriter, r *http.Request) {
	Display(w, r, St2Sl("index.tmpl"), nil)
}

func St2Sl(val ...string) []string {
	return []string(val)
}
