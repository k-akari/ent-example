FROM public.ecr.aws/nginx/nginx:1.23.0-alpine

RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

COPY ./build/web/nginx.conf /etc/nginx/conf.d/