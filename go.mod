module mdnstest

go 1.13

require bonjour v0.0.0

replace bonjour => ./bonjour

require miekgdns v1.1.49

replace miekgdns => ./miekgdns
