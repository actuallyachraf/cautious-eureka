# Project Specification

The coding challenge functionnal specification requires the following :

- Develop a REST microservice that lists the languages used by the top 100 trending
projects on Github.
- For every language, you need to calculate the following attributes :

  - Number of repos using the language.
  - The list of repos using the language.

## Technical Specification

The project was implemented using [Go](https://golang.org) and follows the following
architecture.

The data feed used is Github's own [API](https://developer.github.com/v3/).

We are targeting mainly the following route :

```go
https://api.github.com/search/repositories?q=created:>{date}&sort=stars&order=desc
```

Where date is specified as 30 days before today.

To separate concerns and provide a pragmatic API the following we design the following routes :

- The top trending projects in the last 30 days : ```/api/trending```
  - Filter by language : ```/api/trending/{language}```
