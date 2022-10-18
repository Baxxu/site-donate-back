upstream backend {

	zone upstreams 64K;
	server 127.0.0.1:8080;
	keepalive 2;
}

server {

	listen 80;
	listen [::]:80;

	server_name testest.ru www.testest.ru;

	location / {

		return 301 https://testest.ru$request_uri;
	}
}

server {

	listen 443 ssl http2;
	listen [::]:443 ssl http2;

	server_name www.testest.ru;

	ssl_certificate /etc/nginx/ssl/testest.ru-reg.ru.crt;
	ssl_certificate_key /etc/nginx/ssl/testest.ru-reg.ru.key;

	ssl_session_timeout 1d;
	ssl_session_cache shared:Testest:10m; # about 40000 sessions
	ssl_session_tickets off;

	# modern configuration
	ssl_protocols TLSv1.3;
	ssl_prefer_server_ciphers off;

	# HSTS (ngx_http_headers_module is required) (63072000 seconds)
	add_header Strict-Transport-Security "max-age=63072000" always;

	# OCSP stapling
	ssl_stapling on;
	ssl_stapling_verify on;

	# verify chain of trust of OCSP response using Root CA and Intermediate certs
	ssl_trusted_certificate /etc/nginx/ssl/testest.ru-reg.ru-root.crt;

	# replace with the IP address of your resolver
	resolver 8.8.8.8;

	location / {

		return 301 https://testest.ru$request_uri;
	}
}

server {

	listen 443 ssl http2;
	listen [::]:443 ssl http2;

	server_name testest.ru;

	ssl_certificate /etc/nginx/ssl/testest.ru-reg.ru.crt;
	ssl_certificate_key /etc/nginx/ssl/testest.ru-reg.ru.key;

	ssl_session_timeout 1d;
	ssl_session_cache shared:Testest:10m; # about 40000 sessions
	ssl_session_tickets off;

	# modern configuration
	ssl_protocols TLSv1.3;
	ssl_prefer_server_ciphers off;

	# HSTS (ngx_http_headers_module is required) (63072000 seconds)
	add_header Strict-Transport-Security "max-age=63072000" always;

	# OCSP stapling
	ssl_stapling on;
	ssl_stapling_verify on;

	# verify chain of trust of OCSP response using Root CA and Intermediate certs
	ssl_trusted_certificate /etc/nginx/ssl/testest.ru-reg.ru-root.crt;

	# replace with the IP address of your resolver
	resolver 8.8.8.8;

	add_header Access-Control-Allow-Origin https://testest.ru;

	location /auth {

		proxy_pass http://backend;

		proxy_http_version 1.1;
		proxy_set_header Connection "";
	}

	location /api {

		proxy_pass http://backend;

		proxy_http_version 1.1;
		proxy_set_header Connection "";
	}

	location / {

		root /var/www/testest.ru/html/build;

		index index.html;

		try_files $uri /$uri /index.html;
	}
}