# Go template: Gin Api
This repository contains a template to get started with Gin framework. 

## Features

- A configured Gin server ready to run
- An error handler to convert errors to ProblemDetails
- A Todo handler with sample routes 

## Getting started

You can start a new project with this template using [gonew](https://go.dev/blog/gonew).

```shell
gonew github.com/edirand/go-template-gin-api <destination_module>
```

Once the project ready, run the following command to initialize the project:
```shell
make setup
```

The script will replace all the references to go-template-gin-api with your project name and replace this readme with 
a default one.

## Roadmap

- [] Add configuration for the server based on environment or config files

## Credits

- [meysamhadeli](https://github.com/meysamhadeli) for his problem-details package that has been forked and 
a little modified to provide problem details with gin in this template
- [trstringer](https://github.com/trstringer) for his [article](https://trstringer.com/go-project-templates/) 
on gonew and his setup script