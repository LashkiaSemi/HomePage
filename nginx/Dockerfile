FROM nginx:latest

RUN mkdir /app
RUN mkdir -p /etc/nginx/logs
RUN touch /etc/nginx/logs/access.log
RUN mkdir -p /var/log/nginx
RUN touch /var/log/nginx/homepage-error.log
RUN touch /var/log/nginx/homepage-access.log

COPY ./nginx.conf /etc/nginx/nginx.conf

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]