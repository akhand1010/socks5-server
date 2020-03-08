Micro socks5 proxy on go.

Docker image available at ```vlakam/socks5-server``` for amd64\arm64

Environment variables:
```
SOCKS_USER=user
SOCKS_PASS=pass
```

Use it with:
```
docker run --env SOCKS_USER=user --env SOCKS_PASS=pass -p 3405:1080 vlakam/socks5-server
```
To run on 3405 port with user:pass auth.