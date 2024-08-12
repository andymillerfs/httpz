# Httpz: http without the hassle

## What this package does for you

It simply leverages the http, json/encoding, and io packages to streamline the building, parsing, and sending of http requests.
If you have a struct in the shape of your desired response, and you want to treat any other shaped response as an error, this package offers a very quick way to do that.

## Instructions
Step 1: Build a struct in the shape of your desired response

```
type ExampleResponse struct {
    Foo string `json:"foo"`
    Bar string `json:"bar"`
}
```

Step 2: Construct your request

```
req := httpz.GetRequest{
    URL: "http://localhost:8080/example",
    Headers: []httpz.Header{
        httpz.Header{ Key: "Authorization", Value: "Bearer " + auth_token}
    }
}
```

Step 3: Retrieve a response or an error

```
resp, err := httpz.SendGetRequest[ExampleResponse](req)
if err != nil {
    // handle error
}
```

Step 4: Profit

```
foo := resp.Foo
```

In addition to the errors returned from the send functions, the specific step at which the error occurred will also be logged (i.e. request parsing, status check, response parsing, etc.) along with some useful information about the unexpected data causing the failure.