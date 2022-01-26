# react-env

A simple+secure (single file) tool to create dump react env vars to a file.

This is useful to dynamically configure PWA/SPA static asset sites (running in kubernetes).


## how does it work

in your index.html you'd have something such as:

```html
<script src="/.env.js" />
```

You'd then be able to load data/settings from `window.env.`

..for example:

![image](https://user-images.githubusercontent.com/462087/151183603-965d8c1d-e736-4155-97f8-d404f1221fec.png)



## usage

```
❯ ./react-env  --help
Usage of ./react-env:
      --dest string     Destination path to write to (relative or absolute) (default ".")
      --file string     name of the file to write (default ".env.js")
      --prefix string   Prefixes of env vars to look for (default "REACT_APP_")
pflag: help requested
```

> Sorry, more docs to follow + Docker examples


## attribution

Inspired by https://github.com/andrewmclagan/react-env
