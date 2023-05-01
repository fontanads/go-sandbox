# Go Lang Study Plan

After a serious discussion with an AI generative model, we've decided [Go](https://go.dev/) was a good option as the next programming language I am adding to my skill set.  

Here's the study plan my buddy GPT suggested:

# Study Plan: Learning Go

## Resources

- [A Tour of Go](https://tour.golang.org/welcome/1): A great place to start learning the basics of Go, with interactive examples and exercises.
- [Effective Go](https://golang.org/doc/effective_go.html): A document that covers the idiomatic way of writing Go code.
- [The Go Programming Language Specification](https://golang.org/ref/spec): The official specification of the Go programming language.
- [Go by Example](https://gobyexample.com/): A collection of examples that demonstrate how to use various features of Go.

**Books:**
- [The Go Programming Language](https://www.gopl.io/) by Alan A. A. Donovan and Brian W. Kernighan
- [Introducing Go: Build Reliable, Scalable Programs](https://www.oreilly.com/library/view/introducing-go/9781491941997/) by Caleb Doxsey
- [Concurrency in Go: Tools and Techniques for Developers](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) by Katherine Cox-Buday

**Online Courses:**
- [Learn Go Programming - Golang Tutorial for Beginners](https://www.udemy.com/course/learn-go-programming/)
- [Programming with Google Go Specialization](https://www.coursera.org/specializations/google-golang)


## Topics to Learn

1. Basics of Go programming
2. Variables, types, and operators
3. Control flow statements
4. Functions and methods
5. Pointers and memory management
6. Concurrency and parallelism
7. Packages, modules, and the Go toolchain
8. Error handling
9. Web programming with Go
10. Testing and debugging

## Sample Project

> For a sample project, I suggest building a web application using Go. Here are some possible features you could include:

1. A REST API that provides endpoints for creating, reading, updating, and deleting data from a database
2; Authentication and authorization using JWT tokens
3. Integration with cloud-based services like AWS or Google Cloud Platform
4. Use of goroutines and channels to handle concurrent requests
5. Use of the popular Gorilla toolkit for web development in Go

> By building a web application, you can gain experience with several important aspects of Go programming, including web development, concurrency, and integration with external services. You can also use this project as an opportunity to experiment with deploying your application using Kubernetes or other cloud-native technologies.

> To apply what you've learned, try building a simple RESTful API that exposes a machine learning model for predicting the price of used cars. Here are some high-level requirements:

- The API should be built using the [gin-gonic](https://github.com/gin-gonic/gin) web framework.
- The API should accept a JSON payload containing the car's attributes (e.g., make, model, year, mileage, etc.).
- The API should return a JSON payload containing the predicted price of the car.
- The machine learning model can be built using any library or framework you're comfortable with (e.g., scikit-learn, TensorFlow, etc.).
- You should include a README file that explains how to run the API and how to use it.

## Additional Tips

- Try to code along with the examples from the resources listed above.
- Practice what you've learned by working on small coding challenges.
- Join online communities like Reddit's r/golang or Gophers Slack group to ask questions and learn from experienced Go developers.
