![go-logo](https://go.dev/images/go-logo-blue.svg)

# GO first web app

This repo is my first working on go language.

There is a few branch to execute each step from [Go Web Examples](https://gowebexamples.com/) web page.

## Table of Contents

- [Routers](#routers)
- [Templates](#templates)
- [Forms](#forms)
- [Middlewares](#middlewares)

### Routers

    Go’s net/http package provides a lot of functionalities for the HTTP protocol. One thing it doesn’t do very well is complex request routing like segmenting a request url into single parameters. Fortunately there is a very popular package for this, which is well known for the good code quality in the Go community. In this example you will see how to use the gorilla/mux package to create routes with named parameters, GET/POST handlers and domain restrictions.

### Templates

    Go’s html/template package provides a rich templating language for HTML templates. It is mostly used in web applications to display data in a structured way in a client’s browser. One great benefit of Go’s templating language is the automatic escaping of data. There is no need to worry about XSS attacks as Go parses the HTML template and escapes all inputs before displaying it to the browser.

### Forms

    This example will show how to simulate a contact form and parse the message into a struct.

### Middlewares

    A middleware in itself simply takes a http.HandlerFunc as one of its parameters, wraps it and returns a new http.HandlerFunc for the server to call.

    Here we define a new type Middleware which makes it eventually easier to chain multiple middlewares together. This idea is inspired by Mat Ryers’ talk about Building APIs. You can find a more detailed explaination including the talk here.

    This snippet explains in detail how a new middleware is created. In the full example below, we reduce this version by some boilerplate code.
