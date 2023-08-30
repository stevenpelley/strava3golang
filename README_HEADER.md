# Golang bindings for Strava V3 API
This is the standard golang bindings for Strava's V3 API (current API as of 2023).  This is made available by Strava via Swagger with no modifications here, but it was sufficiently difficult to get working that I thoughts I'd post it on github for others, along with an explanation for why I did it this way (see the Makefile for a simple description of "what" I did).

For general use feel free to require this repos in golang.  I don't intend to update it in the future, but I also make no promises.  If it breaks then your recourse will be to clone and revert this repository and update your requires to your own repository.

Strava uses the same string values in multiple enums.  Golang, which has no built in support for enums, uses the value of these strings to generate static variable names.  The variables from these different enums then collide and produce compiler errors.  You can, of course, simply rename these variables to address the conflict.  I instead used openapi-generator (compatible with Swagger 2 which is what Strava requires) with its -p enumClassPrefix=true option to do this for me.  Openapi-generator complains that the provided swagger input is invalid, so I silence this with --skip-validate-spec.

To use this client you will first need to perform oauth2 authorization, create an http client with the oauth2 token, and assign this client into the strava HTTPClient struct.  I use github.com/stevenpelley/oauth2-accesstoken-golang to create a token for development use.

See example/athlete for an example.  This assumes you have the input oauth_client_config.json used in github.com/stevenpelley/oauth2-accesstoken-golang as well as the token.json that results.  These files are both listed in .gitignore to prevent commiting secrets.

What follows is the README generated from the strava swagger itself.

