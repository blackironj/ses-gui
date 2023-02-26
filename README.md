# SES-gui

<p align="left">
  <a href="https://github.com/blackironj/ses-gui/actions"><img alt="GitHub Actions status" src="https://github.com/actions/setup-go/workflows/build-test/badge.svg"></a>
</p>

This is a GUI tool for uploading, downloading, deleting AWS-SES Email template easily. And also you can test sending email with uploaded template using this tool 

> Currently, AWS-SES does not provide GUI editor yet. So you can use [AWS-CLI](https://awscli.amazonaws.com/v2/documentation/api/latest/index.html) tool to manage email-template instead of this tool. 

## Demo

![ex_img](./img/example-mainview.jpg)

## Build

> I do not test it on linux and macOS yet.

### Prerequisites

- You can find detailed prerequisites for building a `fyne` application on the website shown below.
- https://developer.fyne.io/started/

### Windows

```bash
go build -o ses-gui.exe -ldflags="-H windowsgui"
```
> If you want to check more detail logs on terminal, remove `-ldflags="-H windowsgui" flag in commandline

## Feature

- [x] Login with an aws access key
- [x] Upload an email template
- [x] Download an email template
- [x] Show list of templates
- [x] Delete a template
- [x] Send an email with selected template

### TODO

- [x] enhance UI / UX
- Testing
  - [ ] on Linux
  - [ ] on MacOS