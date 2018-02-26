# hestia

[![Build status](https://badge.buildkite.com/77c1b66942fce33485da9956acfa41fae91bbe889da4581783.svg)](https://buildkite.com/codeclimate/hestia)
[![Maintainability](https://api.codeclimate.com/v1/badges/8a284c45ce0874b1c61e/maintainability)](https://codeclimate.com/github/codeclimate/hestia/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/8a284c45ce0874b1c61e/test_coverage)](https://codeclimate.com/github/codeclimate/hestia/test_coverage)

> Hestia was the Greek goddess of the hearth and the home. She was the sister of
> Zeus and was often associated with Hermes, the two representing domestic life
> on the one hand, and business and outdoor life on the other.

Also, a central presence within the Code Climate organization whose mission is
to improve culture, one function at a time.

## overview

This project is structured as a golang project packaged as lambda functions.

There are currently two lambda functions defined:

- [api](cmd/api/api.go) (receives slack mentions)
- [handler](cmd/handler/handler.go) (processes command requests)

The handler function invokes implemented commands, found within the
[internal/commands](internal/commands) package.

Secrets are stored and retrieved via SSM Parameter Store. The lambda functions
are deployed via terraform in CI.

Implemented commands:

- [whoami](internal/commands/whoami.go): responds with user slack information
- [echo](internal/commands/echo.go): respond with supplied arguments
- [weather](internal/commands/weather.go): respond with local weather or weather from provided zip code
- [nowplaying](internal/commands/nowplaying.go): respond with currently playing tracks from list of last.fm usernames

You can run commands locally by building and running the cli binary.
