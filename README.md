# heat
heat is a gRPC server that wraps Torsten Traenker's heaterControl.exp (http://torsten-traenkner.de/wissen/smarthome/heizung.php).

# heatctl
heatctl is a gRPC client to control the temperature of heating thermostats.

## Usage
heatctl -r livingroom -t 24

Currently supported devices: Sygonix HT100BT

Many thanks to Torsten Traenkner for providing his code.
