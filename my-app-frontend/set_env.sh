#!/bin/zsh
sed -i  "s|<body>|<body hostname=\"$HOSTNAME\" port=\"$PORT\">|"  /usr/share/nginx/html/index.html
nginx -g "daemon off;"