piconetboot:
	go build -o piconetboot github.com/htr/piconetboot/cmd/piconetboot

download_images:
	mkdir testing/static || true
	wget -c -O testing/static/coreos_production_pxe.vmlinuz http://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe.vmlinuz
	wget -c -O testing/static/coreos_production_pxe_image.cpio.gz http://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe_image.cpio.gz
	wget -c -O testing/static/debian-live-8.7.1-amd64-standard.squashfs http://cdimage.debian.org/debian-cd/current-live/amd64/webboot/debian-live-8.7.1-amd64-standard.squashfs
	wget -c -O testing/static/debian-live-8.7.1-amd64-standard.vmlinuz http://cdimage.debian.org/debian-cd/current-live/amd64/webboot/debian-live-8.7.1-amd64-standard.vmlinuz
	wget -c -O testing/static/debian-live-8.7.1-amd64-standard.initrd.img http://cdimage.debian.org/debian-cd/current-live/amd64/webboot/debian-live-8.7.1-amd64-standard.initrd.img

clean_images:
	rm -f testing/static/coreos_production_pxe.vmlinuz testing/static/coreos_production_pxe_image.cpio.gz testing/static/debian-live-8.7.1-amd64-standard.squashfs testing/static/debian-live-8.7.1-amd64-standard.vmlinuz testing/static/debian-live-8.7.1-amd64-standard.initrd.img 

clean:
	rm -f piconetboot

.PHONY: clean download_images clean_images
