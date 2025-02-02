# gouxchecker

## Description

Go package designed to assist developers and designers in evaluating the UX/UI of websites. Utilizing the powerful web scraping capabilities of Go's Colly package, `gouxchecker` focuses on essential elements such as font usage, image accessibility and correct spelling.

## Features

Currently provides the following features:

- Font Detection: Lists the fonts used on a webpage. Ideally websites should have 3 or fewer separate fonts.
- Image Alt Text Verification: Checks for the presence of alt text in images to improve accessibility.
- Typo Detection: Scans the website content for any spelling errors.
- Broken Links Detection: Lists every internal link that doesn't return an http status OK response.
- Responsiveness: Identifies pages that may not be compliant with responsive web design.

## Prerequisites

- Uses Go 1.22

## Usage

To run gouxchecker, you need to set the target URL by modifying the url variable in main.go. For example, if your development server is at port 8080 in localhost:

`	urlPtr := flag.String("url", "http://localhost:8080", "Target URL to check")
`

You can also change `url` to point at a live website. After setting the url, just
`go run cmd/main.go`

Results are currently stored in text files at the root of the project.

Theres's also a dockerfile provided, and you can run the image setting the target url as an argument

`docker build -t gouxchecker .`

`docker run --rm gouxchecker -url=http://localhost:8080`

## License

This project is licensed under the MIT License
