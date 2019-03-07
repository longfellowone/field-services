package main

import "fmt"

func main() {
	s := InitializeServices() // Inject *sql.DB here

	s.CreateAuthor(0, "name", "name@email.com") // Example http/grpc/graphql endpoints
	s.CreatePost(1, 0, "title", "body")
	s.Publish(1)
	s.FindPost(1)
}

func InitializeServices() *Server {
	ar := NewAuthorRepository() // Create repositories, inject *sql.DB here
	pr := NewBlogPostRepository()

	us := NewUserService(ar) // Inject repositories into services
	bs := NewBloggingService(ar, pr)

	return New(us, bs) // Inject services into and return *Server
}

// Server
// package server
type Server struct {
	usvc UserService
	bsvc BloggingService
}

func New(usvc UserService, bsvc BloggingService) *Server {
	return &Server{
		usvc: usvc,
		bsvc: bsvc,
	}
}

func (s *Server) CreateAuthor(id int, name, email string) { // Create an Author
	s.usvc.CreateAuthor(id, name, email)
	fmt.Println("Saved Author:", id, name, email)
}

func (s *Server) CreatePost(id, authorid int, title, body string) { // Create a BlogPost
	s.bsvc.CreatePost(id, authorid, title, body)
	fmt.Println("Saved Post:", id, authorid, title, body)
}

func (s *Server) Publish(id int) { // Publish BlogPost
	s.bsvc.Publish(id)
	fmt.Println("Published post:", id)
}

func (s *Server) FindPost(id int) { // Finds a complete blog post with authors name
	post := s.bsvc.FindPost(id)
	fmt.Println("Found Post:", post)
}

// UserService
// package users
type UserService interface {
	CreateAuthor(id int, name, email string)
}

type authorRepository interface {
	Find(id int) *Author
	Save(a *Author)
}

type userService struct {
	ar authorRepository
}

func NewUserService(ar authorRepository) *userService {
	return &userService{
		ar: ar,
	}
}

func (s *userService) CreateAuthor(id int, name, email string) {
	author := CreateAuthor(id, name, email)
	s.ar.Save(author)
}

// BloggingService
// package blogging
type BloggingService interface {
	CreatePost(id, authorid int, title, body string)
	Publish(id int)
	FindPost(id int) *BlogPostRead
}

// Redeclare this in your BloggingService package
//type authorRepository interface {
//	Find(id int) *Author
//	Save(a *Author)
//}

type blogPostRepository interface {
	Find(id int) *BlogPost
	Save(b *BlogPost)
}

type bloggingService struct {
	ar authorRepository
	pr blogPostRepository
}

func NewBloggingService(ar authorRepository, pr blogPostRepository) *bloggingService {
	return &bloggingService{
		ar: ar,
		pr: pr,
	}
}

func (s *bloggingService) CreatePost(id, authorid int, title, body string) {
	post := CreateBlogPost(id, authorid, title, body)
	s.pr.Save(post)
}

func (s *bloggingService) Publish(id int) {
	post := s.pr.Find(id)
	post.Publish()
	s.pr.Save(post)
}

func (s *bloggingService) FindPost(id int) *BlogPostRead {
	post := s.pr.Find(id)
	author := s.ar.Find(post.AuthorID) // You could also call an external service here to get this information
	return &BlogPostRead{
		BlogPostID: post.BlogPostID,
		AuthorName: author.Name,
		Title:      post.Title,
		Body:       post.Body,
		Published:  post.Published,
	}
}

type BlogPostRead struct { // BlogPost read model
	BlogPostID int
	AuthorName string
	Title      string
	Body       string
	Published  bool
}

// Core domain & business logic
// package
type Author struct {
	AuthorID int
	Name     string
	Email    string
}

func CreateAuthor(id int, name, email string) *Author {
	return &Author{
		AuthorID: id,
		Name:     name,
		Email:    email,
	}
}

type BlogPost struct {
	BlogPostID int
	AuthorID   int
	Title      string
	Body       string
	Published  bool
}

func CreateBlogPost(id, authorid int, title, body string) *BlogPost {
	return &BlogPost{
		BlogPostID: id,
		AuthorID:   authorid,
		Title:      title,
		Body:       body,
		Published:  false,
	}
}

func (b *BlogPost) Publish() {
	b.Published = true
}

// AuthorRepository
// package
type AuthorRepository struct {
	author map[int]*Author
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{author: make(map[int]*Author)}
}

func (r *AuthorRepository) Find(id int) *Author {
	return r.author[id]
}

func (r *AuthorRepository) Save(a *Author) {
	r.author[a.AuthorID] = a
}

// BlogPostRepository
// package
type BlogPostRepository struct {
	post map[int]*BlogPost
}

func NewBlogPostRepository() *BlogPostRepository {
	return &BlogPostRepository{post: make(map[int]*BlogPost)}
}

func (r *BlogPostRepository) Find(id int) *BlogPost {
	return r.post[id]

}

func (r *BlogPostRepository) Save(b *BlogPost) {
	r.post[b.BlogPostID] = b
}
