# Go API client for openfigi

A free & open API for FIGI discovery.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.4.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://openfigi.com/api](https://openfigi.com/api)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./openfigi"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.openfigi.com/{basePath}*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**MappingPost**](docs/DefaultApi.md#mappingpost) | **Post** /mapping | 
*DefaultApi* | [**MappingValuesKeyGet**](docs/DefaultApi.md#mappingvalueskeyget) | **Get** /mapping/values/{key} | 

## Documentation For Models

 - [FigiResult](docs/FigiResult.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [MappingJob](docs/MappingJob.md)
 - [MappingJobResult](docs/MappingJobResult.md)
 - [MappingJobResultFigiList](docs/MappingJobResultFigiList.md)
 - [MappingJobResultFigiNotFound](docs/MappingJobResultFigiNotFound.md)
 - [OneOfMappingJobIdValue](docs/OneOfMappingJobIdValue.md)

## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

support@openfigi.com