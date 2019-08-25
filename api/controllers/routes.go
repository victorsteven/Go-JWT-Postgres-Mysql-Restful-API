package controllers

import "github.com/victorsteven/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/login", s.Login).Methods("POST")

	s.Router.HandleFunc("/", s.Home).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}

// func (s *Server) authOnly(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		err := auth.TokenValid(r)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusUnauthorized, err)
// 			return
// 		}
// 		next(w, r)
// 	}
// }

// func (s *Server) authOnly(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if !currentUser(r).IsAdmin {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		h(w, r)
// 	}
// }

// type Route struct {
// 	Uri          string
// 	Method       string
// 	Handler      func(http.ResponseWriter, *http.Request)
// 	AuthRequired bool
// }

// func Load() []Route {
// 	routes := userRoutes
// 	routes = append(routes, postsRoutes...)
// 	routes = append(routes, loginRoutes...)
// 	return routes
// }

// func SetUpRoutes(r *mux.Router) *mux.Router {
// 	for _, route := range Load() {
// 		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
// 	}
// 	return r
// }

// func SetUpRoutesWithMiddlewares(r *mux.Router) *mux.Router {
// 	for _, route := range Load() {
// 		if route.AuthRequired {
// 			r.HandleFunc(route.Uri,
// 				middlewares.SetMiddlewareLogger(
// 					middlewares.SetMiddlewareJSON(
// 						middlewares.SetMiddlewareAuthentication(route.Handler))),
// 			).Methods(route.Method)
// 		} else {
// 			r.HandleFunc(route.Uri, middlewares.SetMiddlewareLogger(
// 				middlewares.SetMiddlewareJSON(route.Handler)),
// 			).Methods(route.Method)
// 		}
// 	}
// 	return r
// }

// func New() *mux.Router {
// 	r := mux.NewRouter().StrictSlash(true)
// 	return SetUpRoutesWithMiddlewares(r)
// }

// func LoadCORS(r *mux.Router) http.Handler {
// 	headers := handlers.AllowedHeaders([]string{"X-Request", "Content-Type", "Location", "Entity", "Authorization"})
// 	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
// 	origins := handlers.AllowedOrigins([]string{"*"})

// 	return handlers.CORS(headers, methods, origins)(r)
// }
