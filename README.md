# gouxchecker

## Description

Go package designed to assist developers and designers in evaluating the UX/UI of websites. Utilizing the powerful web scraping capabilities of Go's Colly package, `gouxchecker` focuses on essential elements such as font usage, image accessibility, correct spelling.

## Features

Currently provides the following features:

- Font Detection: Lists the fonts used on a webpage.
- Image Alt Text Verification: Checks for the presence of alt text in images to improve accessibility.
- Typo Detection: Scans the website content for any spelling errors.

## Prerequisites

- Uses Go 1.22

## Usage

To run gouxchecker, you need to set the target URL by modifying the url variable in main.go. For example, if your development server is at port 8080 in localhost:

`url := "http://localhost:8080/"`

You can also change `url` to point at a live website. After setting the url, run
`go run cmd/main.go`

## License

This project is licensed under the MIT License
