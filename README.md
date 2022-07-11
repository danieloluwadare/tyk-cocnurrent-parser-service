# tyk-cocnurrent-parser-service

# Getting Started
1. [x] Clone the project, go to the project root folder
2. [x] Run <code> go get ./... </code> to install all dependencies
3. [x] Setup Mysql server (using docker preferably)
4. [x] Create a new file in the project root. The file name is ".env" (Note: Just ".env" no name), then copy the default
   parameters in the "example.env" file into the ."env"
5. [x] Update all Properties with the DB prefix in .env file with the right connection details
6. [x] Open mysql workbench, open the default connection and create a new schema.
   (Note. The name you give it should match the one in the ".env"fle)
7. [x] Update TXT.PATH to the location of data.txt e.g (TXT.PATH=client/resources/data.txt)
8. [x] Update LOG.NAME to the desired name of your log e.g(LOG.NAME=tyk-parser-log.log)
9. [x] Start the GRPC SERVER with this command <code>  cd server && go run . </code>
10. [x] Start the GRPC CLIENT with this command <code>  go run client/main.go </code>