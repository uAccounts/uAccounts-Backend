# uAccounts-Backend

- To run this code first you need to install Dependencies..

## Dependencies

Just run this command to the vscode terminal when you cloned the repository.

` go mod download `

if it is not working;

To the terminal you should write those commands separately

- go get "github.com/gorilla/mux"
- go get "github.com/rs/cors"

- Then you can start the backend with ` go run main.go `
- It will start on http://localhost:8080/
- You can check the api with adding this `/systemhealthcheck` endpoint to the localhost from Postman. 

## Purpose and Process of the Study

- We are creating a platform for League of Legends lovers.
- We open new accounts using a bot app and we offer these accounts to users by raising them to level 30.