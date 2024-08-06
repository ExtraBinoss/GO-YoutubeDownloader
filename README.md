# README

## About

This is the official Wails Svelte template.

## Prerequisites

You will need to have the Wails development environment installed. You can find instructions on how to do this [here](https://wails.io/docs/gettingstarted/installation).

You will also need GO 1.16 or later installed. [link here](https://golang.org/doc/install)

- install youtube-dl, ffmpeg
```bash
sudo apt-get install -y ffmpeg
sudo pip install --break-system-packages --upgrade youtube_dl
```

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
