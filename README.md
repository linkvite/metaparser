<div  id="top"></div>

<!-- PROJECT INTRO -->
<br />
<div  align="center">
<h1  align="center">Metadata Parser</h1>
</div>

<h4  align="center">⚡️ Extract data from web links, including title, description, photos, videos, and more [via OpenGraph].</h4>

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;

![Go](https://img.shields.io/badge/Go-v1.8-blue.svg)

![Dependencies](https://img.shields.io/badge/dependencies-up%20to%20date-brightgreen.svg)

[![GitHub Issues](https://img.shields.io/github/issues/LinkviteApp/metadata-parser.svg)](https://github.com/LinkviteApp/metadata-parser/issues)

[![Contributions welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)](https://github.com/LinkviteApp/metadata-parser/issues/new?labels=enhancement)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

<p  align="center">
    <a  href="#about">About</a> •
    <a  href="#getting-started">Getting Started</a> •
    <a  href="#usage">Usage</a> •
    <a  href="#roadmap">Roadmap</a> •
    <a  href="#contributing">Contributing</a> •
    <a  href="#credits">Credits</a> •
    <a  href="#support">Support</a> •
    <a  href="#license">License</a>
</p>

---

<!-- ABOUT THE PROJECT -->
## About

This package lets you to use Facebook OpenGraph tags to extract information from an HTTP web address and retrieve meta data like title, description, photos, videos, and more.



<!-- GETTING STARTED -->
## Getting Started

1. Clone the repo

```sh

$ git clone https://github.com/linkvite/metadata-parser.git

```

2. Install go mods

```sh

$ go mod

```

3. Run the server

```sh

$ go run .

```



<!-- USAGE EXAMPLES -->
## Usage

Once the program is running, you'll get:

```sh

👋 Enter the url of the web page 👇

================================================================

```

Next provide the link you want to parse.

```sh

👋 Enter the url of the web page 👇

================================================================

https://github.com

```


### ✅ Success Response

If you provided a valid url, you'd get a response that looks like this:

```sh

✅ Valid URL provided.

================================================================

✅ Generated meta data template.

================================================================

⏳ Updating meta data from html document...

================================================================

✅ Updated meta data from html document.

================================================================

⏱ Total time taken: 494 milliseconds.

================================================================

📋 Meta data:

================================================================

```

```json
{
  "name": "GitHub",

  "title": "GitHub: Where the world builds software",

  "description": "GitHub is where over 83 million developers shape the future of software, together. Contribute to the open source community, manage your Git repositories, review code like a pro, track bugs and feat...",

  "domain": "github.com",

  "url": "https://github.com/",

  "type": "object",

  "images": [
    "https://github.githubassets.com/images/modules/site/social-cards/github-social.png"
  ],

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

❌ Failed to parse the url. Reason: The url must be a http or https url.

================================================================

```



<!-- ROADMAP -->
## Roadmap

- [x] Parse any website

- [x] Return custom reponse

- [x] Retrieve favicons

- [x] Retrieve multiple images

See the [open issues](https://github.com/LinkviteApp/metadata-parser/issues) for a full list of proposed features (and known issues).



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

- 🛠 Built on top of [goquery](github.com/PuerkitoBio/goquery)

- ⚡️ Written in [Go](https://go.dev/).

- 📝 MD Template was from [here](https://github.com/othneildrew/Best-README-Template)



<!-- SUPPORT -->
## Support

Feel free to reach out on twitter [@tryLinkvite](https://twitter.com/tryLinkvite) or [@kayode0x](https://twitter.com/kayode0x).



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p  align="right">(<a  href="#top">back to top</a>)</p>
