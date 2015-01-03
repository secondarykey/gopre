package gopre

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"

	"gopre/me"
)

func init() {
	router()
}

func router() {

	http.HandleFunc("/", TopHandler)
	http.HandleFunc("/index", TopHandler)

	//user dashbord
	http.HandleFunc("/me/", DashboardHandler)
	http.HandleFunc("/me/slides", SlideHandler)
	http.HandleFunc("/me/templates", TemplateHandler)

	//user pages

	//publish pages
}

func Display(w http.ResponseWriter, r *http.Request, name string, handler func(http.ResponseWriter, *http.Request, map[string]interface{})) error {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	path := "gopre/templates/"
	var listTmpl = template.Must(
		template.ParseFiles(
			path+"layout.tmpl",
			path+name))

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

func TopHandler(w http.ResponseWriter, r *http.Request) {
	Display(w, r, "index.tmpl", nil)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	function := me.GetDashboardParamFunc()
	Display(w, r, "me/dashboard.tmpl", function)
}

func SlideHandler(w http.ResponseWriter, r *http.Request) {
	function := me.GetSlideParamFunc()
	Display(w, r, "me/slide.tmpl", function)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	function := me.GetTemplateParamFunc()
	Display(w, r, "me/template.tmpl", function)
}
