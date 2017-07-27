FROM tutum/curl
COPY server /usr/bin/server
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:8080/health/
CMD server
