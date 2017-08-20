# Go service with Gopher Sauce

This is a tutorial to help you recreate the project. This service will store an `int` in memory and increase it on each request.

If you don't have Gopher Sauce,  `go get github.com/cheikhshift/gos`

##  Start project

	$ mkdir new-project
	cd new-project
	gos --make


## Create service

Create a new `service.go` file in your project directory. 

### Add the following lines to `service.go`

1. Declare package :

		package main 

2. Declare static `int` :
	
		var Counter = 0

3. Declare `increase` function :

		func Increase(){
		Counter++
		}

## Link Go service to internet 	🕶

1. Open your `gos.gxml` file.
2. Add this end(point) tag within your `endpoints` tag.
		
		<end path="/service" type="GET" ></end>

3. Call the increase function within your `end` tag.

		Increase()

3. Set content type :

		w.Header().Set("Content-Type",  "text/html")

4. Display the current value of the counter :

		w.Write([]byte(fmt.Sprintf("This is the %v request", Counter) ))
	
## Test service
Run `gos --run` and visit : [localhost:8080/service](localhost:8080/service)	
