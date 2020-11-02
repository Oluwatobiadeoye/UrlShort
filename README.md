# UrlShort
A simple url shortener app that redirects to appropriate url based on http path

## How to use
* Add path to urls config in a yaml file. See [example](paths.yaml)
* Start app using ```go run main/main.go ``` and specify the file using ```-file "your_file_name_here"```
* For available configs, run ```go run main/main.go -h```