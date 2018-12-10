const host = process.env.REQHACK_BASEHOST
if (!host) {
  console.error("There is't REQHACK_BASEHOST")
  process.exit(1)
}
const random = process.env.REQHACK_RANDOM
const usessl = process.env.REQHACK_USE_SSL

console.log(`upstream reqhack {
      server reqhack-server:8081;
  }

  resolver 127.0.0.1;
  access_log /var/log/nginx/access2.log;
  error_log /var/log/nginx/error2.log;

  server {
      server_name ${host};
      listen 80 default_server;
      ${
        usessl
          ? `
      listen 443 ssl default_server;
      ssl_certificate /etc/nginx/ssl/cert.crt;
      ssl_certificate_key /etc/nginx/ssl/cert.prv;
      ssl_protocols        SSLv3 TLSv1;
      ssl_ciphers          HIGH:!ADH:!MD5;
      `
          : ''
      }
      location /v1/ {
          proxy_pass http://reqhack;
      }

      location / {
          root /var/www/app;
          try_files $uri /index.html;
      }
  }

  server {
      server_name ~^(?<subdomain>.*?)\.${host.replace(/\./, '\\.')};
      listen 80;
      ${
        usessl
          ? `
      listen 443 ssl;
      ssl_certificate /etc/nginx/ssl/cert.crt;
      ssl_certificate_key /etc/nginx/ssl/cert.prv;
      ssl_protocols        SSLv3 TLSv1;
      ssl_ciphers          HIGH:!ADH:!MD5;
      `
          : ''
      }
      location / {
          rewrite /(.*) /v1/$subdomain/in/$1 break;
          proxy_set_header X-Reqhack-Real-IP-${random} $remote_addr;
          proxy_set_header X-Reqhack-Real-Port-${random} $remote_user;
          proxy_set_header X-Reqhack-Real-Header-${random} $request;
          proxy_set_header Host $http_host;
          proxy_pass http://reqhack;
      }
  }
`)