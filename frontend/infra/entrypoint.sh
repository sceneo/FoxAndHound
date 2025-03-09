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

# Start Nginx
nginx -g 'daemon off;'