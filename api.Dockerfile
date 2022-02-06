FROM golang:1.17.6-alpine AS builder
WORKDIR /app
COPY . ./
RUN apk --no-cache add make
RUN make api.exe

FROM alpine:latest
ARG COMMIT_SHA=${COMMIT_SHA}
ENV COMMIT_SHA=${COMMIT_SHA}
EXPOSE 8080
COPY --from=builder /app/api.exe ./
CMD [ "./api.exe" ]