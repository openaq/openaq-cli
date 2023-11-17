---
title: Getting started
description: A guide in my new Starlight docs site.
sidebar:
  order: 1
---

## Installing OpenAQ CLI

> The OpenAQ CLI tool is still a *work in progress* and may be unstable until a 1.0 release. 

------
### Windows installation using Scoop
**What is Scoop?**  
[Scoop](https://scoop.sh/) is a command-line installer for Windows. It's designed to allow users to install software without the usual Windows GUI and without needing administrative permissions.

**Who should use Scoop?**  
Windows users who are comfortable with the command line and desire easy installation of software on their system.

**Steps to Install via Scoop**:

1. **Add scoop bucket**  
   Buckets in Scoop are similar to taps in Homebrew. By adding a bucket, you're telling Scoop where to find the software you want to install.

    ```bash
    scoop bucket add openaq-bucket https://github.com/openaq/scoop-bucket
    ```

2. **Install**  
   This command installs the OpenAQ CLI tool on your Windows machine.

    ```bash
    scoop install openaq-cli
    ```

---



### MacOS installation using HomeBrew

**What is Homebrew?**  
[Homebrew](https://brew.sh/) is a popular package manager for macOS that allows users to easily install software and tools directly from the command line. It's known for its simplicity and vast library of available packages. Homebrew has also been extended to work on Linux.

**Who should use Homebrew?**  
If you're on macOS or Linux and prefer a straightforward command-line installation method, Homebrew is an excellent choice. It handles dependencies, updates, and uninstallation seamlessly.

**Steps to Install via Homebrew**:

1. **Add tap**  
   This step adds the OpenAQ repository to Homebrew, making the tool available for installation.

    ```bash
    brew tap openaq/homebrew-tap
    ```

2. **Install**  
   This command installs the OpenAQ CLI tool.

    ```bash
    brew install openaq-cli
    ```

---


### Linux installation 
*Linux users can install using HomeBrew, or manually from source.*

------
### From source
Compiled executables for Windows, Mac, Linux are available for download in the releases page: 
https://github.com/openaq/openaq-cli/releases/

After downloading place the exectuable where you keep exectuables and update you system `$PATH` variable to make the executable discoverable by your shell.

---

Alternatively you can install with Golang > 1.18 with:

```
go install github.com/openaq/openaq-cli
```


## Setup

> Note: An OpenAQ API Key is required to use the OpenAQ CLI. Registering for an API Key is free and only requires a valid email address. Register at https://api.openaq.org/register


To set your API Key run: 

```sh
openaq settings set api-key my-super-secret-openaq-api-key-1234-5678
```

replacing `my-super-secret-openaq-api-key-1234-5678` with your API Key

