FROM tutum/curl
COPY server /usr/bin/server
ENV WAIT_TIME_SECONDS=3
ENV SYS_VERSION=1.4
HEALTHCHECK --interval=30s --timeout=30s --retries=3 CMD curl -f http://localhost:8080/health/?status=400S
CMD server
