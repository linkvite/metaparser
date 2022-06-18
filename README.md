<div  id="top"></div>

<!-- PROJECT INTRO -->
<br />
<div align="center">
<h1  align="center">Metadata Parser</h1>
</div>

<h4 align="center">⚡️ Extract data from web links, including title, description, photos, videos, and more [via OpenGraph].</h4>

<p align="center">
    <a href="https://github.com/golang/go">
    <img src="https://img.shields.io/badge/Go-v1.8-blue.svg"
         alt="Golang">
    <a href="https://github.com/LinkviteApp/metaparser/issues">
    <img src="https://img.shields.io/github/issues/LinkviteApp/metaparser.svg"
         alt="GitHub issues">
    <a href="https://github.com/LinkviteApp/metaparser/pulls">
    <img src="https://img.shields.io/github/issues-pr-raw/LinkviteApp/metaparser.svg?&logo=github&logoColor=white"
         alt="GitHub pull requests">
    <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg"
         alt="GitHub License">
    <a href="https://github.com/LinkviteApp/metaparser/issues/new?labels=enhancement">
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields"
        alt="GitHub Contributions">
</p>

<p  align="center">
    <a  href="#about">About</a> •
    <a  href="#installation">Installation</a> •
    <a  href="#usage">Usage</a> •
    <a  href="#roadmap">Roadmap</a> •
    <a  href="#contributing">Contributing</a> •
    <a  href="#credits">Credits</a> •
    <a  href="#support">Support</a> •
    <a  href="#license">License</a>
</p>



<!-- ABOUT THE PROJECT -->
## About

This package lets you to use Facebook OpenGraph tags to extract information from a website (url / link) and retrieve metadata like title, description, photos, videos, and more.



<!-- INSTALLATION -->
## Installation


```sh

go get -u github.com/LinkviteApp/metaparser

```



<!-- USAGE EXAMPLES -->
## Usage

After installing the package, create a go file and paste the example below to get started.

```go

package main

import (
    "fmt"

    parser "github.com/LinkviteApp/metaparser"
)

func main() {
    var link string

    // You can provide your own link.
    link = "https://www.twitter.com/tryLinkvite"

    // Optional: Or use GetLink(). This runs on the terminal.
    link = parser.GetLink()

    // Optional: Pass custom headers parameters.
    // See the table below for the available parameters.
    headers := make(map[string]string)
	headers["Accept-Language"] = "en-US"
	headers["User-Agent"] = "googlebot"

    // Pass the required URL and optional parameters.
    options := parser.Parameters{
		URL:           link,
		Timeout:       10,
		AllowRedirect: false,
		Headers:       headers,
	}

    data, err := parser.ParseLink(options)

    if err != nil {
        panic(err)
    }

    //optional: Convert the parsed data to JSON
    result := parser.ToJSON(data)
    fmt.Println(result)
}

```

### Options

Aside the required URL, you can pass optional parameters which should add more functionality in the parsing of the provided link.

| Property Name                                                                          |                                                                             Result                                                                              |
| -------------------------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------: |
| URL (**required**)                                           |                                 The URL to parse. (Must start with http or https)                                 |
| Headers (**optional**) (ex: { 'user-agent': 'googlebot', 'Accept-Language': 'en-US' }) |                                                                Add request headers to fetch call                                                                |
| Timeout (**optional**) (ex: 1000) (default 10)                                                     |                                                                 Timeout for the request to fail                                                                 |
| AllowRedirect (**optional**) (default false)                                         | For security reasons, the library does not automatically follow redirects, a malicious agent can exploit redirects to steal data, turn this on at your own risk |
        
In your root directory, run

```sh
go run .
```
        
Once the program is running, you'll get

```sh

👋 Enter the url of the web page 👇

================================================================

```

Next provide the link you want to parse

```sh

👋 Enter the url of the web page 👇

================================================================

https://github.com

```


### ✅ Successful Response

If you provided a valid url, you will get a response that looks like this:

```sh

✅ Valid URL provided.

================================================================

✅ Generated metadata template.

================================================================

⏳ Updating metadata from html document...

================================================================

✅ Updated metadata from html document.

================================================================

⏱ Total time taken: 494 milliseconds.

================================================================

📋 Metadata:

================================================================

```

```json
{
    "name": "GitHub",

    "title": "GitHub: Where the world builds software",

    "description": "GitHub is where over 83 million developers shape the future of software, together. Contribute to the open source community, manage your Git repositories, review code like a pro, track bugs and feat...",

    "domain": "github.com",

    "url": "https://github.com/",

    "type": "website",

    "images": ["https://github.githubassets.com/images/modules/site/social-cards/github-social.png"],

    "favicons": ["https://github.githubassets.com/favicons/favicon.svg"]
}
```


### ❌ Error Response

**PS** All links must be of scheme `http` or `https`. An error response would look like this:

```sh

👋 Enter the url of the web page 👇

================================================================

github.com

================================================================

❌ Failed to parse the url. Reason: The url must be of scheme http or https.

================================================================

```



<!-- ROADMAP -->
## Roadmap

- [x] Parse any website

- [x] Return custom reponse

- [x] Retrieve favicons

- [x] Retrieve multiple images
        
- [ ] Your awesome feature 😉

See the [open issues](https://github.com/LinkviteApp/metaparser/issues) for a full list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

Don't forget to give the project a star! Thanks again!

1. Fork the Project

2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)

3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)

4. Push to the Branch (`git push origin feature/AmazingFeature`)

5. Open a Pull Request



<!-- CREDITS -->
## Credits

- 💡 Inspired by [link-preview-js](https://github.com/ospfranco/link-preview-js)

- 🛠 Built on top of [goquery](https://github.com/PuerkitoBio/goquery)

- ⚡️ Written in [Go](https://github.com/golang/go)

- 📝 MD Template was from [here](https://github.com/othneildrew/Best-README-Template)



<!-- SUPPORT -->
## Support

Feel free to reach out on twitter [@tryLinkvite](https://twitter.com/tryLinkvite) or [@kayode0x](https://twitter.com/kayode0x).



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p  align="right">(<a  href="#top">back to top</a>)</p>
