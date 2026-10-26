FROM node:24-alpine

WORKDIR /app

RUN corepack enable

COPY package.json pnpm-lock.yaml ./

RUN pnpm install

COPY ./web ./web

EXPOSE 8080

CMD ["pnpm", "dev"]
