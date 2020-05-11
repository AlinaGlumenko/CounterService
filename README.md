# CounterService
This application contains three pages, two of which refer to the GO-service, which in turn considers the number of users visiting the page.
This application does this thanks to a script that runs in a browser on the client side. This js-script saves the state in cookies to know if the user has visited this page. If the user visited the page for the first time, then the state is saved in cookies and a request is sent to the service to count +1 visits of a specific page. GET-request to the service should include a link to access the current page, which in turn is the key in the Redis database. Also an optional parameter is a visit label. If the tag is present, then the user visited the page for the first time, if not, then the application turned to the service to simply find out the number of visits to a particular page.

## Built With
* [Express](https://www.npmjs.com/package/express) - The web framework for Node.js.
* [Gorilla/Mux](https://github.com/gorilla/mux/) - A request router and dispatcher for matching incoming requests to their respective handler.
* [Redis](https://redis.io/) - Used as database.
* [Redigo](https://github.com/gomodule/redigo) - A Go client for the Redis database.

## Installation
Use the commands below to install packages:
#### Install Redigo
```
go get github.com/gomodule/redigo/redis
```
#### Install gorilla/mux
```
go get github.com/gorilla/mux
```
#### Install NPM packages
```
npm install
```

## Deployment
Use the commands below to run:
##### Run this command in the terminal of ```.../app/``` directory:
```
node app.js
```
##### Run this command in the terminal of ```.../service/``` directory:
```
go run main.go
```

## TODO list
- [x] Сreate a service that counts users which opened current page of website
- [x] Display counter on the same page
- [ ] Tests
- [ ] Check if it is a web-link on the server side.
- [ ] Do not save the link to the page as a key in the Redis repository (cause of URL parameters).

## License
This project is licensed under the MIT License - see the [LICENSE.md](./LICENSE) file for details
