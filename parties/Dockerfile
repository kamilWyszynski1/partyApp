FROM python:3.6

RUN apt-get update && apt-get install -y python-pip python-dev build-essential

WORKDIR /app
COPY . /app

RUN pip install -r requirements.txt
RUN chmod +x docker-entrypoint.sh

CMD ["/bin/bash", "docker-entrypoint.sh"]