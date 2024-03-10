# ASCII-ART-WEB

## Authors

@bisa

@sahahmed

## Project Description

* Ascii-art-web project involves creating and running a server.

* The server provides a web GUI version of an ascii-art program.

* The ascii-art program takes a string as input and generates a graphic representation using ASCII characters.

* The graphic representation is displayed on an HTML page through the web GUI.

* The program should handle inputs with numbers, letters, spaces, and special characters.

## Main page requirements 

* A text input field

* Radio buttons, select options, or any other mechanism to switch between different banners

* A button that triggers a POST request to the '/ascii-art' endpoint and displays the result on the page.

## Available Banners

* standard 

* shadow

* thinkertoy

## HTTP Endpoints

* GET /: This endpoint sends an HTML response, which serves as the main page. It is recommended to use Go templates to receive and display data from the server.

* POST /ascii-art: This endpoint is used to send data (text and a banner) to the Go server. To make the POST request, it is suggested to utilize form and other types of tags.

* The way to present the result from the POST request is flexible and depends on the designer's preference. Here are two recommended approaches:

* Display the result on the /ascii-art route after the POST request is completed. This means navigating from the home page to another page to view the result.
* Alternatively, display the result of the POST request on the home page itself by appending the results.

## Appropriate HTTP status codes

* If everything goes well, the response should have a status code of OK (200).

* If nothing is found, such as templates or banners, the response should have a status code of Not Found.

* For incorrect requests, the response should have a status code of Bad Request.

* In case of unhandled errors, the response should have a status code of Internal Server Error.

## To use the application

* Open the terminal and run the command: $ go run . main.go

* You can access the web page by clicking on the web browser icon located on the right side, and it will take you directly to the webpage.