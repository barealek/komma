FROM node:alpine AS web

WORKDIR /web

RUN --mount=type=bind,source=web/package.json,target=package.json \
    --mount=type=bind,source=web/package-lock.json,target=package-lock.json \
    --mount=type=cache,target=/root/.npm \
    npm ci


COPY web .
RUN npm run build

ENTRYPOINT [ "ls", "-a" ]
