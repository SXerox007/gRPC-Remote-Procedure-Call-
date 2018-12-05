# get the postges
go get -u github.com/lib/pq

export PATH=$PATH:/Applications/Postgres.app/Contents/Versions/latest/bin

sql -U postgres

# start db
\c name_of_db
# list of db
/list
