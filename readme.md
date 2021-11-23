# Golang Pubsub with Firebase Emulator

This project presents a simple example of how to use Google PubSub with Golang.
For that, the firebase emulator was used, providing a simple way to emulate the resources.

## Project Concepts
The purpose here is simple. After run the application, an endpoint will be provided following the
example bellow:

`curl host:port?content=any`

Note that you can send any data as a query parameter. This data will be added as a content of the 
generated messages. When you call this endpoint, a message will be published, and subscribers will
consume it.
In this example, we are considering two subscribers. You can create more and more according to your 
test scenarios.

## Running the project

This project can be executed in two ways. If you just want run and see how the things work, after 
cloning, execute `docker-compose up` and all resources will be started.   

You can also run locally, using your favorite editor or a simple terminal. To do this, make
sure you have golang installed on your machine. This projects was created considering golang in your
`1.16.2`. Don't forget to check `.env.example` file to set the required environment variables.