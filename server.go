package server

import "net/http"

type Server struct {
}

func New() *Server {
	return &Server{}
}

var indexPage = `
<!DOCTYPE html>
<html>
	<body>
		<h1 style="text-align:center;">My first heading</h1>
		<p style="text-align:center;">My first paragraph </p>
	</body>
</html>
`
var userInfo = `{
	"name":"srini",
	"age:21
}`

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexPage))
}
func (s *Server) HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userInfo))
}
