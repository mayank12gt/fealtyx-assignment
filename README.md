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





