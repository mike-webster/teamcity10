# teamcity10

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

A Golang client for interacting with the teamcity v.10 REST API

## Install
`go get github.com/mike-webster/teamcity10`

## Usage
The functionality provided by this package depends on there being a few values present in the context that is provided to the methods:
- A base64 encoded representation of a valid set of teamcity credentials in the form of `username:password`
    - On OSx: `base64 <<< username:password`
- The Base URL for your build server
    - Ex: https://builds.mybuildserver.com

The values should be provided in the context using the ContextKeys found in the keys.go file.

#### Verifying Functionality
To make sure you are passing the values correctly, you can run the following commands (replace the filler values with your own). If it fails, there's something wrong:
```
TC_BASE={builds.example.com} TC_CREDS={testcreds} TEST_ID={testid} go test . -run TestSetup
```

For the value of `TEST_ID` you'll need to include the teamcity ID for a project that you can build as a test.

## Contributing
- If you'd like a feature, create an an issue and assign it to @mike-webster for approval.
    - PRs are accepted, but please don't start working on something without approval.
    - PRs will only be merged by @mike-webster at this time