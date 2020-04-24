# Emma

Emma is a lightweight censorship analyser that tests if the machine it's running
on can use the Tor network.  Emma does so by testing if it can

1. fetch websites such as torproject.org,

2. connect to Tor relays,

3. connect to default bridges,

4. and connect to directory authorities.

If you run emma without any arguments, it will write its analysis report to a
file.  You can pass the argument `-stdout` to make emma write its report to
stdout.

Emma is meant to be run by semi-technical Tor users who believe that they are
subject to censorship.

## Contact

Send email to Philipp Winter <phw@torproject.org>.
