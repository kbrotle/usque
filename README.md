# usque

🥚➡️🍏🍎

Usque is an open-source reimplementation of the Cloudflare WARP client's MASQUE mode. It leverages the [Connect-IP (RFC 9484)](https://datatracker.ietf.org/doc/rfc9484/) protocol and comes with many operation modes including a native tunnel, a SOCKS5 proxy, and a HTTP proxy.

## Table of Contents

- [usque](#usque)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Building from source](#building-from-source)
    - [Docker](#docker)
  - [Usage](#usage)
    - [Registration](#registration)
    - [Enrolling](#enrolling)
    - [Native Tunnel Mode (for Advanced Users, Linux and Windows only!)](#native-tunnel-mode-for-advanced-users-linux-and-windows-only)
      - [On Linux](#on-linux)
      - [On Windows](#on-windows)
      - [Routes on Linux](#routes-on-linux)
      - [Routes on Windows](#routes-on-windows)
    - [SOCKS5 Proxy Mode (easy, cross-platform)](#socks5-proxy-mode-easy-cross-platform)
    - [HTTP Proxy Mode (easy, cross-platform)](#http-proxy-mode-easy-cross-platform)
    - [Port Forwarding Mode (for Advanced Users, cross-platform)](#port-forwarding-mode-for-advanced-users-cross-platform)
    - [Connect/Disconnect Hooks](#connectdisconnect-hooks)
      - [Example on Linux](#example-on-linux)
      - [Example on Windows](#example-on-windows)
    - [TCP and HTTP/2 Support](#tcp-and-http2-support)
      - [HTTP/2 Configuration](#http2-configuration)
    - [Configuration](#configuration)
      - [Fields](#fields)
  - [ZeroTrust support](#zerotrust-support)
  - [Performance](#performance)
    - [Performance Tuning](#performance-tuning)
      - [Linux/BSD](#linuxbsd)
      - [DNS](#dns)
  - [Using this tool as a library](#using-this-tool-as-a-library)
  - [Known Issues](#known-issues)
  - [Miscellaneous](#miscellaneous)
    - [Censorship circumvention](#censorship-circumvention)
  - [Should I replace WireGuard with this?](#should-i-replace-wireguard-with-this)
    - [Why would you still switch?](#why-would-you-still-switch)
  - [Protocol \& research details](#protocol--research-details)
  - [Why was this built?](#why-was-this-built)
  - [Why did you fork connect-ip-go?](#why-did-you-fork-connect-ip-go)
  - [Why the name?](#why-the-name)
  - [Contributing](#contributing)
  - [Acknowledgements](#acknowledgements)
  - [Disclaimer](#disclaimer)

> **Personal fork note:** I'm using this primarily for learning how MASQUE/Connect-IP works. The original repo is at [Diniboy1123/usque](https://github.com/Diniboy1123/usque). Also fixed a small typo in the RFC link text above (`Connnect-IP` → `Connect-IP`).
>
> **My notes so far:**
> - SOCKS5 mode is the easiest way to get started — just point your browser's proxy settings at `127.0.0.1:1080`.
> - Tested on Linux `amd64` only (same as upstream). Haven't tried Windows yet.
> - If you're just experimenting, HTTP proxy mode on port `8080` is also handy for quick `curl` tests: `curl --proxy http://127.0.0.1:8080 https://cloudflare.com/cdn-cgi/trace`
> - **2024-07-15:** Noticed the Table of Contents had a broken anchor for `ZeroTrust support` in some renderers — the section heading uses a capital T in "Trust" so the link works fine on GitHub but worth noting.

## Installation
