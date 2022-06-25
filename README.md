# amandagrams
Go webserver (gin/gonic) for [Amandagrams](https://amandagrams.surge.sh/).  

Deployed at: [https://amandagrams-tgoh98.koyeb.app/](https://amandagrams-tgoh98.koyeb.app/)  
Frontend: [https://github.com/tGoh98/amandagrams-react](https://github.com/tGoh98/amandagrams-react)  

### To run locally
#### Docker
1. `docker build . -t [tag_name]`
2. `docker run -p 8000:8000 --rm [tag_name]`  
3. Should be running at (http://localhost:8000)

#### Manually
Requires that you have [go](https://go.dev/doc/install) installed
1. `go mod download`
2. `go build .`
3. `./amandagrams -g` <-- need to generate word data first
4. `./amandagrams`
5. Should be running at (http://localhost:8000)
