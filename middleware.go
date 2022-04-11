package scapler

import "net/http"

func (s *Scapler) SessionLoad(next http.Handler) http.Handler {
	s.InfoLog.Println("Session Loaded")
	return s.Session.LoadAndSave(next)
}
