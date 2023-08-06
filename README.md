# gostore
Key-Value store Golang implementation with TTL.
Value type is a generic type (any).

Use TTL value of zero to disable invalidation.

The minimum viable TTL value is 1 second.
The recommended maximum size is 1 million values.

For better optimization, use the appropriate constructors:
- New - if key is an alphanumeric string
- NewUUID - if key is an UUID string
- NewNumeric - if key is a numeric string
