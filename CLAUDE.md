# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A Go CLI tool that enumerates prime number dates (dates in YYYYMMDD format that are prime numbers) for a given year.

## Build and Run Commands

```bash
# Run directly
go run main.go [year]

# Build binary
go build -o prime_days main.go

# Run tests (no tests currently exist)
go test ./...
```

## Architecture

Single-file Go application (`main.go`) with:

- **`isPrime(n int)`**: Trial division primality test up to √n
- **`getPrimeDates(start, end time.Time)`**: Iterates through date range, spawns goroutines for concurrent primality checks using channels
- **`main()`**: Parses optional year argument (defaults to current year), outputs prime dates

The program uses goroutines and channels to parallelize primality checks across all dates in the year.
