<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
  <h3 align="center">Captcha solver</h3>

  <p align="center">
     Takes in a url with alphanumberic captcha  
  </p>
</p>
<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#built-with">Built with</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>

  </ol>
</details>



### Built With

* [Gorilla mux](https://github.com/gorilla/mux)
* [Gocv](https://github.com/hybridgroup/gocv)
* [Gosseract](https://github.com/otiai10/gosseract)



<!-- GETTING STARTED -->
## Getting Started

Dependencies are handled by docker. port 5000 is used by default as configured in docker-compose.yml

`docker-compose build` 

`docker-compose up`

### Prerequisites

* Install docker compose
https://docs.docker.com/compose/install/
* If on linux server or non root user, this might be needed 
https://docs.docker.com/engine/install/linux-postinstall/

<!-- USAGE EXAMPLES -->
## Usage

`curl --header "Content-Type: application/json" \
   --request POST \
   --data '{"Url":"https://i.imgur.com/gxmRseR.jpg"}' \
   http://localhost:5000/captcha`
