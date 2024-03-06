go build -o bookings cmd/web/*.go
./bookings.exe -dbname=Hotel1 -dbuser=postgres -dbpassword=AidosK05 -cache=false -production=false