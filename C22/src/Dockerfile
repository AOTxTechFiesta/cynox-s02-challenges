FROM openresty/openresty:alpine

# Install dependencies
RUN apk update 
RUN apk add --no-cache \
  lua-resty-http \
  lua-cjson
  
RUN rm -rf /var/cache/apk/*

# Create working directories
RUN mkdir -p /usr/local/openresty/nginx/html
RUN mkdir -p /usr/local/openresty/nginx/lua

# Copy files
COPY nginx.conf /usr/local/openresty/nginx/conf/nginx.conf
COPY html/ /usr/local/openresty/nginx/html/
COPY lua/ /usr/local/openresty/nginx/lua/

COPY flag.txt /flag.txt

# Expose port
EXPOSE 8080

# Start OpenResty
CMD ["openresty", "-g", "daemon off;"]
