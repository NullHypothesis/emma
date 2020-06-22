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

## Build

You need to have [golang](https://golang.org) installed to compile emma.
First, clone the repository by running

    git clone https://gitlab.torproject.org/tpo/anti-censorship/emma.git

Then, enter your newly-cloned repository and compile the tool by running

    cd emma && make

If everything worked, your emma binary will be in the `bin` directory.

## Example output

Here's what emma's output looks like:

    2020/06/22 09:26:40 Starting analysis.
    Time: 2020-06-22 09:26:40.003730038 -0700 PDT m=+0.018131500
    Testing default bridges:
      ✓ 38.229.1.78:80                        74ms
      ✓ 38.229.33.83:80                       53ms
      ✓ 192.95.36.142:443                     73ms
      ✓ 37.218.240.34:40035                  166ms
      ✓ 37.218.245.14:38224                  147ms
      ✓ 85.31.186.98:443                     167ms
      ✓ 85.31.186.26:443                     165ms
      ✓ 144.217.20.138:80                     72ms
      ✓ 193.11.166.194:27015                 169ms
      ✓ 193.11.166.194:27020                 167ms
      ✓ 193.11.166.194:27025                 167ms
      ✓ 209.148.46.65:443                     68ms
      ✓ 146.57.248.225:22                     69ms
      ✓ 45.145.95.6:27015                    155ms
      ✓ [2a0c:4d80:42:702::1]:27015          153ms
    Testing websites:
      ✓ bridges.torproject.org               664ms
      ✓ drive.google.com                    1.193s
      ✓ github.com                          3.481s
      ✓ archive.org                         15.114s
      ✓ gitlab.com                           475ms
      ✓ accounts.google.com                  130ms
      ✓ mail.riseup.net                      219ms
      ✓ torproject.org                      1.093s
      ✓ gettor.torproject.org                893ms
      ✓ ajax.aspnetcdn.com                   550ms
    Testing directory authorities:
      ✓ 128.31.0.39:9131                      74ms
      ✓ [2001:858:2:2:aabb:0:563b:1526]:443  171ms
      ✓ 86.59.21.38:80                       180ms
      ✓ 45.66.33.45:80                       143ms
      ✓ 66.111.2.131:9030                    106ms
      ✓ [2001:638:a000:4140::ffff:189]:443   163ms
      ✓ 131.188.40.189:80                    156ms
      ✓ [2001:678:558:1000::244]:443         162ms
      ✓ 193.23.244.244:80                    154ms
      ✓ [2001:67c:289c::9]:80                189ms
      ✓ 171.25.193.9:443                     177ms
      ✓ 154.35.175.225:80                     50ms
      ✓ 199.58.81.140:80                      79ms
      ✓ [2620:13:4000:6000::1000:118]:443     41ms
      ✓ 204.13.164.118:80                     35ms
    Testing relays:
      ✓ 193.11.166.196:443                   164ms
      ✓ 81.7.18.7:9001                       162ms
      ✓ 91.143.80.147:995                    162ms
      ✓ 162.247.74.7:443                      93ms
      ✓ 62.102.148.68:443                    164ms
      ✓ 185.220.100.253:9000                 171ms
    2020/06/22 09:27:08 Wrote output to: /dev/stdout

## Contact

Send email to Philipp Winter <phw@torproject.org>.
