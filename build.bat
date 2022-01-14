(
cd config
go build
cd ..
cd portforwarder
go build
cd ..
cd kademlia
cd networking
go build
cd ..
go build
cd ..
go install github.com/kk222mo/godist
) > buildlog.txt
