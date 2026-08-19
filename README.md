# Pokédex CLI

A command-line Pokédex written in Go as part of the Boot.dev backend curriculum.

The application uses the [PokeAPI](https://pokeapi.co/) to explore Pokémon locations, catch Pokémon, and inspect caught Pokémon.

## Features

- Interactive REPL
- Browse paginated Pokémon location areas
- Explore an area to see Pokémon found there
- Catch Pokémon with a probability based on their base experience
- Store caught Pokémon in an in-memory Pokédex
- Inspect a caught Pokémon:
  - height
  - weight
  - stats
  - types
- Cache PokeAPI responses to avoid unnecessary HTTP requests
- Automatic cache expiration
- Thread-safe cache using a mutex

## Commands

- `help` — display available commands
- `map` — display the next 20 location areas
- `mapb` — display the previous 20 location areas
- `explore <area>` — list Pokémon found in a location area
- `catch <pokemon>` — attempt to catch a Pokémon
- `inspect <pokemon>` — display some information about a caught Pokémon
- `pokedex` — list all caught Pokémon
- `exit` — close the application

## Run

```bash
go run .
```

## Tests

```bash
go test ./...
```

## Project structure

- internal/pokeapi — PokeAPI client and API response handling
- internal/pokecache — thread-safe response cache with automatic expiration
- REPL commands and application state are handled in the main package

## Notes

The Pokédex is stored and is reset when the application exits.