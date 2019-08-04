# GopherGold
A platform for joke crowd sourcing which lets the joke creators rate their own jokes 

## Getting started

To get this web project running locally:

- Clone this repo
- To run this project you must have [Go](https://golang.org/dl/) installed.

## Testing the app

- Access the root directory which contains the `Makefile`
- Type `make test` in the terminal

## Running the app

- Access the root directory which contains the `Makefile`
- Type `make run` in the terminal
- Local web server will use the standard port 3000.
  - [http://localhost:3000](http://localhost:3000)

## General Functionality

- The page will load a random joke from the database
- Users can submit their own jokes
- If a joke falls outside of the score range(1 to 10) it will be rejected
- Users will be informed if their joke is accepted or rejected

## Future Improvements

- Full re-design of the user interface
- Create a public server to host all submitted jokes
- Create an API for other sites to embed our jokes and our joke submission forms
- Allow optional images to be posted along with each joke

`Note: This application has been forked from [here](https://github.com/mo-hit/takehome-coding-challenge)`
