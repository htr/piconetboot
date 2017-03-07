# piconetboot
a simple ipxe server

## building
```
$ gvt restore #fetch dependencies
$ make piconetboot
```

## usage
```
Usage of ./piconetboot:
  -addr string
    	ipxe reachable local address (default "127.0.0.1:8085")
  -data-dir string
    	directory containing boot client definitions
  -debug
    	increases logging verbosity level
  -static-dir string
    	directory containing static files
```

to run using the provided test setup:  
```
$ make download_images
$ ./piconetboot -addr 10.0.15.1:8085 -data-dir=$(pwd)/testing/datadir -static-dir=$(pwd)/testing/static -debug 
```

assuming the network interface `netboot0` exists and is configured with the subnet `10.0.15.0/x`:
```
$ qemu-system-x86_64 --curses -boot d -cdrom $(pwd)/testing/ipxe.iso -net nic,model=virtio,macaddr=00:00:00:11:11:12 -net bridge,br=netboot0 -uuid 4c4c4544-0032-3510-8034-cac04f343832 -smbios type=1,serial=A259481 -m 1024
```

## setting up a local testing environment


```
# /etc/systemd/network/10-netboot0.netdev
[NetDev]
Name=netboot0
Kind=bridge
```
```
# /etc/systemd/network/10-netboot0.network
[Match]
Name=netboot0

[Network]
Address=10.0.15.1/24
LinkLocalAddressing=yes
IPMasquerade=yes
DHCPServer=yes

[DHCPServer]
EmitDNS=yes
DNS=8.8.8.8
```

### building ipxe
checkout ipxe:

```
$ git clone git://git.ipxe.org/ipxe.git
```

create a script to be embedded into ipxe's images:

```
$ cat ipxe/script.ipxe
#!ipxe

dhcp
chain http://10.0.15.1:8085/
```

build ipxe:

```
cd ipxe/src
make EMBED=../script.ipxe -j8
```

relevant images files:

* src/bin/undionly.kpxe 
* src/bin/ipxe.iso 


The iso image can be used during development with a VM.
The PXE image can be used with dnsmasq or dhcpd:

```
#dhcpd.conf
...
...

group {
	filename "undionly.kpxe"
	host yada { hardware ethernet 0:1:2:3:4:5; }
}

...
...
```

