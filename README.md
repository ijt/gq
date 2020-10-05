gq
==

GraphQL Querier

Prerequisite
------------
Install Go.

Installing
----------
`go get github.com/ijt/gq`

Example
-------
```
gq -ep=https://rickandmortyapi.com/graphql -q='
  query {
    characters(page: 0){
      results {
        name
      }
    }
  }
'
```

