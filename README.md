# Wifi Speed Watcher

A little program I designed to continuously monitor and record your WiFi speed over time and then present the data in a graph.
 
Checked on Linux, no clue if it works on Windows or macOS but will be tested later on.

## TODO
- [ ] Add CLI params (Options to set the port for the dashboard and to set the interval that the wifi speed is recorded at)
- [ ] Maybe release builds at https://github.com/nekoify/wifi-speed-watcher/releases
- [ ] Fix the datastore so it will check that all files exist instead of only one before creating files
- [ ] If the set interval is lower then the time to finish a speed test, wait till speed test finishes

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
`-port` - Sets the port that the dashboard runs on
 
 Default: 8080
  
 Example: `-port 8000`
  
`-interval` - Sets the interval that the speed is recorded at in seconds\
  
 Default: 300
  
 Example: `-interval 150`
 
 `-clean` - Wipes all previous data
  
 Default: N/A
  
 Example: `-clean`

## Screenshot

![Screenshot](https://raw.githubusercontent.com/nekoify/wifi-speed-watcher/main/assets/screenshot1.png)


