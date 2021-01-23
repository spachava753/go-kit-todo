# go-kit-todo

This is a simple example of a todo microservice using [go-kit](https://github.com/go-kit/kit).

# Thoughts

While Go-kit is great, it does have some drawbacks. Its very verbose, and has a high upfront cost for creating services. I wouldn't really suggest using the library for prototyping. Moving even further, in a containerized and cloud-native world, we don't need to concern ourselves with circuit-breaking, load-balancing, metrics, etc, since our platform can handle that. Of course, some things are still necessary like application logging and service level metrics, but things like circuit breaking and load-balancing is done by platforms like AWS or Kubernetes. Kubernetes inherently gives you service discovery and simple load-balancing, but taking things further with service meshes, we can utilize service to service encryption, reverse-proxying, circuit breaking, metrics, and distributing tracing. As such, in 2020, we don't need to concern ourselves with such matters.

Yet Go-kit introduces an interesting, but well-known concept of middleware. We would write our core business logic, and test it, and we then we can wrap our business logic with extra functionality like application logging. We can "append" functionality without actually modifying previous layer. This is a powerful concept, as we can separate functionality through layers, and gives the flexibility to enable and disable certain layers at startup, perhaps even during runtime.

Another major issue with Go-kit is that it translated endpoints to stdlib `http.Handler`. This sort of limits us in the ways we can service our app (what if I want to use [fiber](https://github.com/gofiber/fiber)?). Everything past the Service layers are dedicated to handling incoming requests and responses, with can change from library to library, framework to framework, architecture to architecture, platform to platform. As such, we can still use Go-kits concepts of Service layers, but wire up our transports separately, based on the end-coder's implementation. Notably, most web servers and router Golang libraries nowadays come with their own ways to add middleware. While Go-kit tries generalize transport middleware across transport protocols (http, grpc, serverless, etc.), I think it's completely fine to use an implementation specific middleware.