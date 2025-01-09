# Changelog

## [1.4.0](https://github.com/umeh-promise/blog/compare/v1.3.0...v1.4.0) (2025-01-09)


### Features

* add comment retrieval to UpdatePost and GetAllPost handlers ([96e539f](https://github.com/umeh-promise/blog/commit/96e539f59e9c3977b1942619af1b2cdb6bda5227))
* add CORS support and configure rate limiter in main application ([9b8b1df](https://github.com/umeh-promise/blog/commit/9b8b1dfcf8a556089cbe063533d77f7aeefac48b))
* add CORS support and rate limit exceeded response handling ([c12fc2b](https://github.com/umeh-promise/blog/commit/c12fc2bece7d194af218a3be35e46ec7efd93021))
* add database indexes for improved query performance ([99b5fd2](https://github.com/umeh-promise/blog/commit/99b5fd25d0708c4e070ab142e210a9e878b8f491))
* implement pagination and sorting for GetAllPost handler ([0831a5e](https://github.com/umeh-promise/blog/commit/0831a5e03cb7e94737bf250f07541c705a1af3b0))
* implement rate limiting middleware and configuration ([eb765cc](https://github.com/umeh-promise/blog/commit/eb765cc164604d0d22d308dcb9cbd59ac05bb2dc))
* retrieve comments for each post in GetAllPost handler ([93161af](https://github.com/umeh-promise/blog/commit/93161af0c9844c4746d037a4979c6376aa0e4937))

## [1.3.0](https://github.com/umeh-promise/blog/compare/v1.2.0...v1.3.0) (2025-01-09)


### Features

* add comment retrieval to post handler and update post model to include comments ([5be34e5](https://github.com/umeh-promise/blog/commit/5be34e527419d24b46cf460c263b748183f39cd0))
* enhance user registration and login payloads with email validation and password handling ([aab81ee](https://github.com/umeh-promise/blog/commit/aab81eeed4e27ab6a907395ccf55752fc1d04268))
* implement comment functionality with handler, service, repository, and model ([e200c77](https://github.com/umeh-promise/blog/commit/e200c774b218ecc86d96bdf3e6aeb33b4c933e2c))
* integrate comment handling into post routes and restructure main application setup ([975287d](https://github.com/umeh-promise/blog/commit/975287d1fada252bba1043b2776ad87ad1df9257))

## [1.2.0](https://github.com/umeh-promise/blog/compare/v1.1.0...v1.2.0) (2025-01-09)


### Features

* add constants for token expiration, issuer, and authentication secret ([ca1f478](https://github.com/umeh-promise/blog/commit/ca1f4784cb49fb596ab2f30e25f21528a723841d))
* add error handling for duplicate email and username, and unauthorized request logging ([b9ed2d7](https://github.com/umeh-promise/blog/commit/b9ed2d743d3faac47751545491493ba8c24017d8))
* add JWT utility functions for token generation and validation ([50019d7](https://github.com/umeh-promise/blog/commit/50019d7c02dae843f4b1f4224fa5f89fdbe364bf))
* add role management interface and repository implementation ([338d7b0](https://github.com/umeh-promise/blog/commit/338d7b0f2bd241bbc19f8dd1b52dc34f0dbe9883))
* add user roles management and versioning to users table ([2d41d98](https://github.com/umeh-promise/blog/commit/2d41d9875f60c9c3d7f8566a57ad6e08d1f10ea5))
* enhance user model with role object and add forbidden error handling ([9c1f5c8](https://github.com/umeh-promise/blog/commit/9c1f5c829f9c1d707ecdd522044185d7f41e39ee))
* implement role management middleware and service for post ownership checks ([0f68bc3](https://github.com/umeh-promise/blog/commit/0f68bc3cb468fa7045434fc0cc4e17e808ee79b5))
* implement user management functionality with CRUD operations ([03e36e7](https://github.com/umeh-promise/blog/commit/03e36e7a7109058491362f4920e3e2010d6e6714))
* implement user registration and login handlers with JWT authentication ([c91e0b0](https://github.com/umeh-promise/blog/commit/c91e0b0f282da985e8f49a9335b73eb0acd9bff6))
* refactor post handler and middleware for user context integration ([18f4541](https://github.com/umeh-promise/blog/commit/18f4541451f285298eebce0d1d3af296efac80b2))

## [1.1.0](https://github.com/umeh-promise/blog/compare/v1.0.0...v1.1.0) (2025-01-08)


### Features

* add post middleware for retrieving posts by ID and handling errors ([e3f31b3](https://github.com/umeh-promise/blog/commit/e3f31b3126d5c053cf1c6c0b3f3cd74f6929adb3))
* add post middleware to enhance post handling in the application ([023d59c](https://github.com/umeh-promise/blog/commit/023d59cd6086b9297f83eebda52f406c1d4b39aa))
* implement CRUD operations for posts; add logger and update Post model ([3c6ac5c](https://github.com/umeh-promise/blog/commit/3c6ac5c985ab73d46c72633e6104d7df67e61c63))
* implement GetAllPost method and update post routes to include retrieval of all posts ([f82cfcd](https://github.com/umeh-promise/blog/commit/f82cfcdeaef8f5a764bd7b362e7f968861895feb))
* refactor post handler methods to utilize middleware for post retrieval and add GetAllPost method ([ba3e1aa](https://github.com/umeh-promise/blog/commit/ba3e1aa59e59783e8f84a8070f3b162a2930eb86))

## 1.0.0 (2025-01-07)


### Bug Fixes

* correct GitHub Actions token reference and update Go module dependencies ([8e3cb82](https://github.com/umeh-promise/blog/commit/8e3cb82fdd64af4b4031c7e0971e98f3f1946d70))
* update Go version in GitHub Actions workflow to 1.23.4 ([d250bce](https://github.com/umeh-promise/blog/commit/d250bce428c40cc229f93066d78726e34d648e93))
* update go.sum to include new dependencies and remove outdated ones ([91b4db8](https://github.com/umeh-promise/blog/commit/91b4db8a3fed673ea602f358334e7105fa67fd3d))
