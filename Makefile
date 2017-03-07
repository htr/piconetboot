piconetboot:
	go build -o piconetboot github.com/htr/piconetboot/cmd/piconetboot

download_images:

clean:
	rm -f piconetboot

.PHONY: clean download_images
