# Server Package

A group of administrative server packages to aid in the trafic and access control of a server application

## User Stroies

- [x] Traffic to server is routed through a reverse proxy
- [x] All interactions with server are logged through a logger package
- [ ] All traffic from the reverse proxy is routed through a firewall package
- [ ] Traffic is balanced between servers using a load balancer package
- [ ] ServerManager is used to start all packages with respective flags pulled from serverConfig
- [ ] Load balancer can call ServerManager to dynamically create and remove servers to fit load
