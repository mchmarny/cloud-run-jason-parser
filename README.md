# cloud-run-jason-parser

Simple [Cloud Run](https://cloud.google.com/run/) service that parses posted JSON using provided select query and returns matching results.

The currently implemented services are:

* `POST /find` - Returns all matching elements anywhere in the posted content based on provided select query
* `POST /select` - Selects specific element from the posted content based on a fully-qualified query

## Example

You can invoke the service using HTTP to select on specific element of that JSON

```shell
curl -X POST \
     -H "Select-query: contributions" \
     -H "Content-Type: application/json" \
     -d '{"login": "test","id": 123456,"type": "User","contributions": 551}' \
     https://json-selector-2gtouos2pq-uc.a.run.app/select
```

The resulting JSON response from the service would be `551`.

## Build the image

Start by [building a container image](bin/image) by submitting job to Cloud Build using the included [Dockerfile](./Dockerfile) and results in versioned, non-root container image URI which will be used to deploy your service to Cloud Run.

```shell
bin/image
```

## Create IAM account

It's a good practice to run Cloud Run service under a specific IAM account. [Create a service account](bin/account) to configure least privilege IAM service account who's identity the deployed Cloud Run service will run under.

```shell
bin/account
```

## Deploy Service

Finally [deploy public Cloud Run service](bin/deploy) configured with the previously configured service account identity container image.

```shell
bin/deploy
```

## Test

You can use one of the provided sample files to test your newly deployed service. Start by first capturing the Cloud Run URL

```shell
export SERVICE_URL=$(gcloud beta run services describe json-selector \
    --region us-central1 \
    --format="value(status.domain)")
```

Now you can CURL on that service with the provided sample file which is a result of the GitHub public API query on the [Knative Serving repository](https://github.com/knative/serving) contributors

```shell
curl -X POST \
     -H "Select-query: login" \
     -H "Content-Type: application/json" \
     -d "@sample/github-repo-contributors.json" \
     "${SERVICE_URL}/find"
```
The result should be `test`.

## Cleanup

To cleanup all resources created by this sample execute

```shell
bin/cleanup
```

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.