curl -iX POST "http://localhost:30000/api/v0/group" -H  "accept: application/json" -H  "Content-Type: application/json" -d '[{"name":"UI","snaps":["ui-snap1","ui-snap2"]}, {"name":"Daemon","snaps":["daemon-snap1"]}]'
sleep 1
curl -iX POST "http://localhost:30000/api/v0/group" -H  "accept: application/json" -H  "Content-Type: application/json" -d '[{"limits":{"maxMemory":512000000},"name":"application","subGroups":["UI","Daemon"]},{"limits":{"maxMemory":256000000},"name":"Special-snap","snaps":["special-snap"]}]'
sleep 1
curl -iX GET "http://localhost:30000/api/v0/group/limits" -H  "accept: application/json" -H  "Content-Type: application/json" -d '["application","Special-snap"]'
