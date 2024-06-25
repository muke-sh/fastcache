package main

import "net"


type Config struct {
 ListenAddress string 
}

type Server struct {
 Config 
 ln net.Listener
}

func NewServer(cfg Config) *Server {
  return &Server{
   Config: cfg,
  }
}

func (s *Server) Start() error {
  
  ln, err := net.Listen("tcp", s.ListenAddress)
 
  if err != nil {
    return err
  }

  s.ln = ln

  return nil
}

func main() {
  
}
