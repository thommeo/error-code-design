# Error Codes Documentation

This document is auto-generated. Do not edit manually.

## Error Code Format

All error codes follow the format: E<type><data> where:
- E: Fixed prefix identifying this as an error code
- type: Single base-36 character (0-9,A-Z) identifying the error code format
- data: Variable-length base-36 encoded data specific to each format

Base-36 encoding uses digits 0-9 and letters A-Z to pack more information into fewer characters while remaining human-readable.


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
| E20MTQ8 | backend.handler.unknown.unknown | Unknown handler error | 
| E20MTXD | backend.handler.users.validation_error | Input validation failed for user operation | 
| E20MTXE | backend.handler.users.authorization_error | User lacks required permissions for operation | 
| E20MU4H | backend.handler.records.validation_error | Input validation failed for record operation | 
| E20MU4I | backend.handler.records.authorization_error | User lacks required permissions for record operation | 
| E20MUBL | backend.handler.analytics.validation_error | Input validation failed for analytics operation | 
| E20MUBM | backend.handler.analytics.authorization_error | User lacks required permissions for analytics operation | 
| E20N6DC | backend.job.unknown.unknown | Unknown job error | 
| E20N6KH | backend.job.sync.database_error | Database operation failed during sync | 
| E20N6KI | backend.job.sync.external_api_error | External API call failed during sync | 
| E20N6KJ | backend.job.sync.timeout | Operation timed out during sync | 
| E20N6RL | backend.job.analytics.database_error | Database operation failed during analytics processing | 
| E20N6RM | backend.job.analytics.external_api_error | External API call failed during analytics processing | 
| E20N6RN | backend.job.analytics.timeout | Operation timed out during analytics processing | 
| E219ATC | frontend.ui.unknown.unknown | Unknown UI error | 
| E219B0H | frontend.ui.forms.validation_error | Form validation failed | 
| E219B0I | frontend.ui.forms.submission_error | Form submission failed | 
| E219B7L | frontend.ui.routing.not_found | Route not found | 
| E219B7M | frontend.ui.routing.unauthorized | Route access unauthorized | 
| E219NGG | frontend.state.unknown.unknown | Unknown state error | 
| E219NNL | frontend.state.store.update_failed | State update operation failed | 
| E219NNM | frontend.state.store.invalid_action | Invalid state action dispatched | 
| E219NUP | frontend.state.persistence.storage_error | Local storage operation failed | 
| E219NUQ | frontend.state.persistence.sync_error | State synchronization failed | 
| E21A03K | frontend.api.unknown.unknown | Unknown API error | 
| E21A0AP | frontend.api.request.network_error | Network request failed | 
| E21A0AQ | frontend.api.request.timeout | Request timed out | 
| E21A0AR | frontend.api.request.invalid_response | Invalid response received | 
| E21A0HT | frontend.api.cache.cache_miss | Cache miss error | 
| E21A0HU | frontend.api.cache.cache_invalid | Cache invalidation error | 


