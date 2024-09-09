# Example-3

Example number 3 from Chapter 2 of [Bootstrapping Microservices](https://www.bootstrapping-microservices.com).

You need Go installed to run this.

To start the microservice:

    PORT=4000 go run main.go

Then point your browser at http://localhost:4000

Alternatively, with [air](https://github.com/air-verse/air "air") installed, start the microservice with:

    PORT=4000 air

This will create a `./tmp` path to build main into during development, so add it to your `.gitignore` file.
