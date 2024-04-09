# Goformix

A lightweight and powerful form validation library designed for backend usage when parsing data from a JSON payload. It is still in development phase and is **not yet production ready**.

Goormix has not been tested yet. Suggestions and contributions are appreciated!

# Installation
```bash
go get -u github.com/laughing-nerd/goformix
```

# Examples (More examples incoming!)
- [fiber example](https://github.com/laughing-nerd/goformix/blob/main/examples/fiber_ex.go)

# Supported Types
- string
- number
- boolean
- email
- url
- phone
- date
- array
- embedded (Nested struct)
- []embedded (Array of nested structs)

# How to use the types
 Using goformix is super simple. You can checkout examples for more information. When you declare a struct, just add the tag `goformix:"<TYPE>"` to a field and goformix will take care of everything! The types have certain parameters which can be used for further customization. The parameters are separated by '|'.

Type | Parametes | Example
-|-|-
string| min \| max | goformix:"string\|10\|20"
number| N/A | goformix:"number"
boolean| N/A | goformix:"boolean"
email| domain1 \| domain2 \| ... | goformix:"email\|gmail.com\|example.com"
url| domain1 \| domain2 \| ... | goformix:"url\|www.linkedin.com\| ..."
phone| country | goformix:"phone\|India"
date| N/A | goformix:"date"
array| N/A | goformix:"array"
embedded| N/A | goformix:"embedded"
[]embedded| N/A | goformix:"[]embedded"
