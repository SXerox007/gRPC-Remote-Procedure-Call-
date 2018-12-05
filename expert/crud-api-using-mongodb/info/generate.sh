# go to mongodb folder 
cd mongodb-osx-X86_64-4.0.4
# run the db 
bin/mongodb
# Error happen not specify the folder of db
# first ll there is bin folder
mkdir data
mkdir data/db
# here is db folder
ls data
#For help
bin/mongod help
#set the db path
bin/mongod --dbpath data/db
# start db
bin/mongodb



# some error
# Failed to set up listener: SocketException: Address already in use
sudo service mongod stop
sudo mongod

# or
sudo killall -15 mongod
sudo mongod