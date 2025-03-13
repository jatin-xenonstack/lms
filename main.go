package main

import (
	"fmt"
	"library-management1/database"
	"library-management1/handlers"
	"library-management1/middlewares"

	_ "library-management1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.InitDB()
	fmt.Println("Hello!")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are on the home page"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.POST("/getdata", handlers.Test)
	// router.DELETE("/del", handlers.Del)
	// @title Library Management System API
	// @version 1.0
	// @description This is a sample server for a library management system.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host localhost:8000
	// @BasePath /

	// @Summary Sign up a new user
	// @Description Create a new user account
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param user body handlers.User true "User"
	// @Success 200 {object} handlers.UserResponse
	// @Router /auth/signup [post]

	// @Summary Log in a user
	// @Description Authenticate a user and return a token
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param credentials body handlers.Credentials true "Credentials"
	// @Success 200 {object} handlers.TokenResponse
	// @Router /auth/login [post]

	// @Summary Get all libraries
	// @Description Retrieve a list of all libraries
	// @Tags library
	// @Produce json
	// @Success 200 {array} handlers.Library
	// @Router /getalllibrary [get]

	// @Summary Search for a book
	// @Description Search for a book by title or author
	// @Tags book
	// @Produce json
	// @Param search path string true "Search"
	// @Success 200 {array} handlers.Book
	// @Security ApiKeyAuth
	// @Router /book/{search} [get]

	// @Summary Get all books
	// @Description Retrieve a list of all books
	// @Tags book
	// @Produce json
	// @Success 200 {array} handlers.Book
	// @Security ApiKeyAuth
	// @Router /book/all [get]

	// @Summary Create a new library
	// @Description Create a new library
	// @Tags owner
	// @Accept json
	// @Produce json
	// @Param library body handlers.Library true "Library"
	// @Success 200 {object} handlers.LibraryResponse
	// @Security ApiKeyAuth
	// @Router /owner/create-library [post]

	// @Summary Assign an admin to a library
	// @Description Assign an admin to a library
	// @Tags owner
	// @Accept json
	// @Produce json
	// @Param admin body handlers.AdminAssignment true "Admin Assignment"
	// @Success 200 {object} handlers.AdminResponse
	// @Security ApiKeyAuth
	// @Router /owner/assign-admin [post]

	// @Summary Register a user to a library
	// @Description Register a user to a library
	// @Tags user
	// @Accept json
	// @Produce json
	// @Param registration body handlers.Registration true "Registration"
	// @Success 200 {object} handlers.RegistrationResponse
	// @Security ApiKeyAuth
	// @Router /user/register [post]

	// @Summary Insert a new book
	// @Description Insert a new book into the library
	// @Tags admin
	// @Accept json
	// @Produce json
	// @Param book body handlers.Book true "Book"
	// @Success 200 {object} handlers.BookResponse
	// @Security ApiKeyAuth
	// @Router /admin/insert-book [post]

	// @Summary Update book copies
	// @Description Update the number of copies of a book
	// @Tags admin
	// @Accept json
	// @Produce json
	// @Param book body handlers.BookUpdate true "Book Update"
	// @Success 200 {object} handlers.BookResponse
	// @Security ApiKeyAuth
	// @Router /admin/update-book [post]

	// @Summary Request an event
	// @Description Request an event in the library
	// @Tags user
	// @Accept json
	// @Produce json
	// @Param event body handlers.EventRequest true "Event Request"
	// @Success 200 {object} handlers.EventResponse
	// @Security ApiKeyAuth
	// @Router /user/request [post]

	// @Summary See all requests
	// @Description See all event requests
	// @Tags admin
	// @Produce json
	// @Success 200 {array} handlers.EventRequest
	// @Security ApiKeyAuth
	// @Router /admin/see-request [get]

	// @Summary Handle a request
	// @Description Handle an event request
	// @Tags admin
	// @Accept json
	// @Produce json
	// @Param request body handlers.RequestHandle true "Request Handle"
	// @Success 200 {object} handlers.RequestHandleResponse
	// @Security ApiKeyAuth
	// @Router /admin/handle-request [post]
	router.POST("/auth/signup", handlers.CreateUser)
	router.POST("/auth/login", handlers.Login)
	router.GET("/getalllibrary", handlers.GetAllLibrary)
	router.GET("/book/:search", middlewares.CheckAuth(), handlers.SearchBook)
	router.GET("/book/all", middlewares.CheckAuth(), handlers.GetAllBook)
	// router.GET("/user/profile", middlewares.CheckAuth(), handlers.GetUserProfile)
	router.POST("/owner/create-library", middlewares.CheckAuth(), handlers.CreateLibrary)
	router.POST("/owner/assign-admin", middlewares.CheckAuth(), handlers.AssignAdmin)
	router.POST("/user/register", middlewares.CheckAuth(), handlers.Register)
	router.POST("/admin/insert-book", middlewares.CheckAuth(), handlers.InsertBook)
	router.POST("/admin/update-book", middlewares.CheckAuth(), handlers.UpadateBookCopies)
	router.POST("/user/make-request", middlewares.CheckAuth(), handlers.RequestEvent)
	router.GET("/admin/see-request", middlewares.CheckAuth(), handlers.SeeRequest)
	router.POST("/admin/handle-request", middlewares.CheckAuth(), handlers.HandleRequest)

	router.Run("localhost:8000")
}
