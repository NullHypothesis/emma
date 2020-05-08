package main

// Default bridges.  See
// <https://trac.torproject.org/projects/tor/wiki/doc/TorBrowser/DefaultBridges>
// for the full list.
var defaultBridges = []string{
	"38.229.1.78:80",
	"38.229.33.83:80",
	"192.95.36.142:443",
	"37.218.240.34:40035",
	"37.218.245.14:38224",
	"85.31.186.98:443",
	"85.31.186.26:443",
	"144.217.20.138:80",
	"193.11.166.194:27015",
	"193.11.166.194:27020",
	"193.11.166.194:27025",
	"209.148.46.65:443",
	"146.57.248.225:22",
	"45.145.95.6:27015",
	"[2a0c:4d80:42:702::1]:27015",
}

// Directory authorities.  See
// <https://gitweb.torproject.org/tor.git/tree/src/app/config/auth_dirs.inc>
// for the full list.
var dirAuths = []string{
	"128.31.0.39:9131",                    // moria1
	"[2001:858:2:2:aabb:0:563b:1526]:443", // tor26
	"86.59.21.38:80",                      // tor26
	"45.66.33.45:80",                      // dizum
	"66.111.2.131:9030",                   // Serge
	"[2001:638:a000:4140::ffff:189]:443",  // gabelmoo
	"131.188.40.189:80",                   // gabelmoo
	"[2001:678:558:1000::244]:443",        // dannenberg
	"193.23.244.244:80",                   // dannenberg
	"[2001:67c:289c::9]:80",               // maatuska
	"171.25.193.9:443",                    // maatuska
	"154.35.175.225:80",                   // Faravahar
	"199.58.81.140:80",                    // longclaw
	"[2620:13:4000:6000::1000:118]:443",   // bastet
	"204.13.164.118:80",                   // bastet
}

// Guard and exit relays.  See
// <https://nusenu.github.io/OrNetStats/> to pick some.
var relays = []string{
	"193.11.166.196:443",   // Karlstad1 (Guard)
	"81.7.18.7:9001",       // Freebird32 (Guard)
	"91.143.80.147:995",    // Ichotolot63 (Guard)
	"162.247.74.7:443",     // CalyxInstitute01 (Exit)
	"62.102.148.68:443",    // weizenbaum (Exit)
	"185.220.100.253:9000", // F3Netze (Exit)
}

// Websites and a string that's meant to be in the website's body.
var websites = map[string]string{
	"https://bridges.torproject.org":                                           "The Tor Project",
	"https://torproject.org":                                                   "Tor Browser",
	"https://gettor.torproject.org":                                            "GetTor",
	"https://ajax.aspnetcdn.com":                                               "Microsoft Ajax Content Delivery Network",
	"https://archive.org/details/@gettor":                                      "torbrowser-install",
	"https://drive.google.com/drive/folders/13CADQTsCwrGsIID09YQbNz2DfRMUoxUU": "tor-browser-linux",
	"https://github.com/torproject/torbrowser-releases/":                       "GetTor",
	"https://gitlab.com/thetorproject/torbrowser-windows":                      "torbrowser",
	"https://accounts.google.com/ServiceLogin":                                 "Sign in",
	"https://mail.riseup.net/rc/":                                              "Welcome to mail.riseup.net",
}
