Remove-Item -Path ./output -Recurse -Force
mkdir output
go build -o ./output/ias-simulator.exe ./cmd/ias-simulator/main.go