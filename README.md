# Wifi Speed Watcher

A little program I designed to continuously monitor and record your WiFi speed over time and then present the data in a graph.
 
Works on Linux and Windows, no clue if it works on macOS but it should do.

## TODO
- [x] Add CLI args (Options to set the port for the dashboard and to set the interval that the wifi speed is recorded at)
- [x] Maybe release builds at https://github.com/nekoify/wifi-speed-watcher/releases
- [ ] Fix the datastore so it will check that all files exist instead of only one before creating files
- [ ] Docker Support


## How to run
### Building from source
```
git clone https://github.com/nekoify/wifi-speed-watcher
cd wifi-speed-watcher
go get .
go build .
./main [args]
```
#### Arguments
`-port` - Sets the port that the dashboard runs on
 
 Default: 8080
  
 Example: `-port 8000`
 
  
`-interval` - Sets the interval that the speed is recorded at in seconds. Note: the average time a speed test takes is ~20 seconds and the time taken is taken off the interval.
  
 Default: 300
  
 Example: `-interval 150`
 

## Screenshot

![Screenshot](https://raw.githubusercontent.com/nekoify/wifi-speed-watcher/main/assets/screenshot1.png)


