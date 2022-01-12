(

cd portforwarder
go build
cd ..
cd kademlia
go build
cd ..
go install github.com/kk222mo/godist
) > buildlog.txt
