# Backend Coding Challenge

The following repository provides an implementation of the [gemography backend coding challenge](https://github.com/gemography/backend-coding-challenge/blob/master/README.md).

## Usage

### Building

```sh
 git clone git@github.com:actuallyachraf/cautious-eureka.git
 cd cmd/
 go build -o app
```

### Querying

- Fetch the top trending repositories in the last 30 days :

    ```sh
    curl localhost:3000/api/trending
    ```

- Fetch the number of repos and their URLs by language

    ```sh
    curl localhost:3000/api/trending/Python
    ```
