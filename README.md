# go-go-url-go

![Actions Workflow](https://github.com/joaomarcuslf/go-go-url-go/workflows/go/badge.svg)
![Actions Workflow](https://github.com/joaomarcuslf/go-go-url-go/workflows/CodeQL/badge.svg)

This is and URL shortener repo written in Go, using Redis and Docker.

It is deployed in Google Cloud Platform.

## Getting Started

1. Copy ```sample.env``` to ```.env``` and rename the variables if you need
2. Build the images and run the containers:

```sh
make docker-start
```

You can stop the docker by running: `make docker-stop`

## Running Tests

```sh
make test
```

## Deploying

Make sure you have `heroku cli`in you `path`, and access to the project.

```sh
make deploy version=4
```

## Collaborators

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/joaomarcuslf">
        <img src="https://avatars.githubusercontent.com/u/53450523?v=4" width="100px;" alt="Joaomarcuslf's Github picture"/><br>
        <sub>
          <b>joaomarcuslf</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

[⬆ Scroll top](#go-go-url-go)<br>
