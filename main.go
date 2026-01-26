package main

import "fmt"

type Repository interface {
	Save()
	Get()
}

type sqlite struct {
	DSN string
}

func NewSqlite(dsn string) *sqlite {
	return &sqlite{
		DSN: dsn,
	}
}

func (s sqlite) Save() {
	fmt.Printf("Saved in sqlite in DSN: %v \n", s.DSN)
}

func (s sqlite) Get() {
	fmt.Printf("Take your result from sqlite in DSN: %v \n", s.DSN)
}

type postgres struct {
	Host string
	Port string
}

func NewPostgres(host, port string) *postgres {
	return &postgres{
		Host: host,
		Port: port,
	}
}

func (p *postgres) Save() {
	fmt.Printf("Save in postgres, host: %s, port: %s \n", p.Host, p.Port)
}

func (p *postgres) Get() {
	fmt.Printf("Get from postgres, host: %s, port: %s, ENJOY IT! \n", p.Host, p.Port)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SaveAll() {
	s.repo.Save()
}

func (s *Service) GetAll() {
	s.repo.Get()
}

func Prnt(val any) {
	switch v := val.(type) {
	case int:
		fmt.Print("Это инт! ", v)
	case string:
		fmt.Print("Это строка! ", v)
	default:
		fmt.Print("Я хз че эт:(")
	}
}

func main() {
	// pg := NewPostgres("localhost", "5432")
	// s := NewService(pg)
	// s.GetAll()
	// s.SaveAll()
	// 
	Prnt(3i+1)
}
