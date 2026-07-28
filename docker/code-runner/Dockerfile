FROM python:3.13-slim

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

WORKDIR /app

RUN useradd --create-home --shell /usr/sbin/nologin runner

COPY docker/code-runner/server.py /app/server.py

RUN chown -R runner:runner /app

USER runner
EXPOSE 8080

CMD ["python", "/app/server.py"]
