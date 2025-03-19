# THG Stack

Thg is a lighweight template for spinning up websites using <b>T</b>ailwind CSS, <b>H</b>TML, <b>G</b>olang, and [Chi](https://github.com/go-chi/chi) for routing.

## Features
- Live reloading on save using [Air](https://github.com/air-verse/air)
- Templates through native Golang text/template package 
- Reusable components

## Getting Started
Clone repo
```bash
git clone https://github.com/pisgahi/thg.git
cd thg
```

<b>Prerequisites</b>  
- [Install Tailwind CSS](https://tailwindcss.com/blog/standalone-cli) (preferably the standalone CLI).  
- Create a `.env` file and add the following. `PORT=8080` is just an example port:  
```bash
PORT=8080
```

<h4>Dev</h4>

1. Begin Tailwind file watching and generate the static CSS file 
```bash
make wind
```

2. In another terminal session start Air for live reloading
```bash
make air
```

<h4>Production</h4>

1. Run the final product
```bash
make run
```

2. Build the final executable
```bash
make build
```
