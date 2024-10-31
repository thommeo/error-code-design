# Error Codes Documentation

This document is auto-generated. Do not edit manually.

## Error Code Format

All error codes follow the format: E<type><data> where:
- E: Fixed prefix identifying this as an error code
- type: Single base-36 character (0-9,A-Z) identifying the error code format
- data: Variable-length base-36 encoded data specific to each format

Base-36 encoding uses digits 0-9 and letters A-Z to pack more information into fewer characters while remaining human-readable.


## Tiny Format

Simplest possible error code format using just an error type value.

The format provides:
- Values from 0 to 1295 (00 to ZZ in base-36)
- Total of 1,296 possible unique error codes

The code is encoded as E0XX where:
- E: Fixed prefix
- 0: Fixed type identifier
- XX: Two base-36 characters encoding the error type (00-ZZ)

Examples:
- E000: Unknown error
- E001: Validation error
- E0ZZ: Maximum value (1295)

| Code | Type | Description | 
|----|----|----|
| E000 | unknown | Unknown error | 
| E001 | validation | Validation error | 
| E002 | not_found | Resource not found | 
| E003 | unauthorized | Unauthorized access | 
| E004 | bad_request | Bad request | 
| E0ZZ | max | Maximum error value (ZZ) | 



## Simple Format

Each error code is composed of two bytes encoded as follows:
- Class (8 bits): Identifies the error class (allows up to 256 distinct classes)
- ErrorType (8 bits): Identifies the specific error (allows up to 256 errors per class)

The format provides:
- Up to 256 different classes
- Up to 256 different error types per class
- Total of 65,536 possible unique error codes

The code is encoded as E<type><data> where:
- E: Fixed prefix
- type: 1 character in base-36 encoding the error type
- data: 4 characters in base-36 encoding the class and error type bits

Bit layout before encoding:
```
[CCCCCCCC][EEEEEEEE]
C: Class bits
E: ErrorType bits
```

| Code | Class.Type | Description | 
|----|----|----|
| E10000 | unknown.unknown | Unknown API error | 
| E10074 | api.unknown | Unknown API error | 
| E10075 | api.validation_error | API validation error | 
| E10076 | api.authorization_error | API authorization error | 
| E100E8 | jobs.unknown | Unknown job error | 
| E100E9 | jobs.database_query | Database query error in job | 
| E100EA | jobs.timeout | Job execution timeout | 
| E11EKF | max.max | Max error type number | 



## Simple 5-11 Format

Each error code is composed of 16 bits encoded as follows:
- Class (5 bits): Identifies the error class (allows up to 32 distinct classes)
- ErrorType (11 bits): Identifies the specific error (allows up to 2048 errors per class)

The format provides:
- Up to 32 different classes
- Up to 2048 different error types per class
- Total of 65,536 possible unique error codes

The code is encoded as E<type><data> where:
- E: Fixed prefix
- type: 1 character in base-36 encoding the error type
- data: 4 characters in base-36 encoding the class and error type bits

Bit layout before encoding:
```
[CCCCCEEE][EEEEEEEE]
C: Class bits (5)
E: ErrorType bits (11)
```

| Code | Class.Type | Description | 
|----|----|----|
| E30000 | unknown.unknown | Unknown error | 
| E301KW | http.unknown | Unknown HTTP error | 
| E301KX | http.bad_request | Bad request error (400) | 
| E301KY | http.unauthorized | Unauthorized error (401) | 
| E301KZ | http.forbidden | Forbidden error (403) | 
| E301L0 | http.not_found | Not found error (404) | 
| E31EKF | max.max | Maximum error type value | 



## App Component Format

Each error code is composed of 24 bits of data encoded as follows:
- App (4 bits): Identifies the application (allows up to 16 apps)
- Component (6 bits): Identifies the major component (allows up to 64 components per app)
- SubComponent (6 bits): Identifies the specific sub-component (allows up to 64 sub-components per component)
- ErrorType (8 bits): Identifies the specific error (allows up to 256 error types per sub-component)

The format provides:
- Up to 16 different applications
- Up to 64 different components per application
- Up to 64 different sub-components per component
- Up to 256 different error types per sub-component
- Total of 16,777,216 possible unique error codes (16 * 64 * 64 * 256)

The code is encoded as E<type><data> where:
- E: Fixed prefix
- type: 1 base-36 character encoding the type (2)
- data: 5 base-36 characters encoding the packed 24 bits

Bit layout before encoding:
```
[AAAACCCC][CCSSSSSS][EEEEEEEE]
A: App bits
C: Component bits
S: SubComponent bits
E: ErrorType bits
```

| Code | App.Component.SubComponent.Type | Description | 
|----|----|----|
| EA0MTQ8 | backend.handler.unknown.unknown | Unknown handler error | 
| EA0MTXD | backend.handler.users.validation_error | Input validation failed for user operation | 
| EA0MTXE | backend.handler.users.authorization_error | User lacks required permissions for operation | 
| EA0MU4H | backend.handler.records.validation_error | Input validation failed for record operation | 
| EA0MU4I | backend.handler.records.authorization_error | User lacks required permissions for record operation | 
| EA0MUBL | backend.handler.analytics.validation_error | Input validation failed for analytics operation | 
| EA0MUBM | backend.handler.analytics.authorization_error | User lacks required permissions for analytics operation | 
| EA0N6DC | backend.job.unknown.unknown | Unknown job error | 
| EA0N6KH | backend.job.sync.database_error | Database operation failed during sync | 
| EA0N6KI | backend.job.sync.external_api_error | External API call failed during sync | 
| EA0N6KJ | backend.job.sync.timeout | Operation timed out during sync | 
| EA0N6RL | backend.job.analytics.database_error | Database operation failed during analytics processing | 
| EA0N6RM | backend.job.analytics.external_api_error | External API call failed during analytics processing | 
| EA0N6RN | backend.job.analytics.timeout | Operation timed out during analytics processing | 
| EA19ATC | frontend.ui.unknown.unknown | Unknown UI error | 
| EA19B0H | frontend.ui.forms.validation_error | Form validation failed | 
| EA19B0I | frontend.ui.forms.submission_error | Form submission failed | 
| EA19B7L | frontend.ui.routing.not_found | Route not found | 
| EA19B7M | frontend.ui.routing.unauthorized | Route access unauthorized | 
| EA19NGG | frontend.state.unknown.unknown | Unknown state error | 
| EA19NNL | frontend.state.store.update_failed | State update operation failed | 
| EA19NNM | frontend.state.store.invalid_action | Invalid state action dispatched | 
| EA19NUP | frontend.state.persistence.storage_error | Local storage operation failed | 
| EA19NUQ | frontend.state.persistence.sync_error | State synchronization failed | 
| EA1A03K | frontend.api.unknown.unknown | Unknown API error | 
| EA1A0AP | frontend.api.request.network_error | Network request failed | 
| EA1A0AQ | frontend.api.request.timeout | Request timed out | 
| EA1A0AR | frontend.api.request.invalid_response | Invalid response received | 
| EA1A0HT | frontend.api.cache.cache_miss | Cache miss error | 
| EA1A0HU | frontend.api.cache.cache_invalid | Cache invalidation error | 
| EA9ZLDR | max.max_component.max_subcomponent.max_error | Maximum possible error code value | 


