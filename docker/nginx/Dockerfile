FROM nginx:latest AS base_stage

ADD ./etc/nginx/nginx.conf /etc/nginx/nginx.conf

FROM base_stage AS dev_stage
ADD ./etc/nginx/dev.conf /etc/nginx/default.conf

FROM base_stage AS prod_stage
ADD ./etc/nginx/prod.conf /etc/nginx/default.conf
