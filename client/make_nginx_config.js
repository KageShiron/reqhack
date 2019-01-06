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
  access_log /var/log/nginx/access.log;
  error_log /var/log/nginx/error.log;
  add_header Content-Security-Policy "default-src 'self'; script-src 'self'; img-src 'self' data:; style-src 'self' 'unsafe-inline'";

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
          set_by_lua_block $reqhack_raw_request {
            local b64 = require("ngx.base64")
            return b64.encode_base64url( ngx.req.raw_header() )
          }
          rewrite /(.*) /v1/$subdomain/in/$1 break;
          proxy_set_header X-Reqhack-Real-IP-${random} $remote_addr;
          proxy_set_header X-Reqhack-Real-UserPort-${random} $remote_port;
          proxy_set_header X-Reqhack-Real-ServerPort-${random} $server_port;
          proxy_set_header X-Reqhack-Real-User-${random} $remote_user;
          proxy_set_header X-Reqhack-Real-Scheme-${random} $scheme;
          proxy_set_header X-Reqhack-Real-Request-${random} $reqhack_raw_request;
          proxy_set_header Host $http_host;
          proxy_pass http://reqhack;
      }
  }
`)
