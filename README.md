# Go Parcel Server
Basic Go Server with a React frontend, built with parcel bundler. Includes HMR and parcel language support uses realize to watch and run go code.

## Setup and Running
```bash
go get github.com/tockins/realize

git clone https://github.com/ccutch/go-parcel my-project
cd my-project
npm i

realize start
```

# Layout
 - /frontend: All React code
 - /frontend-dist (ignored): Compiled static assets
 - /: Go files like any go project with root directory being main
