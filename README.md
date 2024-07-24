# About
A simple REST API implemented in Go that performs basic CRUD (Create, Read, Update, Delete) operations on a list of students

# How to run locally
- Clone the repo on your local machine<br/>
- Run Ollama server with llama3<br/>
- Change the ```OLLAMA_BASE_URL``` env variable in the env file if your Ollama server is not running on the default http://localhost:11434<br/>
- Change the ```PORT``` env variable if you want, the default is 4000<br/>
- Run ```make api/run``` in the cmd<br/>


# Summary
Summary Generation does not work on the deployed service. It only works locally as it requires Ollama on your machine. 
If you try to access the summary endpoint on the hosted service, it will return an error.

# Endpoints
Postman Doc- https://documenter.getpostman.com/view/26059341/2sA3kXDfnA

# Architecture
I have implemented a 3 layer architecture (repository pattern) along with dependency injection to make the application more robust and maintainable. There are 3 layers - handler, service and repository. The respository layer provides access to and performs operations on the data, the service layer is used for business logic like validating requests and integrating third party services in this case Ollama, the handler layer is the topmost layer, which accepts the requests performs some basic tasks like extracting query params and request body and then calls the respective method of the service layer which performs the business logic and then calls the repository layer. In this way the actions moves from top to bottom and data moves from bottom to top which makes it very easy to add/modify functionality in the future
## Dependency Injection
The ```App``` struct in the main.go file acts as a dependency injection container, it encapsulates all dependendcies and configs making it clear what components the application depends on. By injecting dependencies through the App struct, the application layers become decoupled from each other as each component (like StudentService) does not need to know how to create its dependencies. it just receives them through the App struct, which makes the components more modular and testable.
## Error Handling
The ```APIError``` struct defined in the apierror package represents the API Error which is returned whenever there is an error. It implements the Error interface, so it can be used interchangeably with the golang error package as well. Whenever an error occurs in our app, like in the repository layer or during request validation, it returns a golang error, we handle this error and create our own ```APIError``` from this error, instead of returning the golang error to the caller, we return this ```APIError``` instead, until it reaches the handler layer where we return this ```APIError``` as a JSON response, we do this so that we can easily create json response from the error and the error handling logic is modularized  
## Validation
The ```validators``` package contains the validation logic. In the ```commons.go``` I have defined common validation functions like email validation, integer range validation so that we can reuse them whenever needed. The ```Studentvalidator``` function takes the ```Student``` struct and validates it using the common validation functions. This way the validation logic becomes reusable.





