# Error Codes Documentation

This document is auto-generated. Do not edit manually.


## Simple Format

Each error code is composed of two bytes encoded as follows:
- Class (8 bits): Identifies the error class (allows up to 256 distinct classes)
- ErrorType (8 bits): Identifies the specific error (allows up to 256 errors per class)

The format provides:
- Up to 256 different classes
- Up to 256 different error types per class
- Total of 65,536 possible unique error codes

Byte layout:
```
[CCCCCCCC][EEEEEEEE]
C: Class bits
E: ErrorType bits
```

| Code | Class.Type | Description | 
|----|----|----|
| E010000 | unknown.unknown | Unknown API error | 
| E010100 | api.unknown | Unknown API error | 
| E010101 | api.validation_error | API validation error | 
| E010102 | api.authorization_error | API authorization error | 
| E010200 | jobs.unknown | Unknown job error | 
| E010201 | jobs.database_query | Database query error in job | 
| E010202 | jobs.timeout | Job execution timeout | 



## App Component Format

Each error code is composed of three bytes encoded as follows:
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

Byte layout:
```
[AAAACCCC][CCSSSSSS][EEEEEEEE]
A: App bits
C: Component bits
S: SubComponent bits
E: ErrorType bits
```

| Code | App.Component.SubComponent.Type | Description | 
|----|----|----|
| E02104000 | backend.handler.unknown.unknown | Unknown handler error | 
| E02104101 | backend.handler.users.validation_error | Input validation failed for user operation | 
| E02104102 | backend.handler.users.authorization_error | User lacks required permissions for operation | 
| E02104201 | backend.handler.records.validation_error | Input validation failed for record operation | 
| E02104202 | backend.handler.records.authorization_error | User lacks required permissions for record operation | 
| E02104301 | backend.handler.analytics.validation_error | Input validation failed for analytics operation | 
| E02104302 | backend.handler.analytics.authorization_error | User lacks required permissions for analytics operation | 
| E02108000 | backend.job.unknown.unknown | Unknown job error | 
| E02108101 | backend.job.sync.database_error | Database operation failed during sync | 
| E02108102 | backend.job.sync.external_api_error | External API call failed during sync | 
| E02108103 | backend.job.sync.timeout | Operation timed out during sync | 
| E02108201 | backend.job.analytics.database_error | Database operation failed during analytics processing | 
| E02108202 | backend.job.analytics.external_api_error | External API call failed during analytics processing | 
| E02108203 | backend.job.analytics.timeout | Operation timed out during analytics processing | 
| E02204000 | frontend.ui.unknown.unknown | Unknown UI error | 
| E02204101 | frontend.ui.forms.validation_error | Form validation failed | 
| E02204102 | frontend.ui.forms.submission_error | Form submission failed | 
| E02204201 | frontend.ui.routing.not_found | Route not found | 
| E02204202 | frontend.ui.routing.unauthorized | Route access unauthorized | 
| E02208000 | frontend.state.unknown.unknown | Unknown state error | 
| E02208101 | frontend.state.store.update_failed | State update operation failed | 
| E02208102 | frontend.state.store.invalid_action | Invalid state action dispatched | 
| E02208201 | frontend.state.persistence.storage_error | Local storage operation failed | 
| E02208202 | frontend.state.persistence.sync_error | State synchronization failed | 
| E0220C000 | frontend.api.unknown.unknown | Unknown API error | 
| E0220C101 | frontend.api.request.network_error | Network request failed | 
| E0220C102 | frontend.api.request.timeout | Request timed out | 
| E0220C103 | frontend.api.request.invalid_response | Invalid response received | 
| E0220C201 | frontend.api.cache.cache_miss | Cache miss error | 
| E0220C202 | frontend.api.cache.cache_invalid | Cache invalidation error | 


