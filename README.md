# Basic User Authentication Infrastructure

This is a basic user authentication infrastructure built using AWS's CDK and written in Go.
The infrastructure for this project consists of AWS's API Gateway, Lambda, and DynamoDB.
Included are API endpoints that handle registering users, logging the in, and also simple middleware using JSON Web Token (JWT) to grant users access tokens to acesss protected routes for necessary data.
The current error handling should tackle general cases and output generally understandable error outputs, but is not very extensive as of right now.
This project was more intended for learning Go and working with AWS's infrastructure as code principles but I intend to expanded upon it for my future projects.
