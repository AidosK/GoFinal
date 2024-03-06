## Booking and Reservation Project

This project is a simple hotel booking and reservation project with key features like;

- Authentication
- Create and manage rooms
- Showcase the rooms
- Check for rooms availability
- Allow booking of one room
- Have a backend for management
- Review existing bookings
- Show a calendar of bookings
- Change or cancel a booking
- Todo

### Live Website and login details

- username/email: aidosuniversity@gmail.com
- password: Password

### Packages used

- [Alex Edwards SCS](https://github.com/alexedwards/scs/v2) package for managing sessions
- [Chi router](https://github.com/go-chi/chi/v5)
- [Justinas nosurf](https://github.com/justinas/nosurf)
- [JackC PGX](https://github.com/jackc/pgx/v5) pgx is a pure Go driver and toolkit for PostgreSQL.
- [Go Simple Mail](https://github.com/xhit/go-simple-mail) Used for sending mails.
- [Simple DataTable](https://github.com/fiduswriter/Simple-DataTables) Used for tables.
- [Buffalo Soda](https://gobuffalo.io/pt/documentation/database/soda/) Used for tables.

### Note:

- Create your own go mod file and delete the one used here, run the following command `go mod init your-project-name`
- your-project-name is usually your github link and the name of your project, example "github.com/diaslmb/go-project". This is not a must, but a recommendation.
- Change the name of every import to your current go mod name. Example, open the main.go file, in the `required imports` section, replace these "github.com/diaslmb/go-project/pkg/config" to "github.com/diaslmb/your-project-name/pkg/config". Go through all files and make this replacement
- After all the necessary changes, run the app `go run cmd/web/*.go` this will install all the third party packages and run the server on the selected port.
- Create your postgres db
- Setup the flags in main.go file - `cmd/web/main.go`
- Setup the flags in run.sh file
- Setup the .env file.
- Setup the `database.yml`.

### Run the server

- Manual: `go run cmd/web/main.go cmd/web/middleware.go cmd/web/routes.go cmd/web/sendMail.go`
- Batch:  
**On Windows** - create a `run.bat` file in the root directory of the project and paste the below code

  ```
  go build -o bookings cmd/web/*.go
  ./bookings.exe
  ```

  Then run `run.bat` in the terminal
