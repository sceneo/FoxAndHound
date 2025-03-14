#!/bin/sh

# Replace placeholders with actual environment variable values

# Check if API_BASE_PATH is set
if [ -n "${BACKEND_URL}" ]; then
  # Generate json file that can be fetched from webapp
  echo "Generating config.json with API_BASE_PATH: ${BACKEND_URL}/api"
  echo "{\"apiBaseUrl\": \"${BACKEND_URL}/api\"}" > /usr/share/nginx/html/config.json
else
  # Generate empty json file
  echo "{}" > /usr/share/nginx/html/config.json
fi

PORT="${PORT:-8080}"
if [ "8080" != "${PORT}" ]; then
  echo "Setting Nginx port to ${PORT}"
  sed -i -e "s|listen 8080|listen ${PORT}|g" /etc/nginx/conf.d/default.conf
fi

# Start Nginx
nginx -g 'daemon off;'