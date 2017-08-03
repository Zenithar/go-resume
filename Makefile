all: clean build md pdf 

clean:
	@rm -f cv.* optimized.pdf 

build:
	@echo "> Compiling resume transformer"
	@go build

md: build
	@echo "> Building Markdown file"
	@./resume -yaml examples/fr.yaml -tmpl md.tmpl > README.md

pdf: build
	@echo "> Building XSL-FO file"
	@./resume -yaml examples/fr.yaml -tmpl fo.tmpl > cv.fo
	@echo "> Compiling to PDF"
	@fop cv.fo cv.pdf
	@echo "> Optimizing PDF"
	@gs -sDEVICE=pdfwrite -dCompatibilityLevel=1.4 -dPDFSETTINGS=/screen -dNOPAUSE -dQUIET -dBATCH -sOutputFile=optimized.pdf cv.pdf
	@echo "> Signing PDF"
	@gpg --clearsign --output=signed.pdf optimized.pdf

latex: build
	@echo "> Building Latex file"
	@./resume -yaml examples/fr.yaml -tmpl latex.tmpl > cv.tex
	@echo "> Compiling using XeTeX"
	@xetex cv.tex
