# Wifi Speed Watcher

A little program I designed to continuously monitor and record your WiFi speed over time and then present the data in a graph.

## TODO
- [ ] Add CLI params (Options to set the port for the dashboard and to set the interval that the wifi speed is recorded at)
- [ ] Maybe release builds at https://github.com/nekoify/wifi-speed-watcher/releases
- [ ] Fix the datastore so it will check that all files exist instead of only one before creating files

## How to run
### Building from source
```
git clone https://github.com/nekoify/wifi-speed-watcher
cd wifi-speed-watcher
go get .
go build .
./main
```
OR
```
git clone https://github.com/nekoify/wifi-speed-watcher
cd wifi-speed-watcher
go get .
go run .
```
#### Parameters (Not Implemented)
`-port` 
Default: 8080
Example: `-port 8000`
 
`-interval`
Default: 300 (seconds)
 Example: `-interval 150`
