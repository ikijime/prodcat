# entclean:
# 	entclean User Group
# if you want good hot reload use 3 different cmd with tailwatch + temlwatch + air . with default config
entgenb:
	go generate ./ent

entgen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery ./ent/schema

templ:
	templ generate

templwatch:
	templ generate -watch

run:
	- tail
	- templ generate .\views 
	- air .

build:
	- tailmin
	- templ generate .\views 
	- go build -o bin/app

tail:
	tailwindcss-windows-x64.exe --config .\config\tailwind.config.js -i .\config\input.css -o .\static\css\output.css

tailwatch:
	tailwindcss-windows-x64.exe --config .\config\tailwind.config.js -i .\config\input.css -o .\static\css\output.css --watch

# Compile and minify your CSS for production
tailmin:
	tailwindcss-windows-x64.exe -i .\static\css\input.css -o .\static\css\output.css --minify
