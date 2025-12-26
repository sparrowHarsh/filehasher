# Concurrent File Hasher using Go Channels and Semaphores

## Problem Description

You are given a directory path. Your task is to recursively scan all files inside the directory and compute a cryptographic hash (e.g., SHA-256) for each file.

### Requirements

The system must:

- Traverse directories safely
- Read file contents
- Hash files concurrently
- Control concurrency using semaphores
- Coordinate workers using channels
- Handle large directories without exhausting memory or CPU

## Why this problem?

This problem simulates real backend and systems workloads:

- File indexing
- Antivirus scanners
- Backup systems
- Content-addressable storage (Git, Docker layers)
- Data deduplication systems

## Learning Objectives

It forces you to learn:

- Goroutines lifecycle
- Channels as pipelines
- Semaphores for resource limits
- Worker pools
- Graceful shutdown
- Error propagation