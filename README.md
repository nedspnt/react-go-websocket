# react-go-websocket

Run backend server
```
go run main.go
```
This server runs at localhost port 8080

Then run frontend server
```
cd websocket-react-app
npm start
```
React app by default runs at localhost:3000

Put any text in the input form and submit.
You will be able to see the submitted text appear in:
(1) Backend server's Terminal
(2) Browser console
(3) Web interface

TO DO: 
1. Explore how it works when multiple client servers are running and sending messages (broadcasting)
2. Create the feature where all users can see the broadcasted message
