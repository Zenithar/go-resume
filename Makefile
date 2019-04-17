all: clean build md pdf

clean:
	@rm -f cv.* optimized.pdf

build:
	@echo "> Compiling resume transformer"
	@go build

md.fr: build
	@echo "> Building Markdown file"
	@./resume -revision $(shell git rev-parse --short HEAD) -yaml examples/fr.yaml -tmpl md.fr.tmpl > README.md

md.en: build
	@echo "> Building Markdown file"
	@./resume -revision $(shell git rev-parse --short HEAD) -yaml examples/en.yaml -tmpl md.en.tmpl > README.md

pdf.fr: build
	@echo "> Building XSL-FO file"
	@./resume -revision $(shell git rev-parse --short HEAD) -yaml examples/fr.yaml -tmpl fo.fr.tmpl > fr.fo
	@echo "> Compiling to PDF"
	@fop fr.fo fr.pdf
	@echo "> Optimizing PDF"
	@pdfcpu optimize fr.pdf fr_full.pdf 

pdf.en: build
	@echo "> Building XSL-FO file"
	@./resume -revision $(shell git rev-parse --short HEAD) -yaml examples/en.yaml -tmpl fo.en.tmpl > en.fo
	@echo "> Compiling to PDF"
	@fop en.fo en.pdf
	@echo "> Optimizing PDF"
	@pdfcpu optimize en.pdf en_full.pdf 

latex: build
	@echo "> Building Latex file"
	@./resume -revision $(shell git rev-parse --short HEAD) -yaml examples/fr.yaml -tmpl latex.tmpl > cv.tex
	@echo "> Compiling using XeTeX"
	@xetex cv.tex
