----------

<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">HTML Metadata Parser</h3>
  <p align="center">
    ⚡️ Pull data from web links, including title, description, photos, videos, and more [via OpenGraph]
    <br />
    <br />
    <a href="https://github.com/LinkviteApp/metadata-parser">View Demo</a>
    ·
    <a href="https://github.com/oLinkviteApp/metadata-parser/issues">Report Bug</a>
    ·
    <a href="https://github.com/LinkviteApp/metadata-parser/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About

This package lets you to use Facebook OpenGraph tags to extract information from an HTTP web address and retrieve meta data like title, description, photos, videos, and more. Inspired by [link-preview-js](https://github.com/ospfranco/link-preview-js), and built with [Go](https://go.dev/).



<!-- GETTING STARTED -->
## Getting Started

The app runs on a backend, so to get started, clone the repo, install dependencies and run the server to start making http requests.

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

This package is accessible via http requests and operates on a backend. to make requests, either use the query parameter in your browser or the body parameter in a tool like [Postman](https://www.postman.com/).

### Query Parameter
 `GET` `http://localhost:8080/api/parse?link=https://example.com`

### Body Parameter
 `GET` `http://localhost:8080/api/parse`
 
 ```js
{
	"link": "https://example.com"
}
```

### ✅ Success Response

```js
{
    "data":  {
	    "name":  "YouTube",
	    "title":  "Rick Astley - Never Gonna Give You Up (Official Music Video)",
	    "description":  "The official video for “Never Gonna Give You Up” by Rick Astley....",
	    "domain":  "youtube.com",
	    "image":  "https://i.ytimg.com/vi/dQw4w9WgXcQ/maxresdefault.jpg",
	    "url":  "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	    "type":  "video.other",
	    "keywords":  ["rick astley, Never Gonna Give You Up, gonna give you up lyrics"]
	},
    "timeInMs":  1962
}
```

### ❌ Error Response

```js
{
    "error":  "Please provide a valid url.",
    "data":  null,
}
```



<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Parse any website
- [x] Return custom reponse
- [ ] Retrieve favicons
- [ ] Retrieve multiple images & videos

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

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Feel free to reach out at [@tryLinkvite](https://twitter.com/tryLinkvite) or [@kayode0x](https://twitter.com/kayode0x)  on twitter.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

This project wouldn't be possible without the following packages.  We appreciate your efforts 🙏

* [goquery](github.com/PuerkitoBio/goquery)
* [gin](github.com/gin-gonic/gin)



<!-- MISC -->
## Misc

This README template was from [here](https://github.com/othneildrew/Best-README-Template).

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/linkviteApp/metadata-parser.svg?style=for-the-badge
[contributors-url]: https://github.com/LinkviteApp/metadata-parser/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/linkviteApp/metadata-parser.svg?style=for-the-badge
[forks-url]: https://github.com/LinkviteApp/metadata-parser/network/members
[stars-shield]: https://img.shields.io/github/stars/linkviteApp/metadata-parser.svg?style=for-the-badge
[stars-url]: https://github.com/LinkviteApp/metadata-parser/stargazers
[issues-shield]: https://img.shields.io/github/issues/linkviteApp/metadata-parser.svg?style=for-the-badge
[issues-url]: https://github.com/LinkviteApp/metadata-parser/issues
[license-shield]: https://img.shields.io/github/license/linkviteApp/metadata-parser.svg?style=for-the-badge
[license-url]: https://github.com/LinkviteApp/metadata-parser/blob/main/LICENSE.txt
