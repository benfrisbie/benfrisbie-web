# benfrisbie-web
benfrisbie-web is a golang web server for my personal website. It uses the [gin](https://github.com/gin-gonic/gin) http framework.

## Fresh Deployment
Follow these instructions to deploy to a brand new server.
### Install golang
```sh
sudo apt update
sudo apt install golang
```

### Build application
Copy application to server
```sh
git clone https://github.com/benfrisbie/benfrisbie-web.git
cd benfrisbie-web/
```

Install golang
```sh
sudo apt update
sudo apt install golang
```

Build
```sh
go build main.go
```

### Systemd
Create a [systemd service](https://man7.org/linux/man-pages/man5/systemd.service.5.html)
```sh
# /etc/systemd/system/benfrisbie-web.service
[Unit]
Description=benfrisbie-web service
After=network.target

[Service]
Type=simple
WorkingDirectory=/home/ubuntu/benfrisbie-web
ExecStart=/home/ubuntu/benfrisbie-web/main
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

Start service
```sh
sudo systemctl enable benfrisbie-web
sudo systemctl start benfrisbie-web
```

### Nginx
Install [nginx](https://www.nginx.com/) as our webserver.
```sh
sudo apt update
sudo apt install nginx
sudo ufw allow 'Nginx Full'
sudo ufw status
systemctl status nginx
```

Create nginx config
```sh
# /etc/nginx/conf.d/benfrisbie.com.conf
server {
    listen 80 default_server;
    listen [::]:80 default_server;
    server_name benfrisbie.com;

    location / {
      proxy_pass http://localhost:8080/;
      proxy_set_header Host $host;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection upgrade;
      proxy_set_header Accept-Encoding gzip;
    }
}
```

### Lets Encrypt
Use [Let's Encrypt](https://letsencrypt.org/) to get free TLS certificate.

Install Let's Encrypt [certbot](https://certbot.eff.org/lets-encrypt/ubuntufocal-nginx)
```
sudo snap install --classic certbot
sudo ln -s /snap/bin/certbot /usr/bin/certbot
rm /etc/nginx/sites-enabled/default
sudo certbot --nginx -d benfrisbie.com
```

Add the following line to your cron to auto renew certificate.
```sh
# sudo crontab -e
0 12 * * * sudo /usr/bin/certbot renew --quiet
```
