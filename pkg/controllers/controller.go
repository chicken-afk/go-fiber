package controllers

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserController struct {
	// Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		// Inject services
	}
}

type BookController struct {
	//Dependent services
}

func NewBookController() *BookController {
	return &BookController{
		//Inject services
	}
}
