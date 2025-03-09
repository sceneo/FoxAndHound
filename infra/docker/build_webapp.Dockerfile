FROM node:22 AS build

WORKDIR /app

COPY package*.json .
RUN npm ci

COPY . .
RUN npm run build --omit=dev

# Stage 2: Serve the app with Nginx
FROM nginx:alpine

COPY --from=build /app/infra/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/infra/entrypoint.sh /usr/share/nginx/html/entrypoint.sh
COPY --from=build /app/dist/frontend /usr/share/nginx/html

# Expose port 80
EXPOSE 8081

RUN chmod +x /usr/share/nginx/html/entrypoint.sh
ENTRYPOINT ["/usr/share/nginx/html/entrypoint.sh"]