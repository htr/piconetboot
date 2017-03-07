piconetboot:
	go build -o piconetboot github.com/htr/piconetboot/cmd/piconetboot

clean:
	rm -f piconetboot

.PHONY: clean
