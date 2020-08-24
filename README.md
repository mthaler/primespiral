# primespiral

A small program that visualizes prime numbers using polar coordinates. Each prime number is plotted as a point (r, theta) = (p, p) where r is the radius, theta the angle and p a prime number.

# Requirements

To build primespiral [primegen](https://github.com/jbarham/primegen) and [Go Graphics](https://github.com/fogleman/gg) need to be installed. Do `go get github.com/jbarham/primegen` and  `go get github.com/fogleman/gg` to install these packages.

# Usage

Run primespiral with `go run primesprial.go`

Open a web browser and enter `http://localhost:8080/primes?max=1000&pointsize=10`

The *max* query parameter specifies the largest prime number that is drawn and the *pointsize* parameter sets the size of the points.

# Credits

This program was inspired by the wonderful video [Why do prime numbers make these spirals?](https://www.youtube.com/watch?v=EK32jo7i5LQ)