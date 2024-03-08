# Wifi Speed Watcher

A little program I designed to continuously monitor and record your WiFi speed over time and then present the data in a graph.

## TODO
- [ ] Add CLI params (Options to set the port for the dashboard and to set the interval that the wifi speed is recorded at)
- [ ] Maybe release builds at https://github.com/nekoify/wifi-speed-watcher/releases
- [ ] Make the datastore check if all files exist (how it currently checks: does the first file exist? Yes -> Then datastore must exist (if the other files don't exist, they wont be made))

## How to run
### Building from source
```
git clone https://github.com/nekoify/wifi-speed-watcher
cd wifi-speed-watcher
go get .
go build .
./main
```
#### Parameters (Not Implemented)
`-port` 
Default: 8080
Example: `-port 8000`
 
`-interval`
Default: 300 (5 Minutes)
Example: `-interval 150`
