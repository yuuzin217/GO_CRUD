package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"todo_app/app/models"
	"todo_app/config"
)

/*
generateHTML は HTMLを生成します。
*/
func generateHTML(w http.ResponseWriter, data interface{}, fileNames ...string) {
	var files []string
	for _, file := range fileNames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

/*
checkSession は セッションの確認をおこないます。
*/
func checkSession(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		session = models.Session{UUID: cookie.Value}
		if ok, _ := session.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return session, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		q := validPath.FindStringSubmatch(r.URL.Path)
		for _, i := range q {
			log.Println(i)
		}
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}

}

/*
StartMainServer は サーバーを起動します。
*/
func StartMainServer() error {
	// FileServerでルートディレクトリを指定する。
	files := http.FileServer(http.Dir(config.Config.Static))
	/*
		指定されたパターンのハンドラーを DefaultServeMux に登録する。
		StripPrefix では URL のパスから prefix を削除している。
	*/
	http.Handle("/static/", http.StripPrefix("/static/", files))

	/*
		URLに対応するハンドラー関数を登録する
		func(ResponseWriter, *Request)のハンドラー関数を実装する必要がある
	*/
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	return http.ListenAndServe("localhost:"+config.Config.Port, nil)
}
