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
	"216.252.162.21:46089",
	"144.217.20.138:80",
	"193.11.166.194:27015",
	"193.11.166.194:27020",
	"193.11.166.194:27025",
	"209.148.46.65:443",
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

var domains = map[string]map[string]bool{
	"www.torproject.org": map[string]bool{"95.216.163.36": true,
		"138.201.14.197": true, "2a01:4f8:172:1b46:0:abba:5:1": true, "2a01:4f9:c010:19eb::1": true},

	"bridges.torproject.org": map[string]bool{"78.47.38.229": true, "2a01:4f8:211:6e8:0:823:5:1": true},
}