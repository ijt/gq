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
gq -ep=https://countries.trevorblades.com/ -q='
  query {
    countries {
      name
    }
  }
'
```

