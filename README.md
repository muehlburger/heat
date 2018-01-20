# heat
heat is a gRPC server written in go that wraps Torsten Traenker's heaterControl.exp (http://torsten-traenkner.de/wissen/smarthome/heizung.php). 

The heat binary is compiled to run on a Raspberry Pi per default (see Makefile).

# heatctl
heatctl is a gRPC client to control the temperature of heating thermostats.

## Usage
heatctl -r livingroom -t 24

Currently supported devices: Sygonix HT100BT

Many thanks to Torsten Traenkner for providing his code.
