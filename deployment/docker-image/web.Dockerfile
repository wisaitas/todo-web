FROM node:22-alpine as build

WORKDIR /app

COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

COPY frontend/ ./

RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL=${VITE_API_BASE_URL}

RUN echo 'server { \
    listen 8083; \
    root /usr/share/nginx/html; \
    index index.html; \
    location / { \
        try_files $uri $uri/ /index.html; \
    } \
}' > /etc/nginx/conf.d/default.conf

RUN echo '{ "API_BASE_URL": "http://localhost:8082/api/v1" }' > /usr/share/nginx/html/config.json

EXPOSE 8083

CMD ["nginx", "-g", "daemon off;"] 