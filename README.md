# Fake server json

This project is made with [Golang](https://go.dev/) for provide a simple fake server to responde data from the url

## Example

```json
{
    "url":"/url",
    "response":
        {
            "key":"result of response"
        }
}
```

## How to use

If you have already [make](https://www.gnu.org/software/make/manual/make.html) installed, it's only necessary to run this command: `make start_server`

## Feature

- [x] Read json file and make each route
- [ ] Reaload json file when project is already run
- [ ] Make variable parameter
- [ ] Output variable parameter inside response
- [ ] Log output
- [ ] Make a list of url callable
- [ ] Make simple logic to improve output of the message by code inside json