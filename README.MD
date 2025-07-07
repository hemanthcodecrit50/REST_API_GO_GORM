TO RUN THE COMPILE DAEMON : CompileDaemon -command=".\REST_API_WITH_GORM.exe"

INSTALLS NEEDED: 

go install github.com/githubnemo/CompileDaemon@latest
go get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u github.com/gorm.io/driver/postgres
./CompileDaemon -directory=yourproject/
CompileDaemon -command="./MyProgram -my-options"




go get github.com/jackc/pgx/v5