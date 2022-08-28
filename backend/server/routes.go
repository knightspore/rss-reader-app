package server

func (s *Server) Routes() {
	s.Router.HandleFunc("/api/user/create", s.HandleUserCreate())
// 	s.Router.HandleFunc("/api/subscription/create", s.HandleSubscriptionCreate())
// 	s.Router.HandleFunc("/api/subscription/get", s.HandleSubscriptionGet())
// 	s.Router.HandleFunc("/api/readinglist/get", s.HandleReadingListGet())
// 	s.Router.HandleFunc("/api/article/read", s.HandleArticleRead())
}

