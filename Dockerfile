FROM golang:1.20.3-alpine

# SET THE WORKING DIRECTORY TO /APP
WORKDIR /app

# COPY THE GO.MOD AND GO.SUM FILES TO THE WORKING DIRECTORY
COPY go.mod go.sum ./

# DOWNLOAD AND INSTALL ANY REQUIRED GO DEPENDENCIES
RUN go mod download

# COPY THE ENTIRE SOURCE CODE TO THE WORKING DIRECTORY
COPY . .

# BUILD THE GO APPLICATION
RUN go build -o main .

# EXPOSE THE PORT SPECIFIED BY THE PORT ENVIRONMENT VARIABLE
EXPOSE 3000

# SET THE ENTRY POINT OF THE CONTAINER TO THE EXECUTABLE
CMD ["./main"]