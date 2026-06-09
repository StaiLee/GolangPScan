# 🔍 GolangPScan

[![CI](https://github.com/StaiLee/GolangPScan/actions/workflows/ci.yml/badge.svg)](https://github.com/StaiLee/GolangPScan/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/StaiLee/GolangPScan?color=39E0D8&logo=github)](https://github.com/StaiLee/GolangPScan/releases/latest)
[![Go](https://img.shields.io/github/go-mod/go-version/StaiLee/GolangPScan?logo=go&color=00ADD8)](go.mod)
[![License: MIT](https://img.shields.io/github/license/StaiLee/GolangPScan?color=blue)](LICENSE)
[![Stars](https://img.shields.io/github/stars/StaiLee/GolangPScan?logo=github&color=e3b341)](https://github.com/StaiLee/GolangPScan/stargazers)

> A fast, concurrent **TCP connect** port scanner written in pure Go.

---

## 📖 Overview

**GolangPScan** is a lightweight port scanner built around Go's concurrency
primitives. It spins up a pool of goroutine **workers** that pull ports from a
channel, attempt a TCP `Dial` against the target, and report which ports accept
a connection. Results are collected, sorted, and printed.

It's intentionally small and dependency-free (standard library only) — a clean
reference implementation of the **worker-pool pattern** applied to network
scanning.

## ✨ Features

- ⚡ **Concurrent** — 100 goroutine workers by default
- 🎯 **Flexible port parsing** — single ports, comma lists and ranges
  (e.g. `22,80,443` or `1-1024`)
- 🧹 **Zero dependencies** — pure Go standard library
- 📋 **Sorted output** of open ports

## 🚀 Usage

### Option 1 — Download a binary
Grab the prebuilt binary for your OS from the
[latest release](https://github.com/StaiLee/GolangPScan/releases/latest)
(Linux, Windows, macOS — amd64 & arm64).

### Option 2 — Run from source
```bash
git clone https://github.com/StaiLee/GolangPScan.git
cd GolangPScan
go run main.go
```

## ⚙️ Configuration

> **Note:** the target host and port range are currently set as constants at the
> top of `main.go` — edit them before running:

```go
address  := fmt.Sprintf("test.com:%d", p) // <- target host
portsToScan := "1-1024"                    // <- ports (e.g. "22,80,443" or "1-65535")
numWorkers  := 100                         // <- concurrency
```

## 🧠 How it works

```
ports channel  ──▶  [ worker 1 ]
               ──▶  [ worker 2 ]  ──▶  results channel  ──▶  sorted open ports
               ──▶  [  ...100  ]
```

Each worker performs a `net.Dial("tcp", host:port)`. A successful dial means the
port is open; the worker pushes it to the results channel. A `sync.WaitGroup`
tracks completion.

## 🗺️ Roadmap

- [ ] CLI flags (`-host`, `-ports`, `-workers`) instead of source constants
- [ ] Connection timeout / rate limiting
- [ ] Service/banner grabbing
- [ ] JSON output

## ⚠️ Disclaimer

For **educational use and authorized testing only**. Scanning hosts you do not
own or have explicit permission to test may be illegal. Use responsibly.

## 📄 License

MIT — see [LICENSE](LICENSE).
