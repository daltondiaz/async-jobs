# Async Jobs

This project has the goal to "manage" async jobs, when I say manage is not executing the one function
if it is running and can execute many jobs asynchronously.

## Problem/Solution

When I'm using Php sometimes we need to execute tasks or jobs in the background, but the concept of Threads
was not implemented in a simple way, and sometimes we need to use many other tools to do that, for example: Php send a message to one queue in RabbitMQ, and Supervisord call other Php file who read the message of queue and run something stuff in one process.

For solving this, I thought in using the power of Go Lang to create a simple application to solve.

# Required

Go lang
``

# Getting Started

If it's your first time run

`go mod tidy`

and to run the project

`go run.`
